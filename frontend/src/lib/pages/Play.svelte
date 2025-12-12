<script lang="ts">
  import AuthForm from '../components/AuthForm.svelte';
  import UserBar from '../components/UserBar.svelte';
  import CardFeed from '../components/CardFeed.svelte';
  import PlayBar from '../components/PlayBar.svelte';
  import { auth } from '../stores/auth';
  import { Logout } from '../../../wailsjs/go/backend/AuthService';

  // Patching state
  let isPatchComplete = true;
  let patchProgress = 100;
  let patchStatus = 'Ready to play';
  let downloadSpeed = '12.5 MB/s';

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
    console.log('Launching game...');
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
      {isPatchComplete} 
      {patchProgress} 
      {patchStatus} 
      {downloadSpeed}
      version="v1.3.1"
      on:play={play} 
    />
  {/if}
</div>