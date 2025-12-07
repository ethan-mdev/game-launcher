<script lang="ts">
  import AuthForm from '../components/AuthForm.svelte';
  import UserBar from '../components/UserBar.svelte';
  import CardFeed from '../components/CardFeed.svelte';
  import PlayBar from '../components/PlayBar.svelte';

  let username = '';
  let password = '';
  let rememberMe = false;
  let isLoggedIn = false;

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
        'New crystalline cosmetics + wings',
      ],
    },
    {
      title: 'v1.3.0 — The Awakening',
      date: 'November 20, 2025',
      content: [
        'New zone: Crystal Caverns',
        'Level cap increased to 60',
        'New raid: Heart of Stone',
      ],
    },
  ];

  function handleLogin() {
    // Add actual login logic here
    isLoggedIn = true;
  }

  function handleRegister() {
    // Add actual register logic here
    console.log('Registering...');
  }

  function logout() {
    isLoggedIn = false;
    username = '';
    password = '';
  }

  function play() {
    console.log('Launching game...');
  }
</script>

<div class="flex flex-col h-full">
  {#if isLoggedIn}
    <UserBar {username} on:logout={logout} />
  {/if}

  <div class="flex-1 overflow-y-auto no-scrollbar">
    {#if !isLoggedIn}
      <AuthForm 
        bind:username 
        bind:password
        on:login={handleLogin}
        on:register={handleRegister}
      />
    {:else}
      <CardFeed title="Patch Notes" items={patchNotes} />
    {/if}
  </div>

  {#if isLoggedIn}
    <PlayBar 
      {isPatchComplete} 
      {patchProgress} 
      {patchStatus} 
      {downloadSpeed}
      on:play={play} 
    />
  {/if}
</div>