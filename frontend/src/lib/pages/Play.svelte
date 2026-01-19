<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import AuthForm from '../components/AuthForm.svelte';
  import UserBar from '../components/UserBar.svelte';
  import CardFeed from '../components/CardFeed.svelte';
  import PlayBar from '../components/PlayBar.svelte';
  import { auth } from '../stores/auth';
  import { patching, isPatchComplete, patchStatusText } from '../stores/patching';
  import { Logout } from '../../../wailsjs/go/backend/AuthService';
  import { StartGame } from '../../../wailsjs/go/backend/App';

  onMount(() => {
    patching.setupEventListeners();
    
    if ($auth.isLoggedIn && $auth.accessToken) {
      patching.checkAndDownload($auth.accessToken);
    }
  });

  onDestroy(() => {
    patching.cleanupEventListeners();
  });

  $: if ($auth.isLoggedIn && $auth.accessToken && $auth.gameLinked) {
    patching.checkAndDownload($auth.accessToken);
  }

  const patchNotes = [
    {
      title: 'v1.3.1 — Crystalbound',
      date: 'December 5, 2025',
      featured: true,
      content: [
        'New emerald dungeon with scaling mechanics',
        'Forest Trials rotation begins weekly',
        'Boss AI updated in Aeridor Depths',
        'New crystalline cosmetics + wings',
      ],
    },
  ];

  async function logout() {
    try {
      await Logout($auth.refreshToken);
    } catch (err) {
      console.error('Logout error:', err);
    }
    auth.set({
      isLoggedIn: false,
      userId: '',
      username: '',
      role: 'user',
      profileImage: '',
      accessToken: '',
      refreshToken: '',
      gameApiKey: '',
      gameLinked: false
    });
  }

  async function play() {
    if ($isPatchComplete && $auth.gameLinked) {
      try {
        await StartGame($auth.username, $auth.gameApiKey);
      } catch (err) {
        console.error('Failed to start game:', err);
      }
    }
  }
</script>

<div class="flex flex-col h-full">
  {#if $auth.isLoggedIn}
    <UserBar username={$auth.username} on:logout={logout} />
  {/if}

  <div class="flex-1 overflow-y-auto no-scrollbar">
    {#if !$auth.isLoggedIn}
      <AuthForm />
    {:else if !$auth.gameLinked}
      <div class="h-full flex items-center justify-center p-8">
        <div class="text-center space-y-4">
          <div class="text-yellow-500 text-5xl mb-4">⚠️</div>
          <h2 class="text-white text-xl font-medium">Account Not Verified</h2>
          <p class="text-neutral-400 max-w-sm">
            Verify your account on Discord to create your game account and start playing.
          </p>
        </div>
      </div>
    {:else}
      <CardFeed title="Patch Notes" items={patchNotes} />
    {/if}
  </div>

  {#if $auth.isLoggedIn && $auth.gameLinked}
    <PlayBar 
      isPatchComplete={$isPatchComplete} 
      patchProgress={$patching.progress} 
      patchStatus={$patchStatusText} 
      downloadSpeed={$patching.downloadSpeed}
      version={$patching.version}
      on:play={play} 
    />
  {/if}
</div>