import { writable, derived } from 'svelte/store';
import { EventsOn, EventsOff } from '../../../wailsjs/runtime/runtime';
import { CheckForUpdates, DownloadUpdates } from '../../../wailsjs/go/backend/PatchingService';
import type { backend } from '../../../wailsjs/go/models';

export interface PatchingState {
  status: 'idle' | 'checking' | 'downloading' | 'complete' | 'error';
  progress: number;
  currentFile: string;
  totalFiles: number;
  completedFiles: number;
  downloadSpeed: string;
  version: string;
  error: string | null;
  filesToUpdate: backend.FileHash[];
}

const initialState: PatchingState = {
  status: 'idle',
  progress: 0,
  currentFile: '',
  totalFiles: 0,
  completedFiles: 0,
  downloadSpeed: '',
  version: 'unknown',
  error: null,
  filesToUpdate: []
};

function createPatchingStore() {
  const { subscribe, set, update } = writable<PatchingState>(initialState);

  let speedSamples: number[] = [];
  let lastBytesTime = Date.now();

  function formatSpeed(bytesPerSecond: number): string {
    if (bytesPerSecond >= 1024 * 1024) {
      return `${(bytesPerSecond / (1024 * 1024)).toFixed(1)} MB/s`;
    } else if (bytesPerSecond >= 1024) {
      return `${(bytesPerSecond / 1024).toFixed(1)} KB/s`;
    }
    return `${bytesPerSecond.toFixed(0)} B/s`;
  }

  function setupEventListeners() {
    // Listen for download progress
    EventsOn('patch:progress', (data: { current: number; total: number; file: string; status: string }) => {
      update(state => ({
        ...state,
        status: 'downloading',
        currentFile: data.file,
        totalFiles: data.total,
        completedFiles: data.current - 1,
        progress: ((data.current - 1) / data.total) * 100
      }));
    });

    // Listen for file completion
    EventsOn('patch:file-complete', (data: { file: string; progress: number }) => {
      update(state => ({
        ...state,
        completedFiles: state.completedFiles + 1,
        progress: data.progress
      }));
    });

    // Listen for patch complete
    EventsOn('patch:complete', () => {
      update(state => ({
        ...state,
        status: 'complete',
        progress: 100,
        currentFile: '',
        downloadSpeed: '',
        error: null
      }));
    });
  }

  function cleanupEventListeners() {
    EventsOff('patch:progress');
    EventsOff('patch:file-complete');
    EventsOff('patch:complete');
  }

  async function checkForUpdates(accessToken: string): Promise<boolean> {
    update(state => ({ ...state, status: 'checking', error: null }));

    try {
      const result = await CheckForUpdates(accessToken);
      
      if (result.needsUpdate) {
        update(state => ({
          ...state,
          status: 'idle',
          version: result.serverVersion,
          filesToUpdate: result.filesToUpdate,
          totalFiles: result.filesToUpdate.length
        }));
        return true;
      } else {
        update(state => ({
          ...state,
          status: 'complete',
          progress: 100,
          version: result.currentVersion,
          filesToUpdate: []
        }));
        return false;
      }
    } catch (error) {
      update(state => ({
        ...state,
        status: 'error',
        error: error instanceof Error ? error.message : 'Failed to check for updates'
      }));
      return false;
    }
  }

  async function downloadUpdates(accessToken: string): Promise<void> {
    update(state => {
      if (state.filesToUpdate.length === 0) {
        return state;
      }
      return {
        ...state,
        status: 'downloading',
        progress: 0,
        completedFiles: 0,
        error: null
      };
    });

    let currentState: PatchingState = initialState;
    subscribe(s => currentState = s)();

    if (currentState.filesToUpdate.length === 0) {
      return;
    }

    try {
      await DownloadUpdates(currentState.filesToUpdate, accessToken);
      // patch:complete event will handle the final state update
    } catch (error) {
      update(state => ({
        ...state,
        status: 'error',
        error: error instanceof Error ? error.message : 'Download failed'
      }));
    }
  }

  async function checkAndDownload(accessToken: string): Promise<void> {
    const needsUpdate = await checkForUpdates(accessToken);
    if (needsUpdate) {
      await downloadUpdates(accessToken);
    }
  }

  function reset() {
    set(initialState);
  }

  return {
    subscribe,
    setupEventListeners,
    cleanupEventListeners,
    checkForUpdates,
    downloadUpdates,
    checkAndDownload,
    reset
  };
}

export const patching = createPatchingStore();

// Derived stores for convenience
export const isPatchComplete = derived(patching, $p => $p.status === 'complete');
export const isPatching = derived(patching, $p => $p.status === 'downloading' || $p.status === 'checking');
export const patchError = derived(patching, $p => $p.error);

export const patchStatusText = derived(patching, $p => {
  switch ($p.status) {
    case 'idle':
      return $p.filesToUpdate.length > 0 ? `${$p.filesToUpdate.length} files need updating` : 'Checking...';
    case 'checking':
      return 'Checking for updates...';
    case 'downloading':
      return `Downloading ${$p.currentFile}`;
    case 'complete':
      return 'Ready to play';
    case 'error':
      return $p.error || 'Error occurred';
    default:
      return 'Unknown';
  }
});
