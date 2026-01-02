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
      <!-- Account not verified prompt -->
      <div class="h-full flex items-center justify-center p-8">
        <div class="text-center space-y-4">
          <h2 class="text-white text-xl font-medium">Account Not Verified</h2>
          <p class="text-neutral-400 max-w-sm">
            Please verify your account on the website to link your game account and start playing.
          </p>
          <a 
            href="https://dashboard.ethan-mdev.com" 
            target="_blank"
            class="inline-block mt-4 px-6 py-3 bg-gradient-to-r from-yellow-500 to-yellow-600 text-black font-semibold rounded-xl hover:opacity-90 transition"
          >
            Verify Account
          </a>
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