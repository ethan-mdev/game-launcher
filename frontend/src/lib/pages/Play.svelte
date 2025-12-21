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

  // Setup event listeners when component mounts
  onMount(() => {
    patching.setupEventListeners();
    
    // Check for updates if already logged in
    if ($auth.isLoggedIn && $auth.accessToken) {
      patching.checkAndDownload($auth.accessToken);
    }
  });

  onDestroy(() => {
    patching.cleanupEventListeners();
  });

  // React to login state changes
  $: if ($auth.isLoggedIn && $auth.accessToken) {
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
      email: '',
      role: 'user',
      profileImage: '',
      accessToken: '',
      refreshToken: ''
    });
  }

  function play() {
    if ($isPatchComplete) {
      // Pass username and access token to the game launcher
      // The game should authenticate using the token, not password
      StartGame($auth.username, $auth.accessToken);
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
    {:else}
      <CardFeed title="Patch Notes" items={patchNotes} />
    {/if}
  </div>

  {#if $auth.isLoggedIn}
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