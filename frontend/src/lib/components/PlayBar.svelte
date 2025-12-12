<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { StartGame } from '../../../wailsjs/go/backend/App';
  
  const dispatch = createEventDispatcher();
  
  export let isPatchComplete = true;
  export let patchProgress = 100;
  export let patchStatus = 'Ready to play';
  export let downloadSpeed = '';
  export let version = 'v1.0.0';

  async function handlePlay() {
    await StartGame('Ember', 'test');
  }
</script>

<div class="p-6 bg-gradient-to-t from-neutral-900 to-transparent">
  <div class="flex items-center gap-6">
    <!-- Progress Section -->
    <div class="flex-1">
      <div class="flex items-center justify-between mb-2">
        <div class="flex items-center gap-2">
          {#if isPatchComplete}
            <span class="w-2 h-2 bg-green-500 rounded-full"></span>
          {:else}
            <span class="w-2 h-2 bg-yellow-500 rounded-full animate-pulse"></span>
          {/if}
          <span class="text-white text-sm font-medium">{patchStatus}</span>
        </div>
        <span class="text-neutral-600 text-xs">
          {#if !isPatchComplete}{downloadSpeed} • {/if}{version}
        </span>
      </div>
      <div class="w-full h-1 bg-neutral-800 rounded-full overflow-hidden">
        <div
          class="h-full rounded-full transition-all duration-500 {isPatchComplete
            ? 'bg-green-500'
            : 'bg-gradient-to-r from-yellow-500 to-yellow-400'}"
          style="width: {patchProgress}%"
        ></div>
      </div>
    </div>

    <!-- Play Button -->
    <button
      on:click={handlePlay}
      disabled={!isPatchComplete}
      class="px-12 py-4 text-lg font-bold rounded-xl transition-all duration-200
        {isPatchComplete
          ? 'bg-gradient-to-r from-green-500 to-emerald-600 text-white hover:scale-105 hover:shadow-xl hover:shadow-green-500/30'
          : 'bg-neutral-800/50 text-neutral-600 cursor-not-allowed'}"
    >
      {#if !isPatchComplete}
        Updating {patchProgress}%
      {:else}
        Play
      {/if}
    </button>
  </div>
</div>