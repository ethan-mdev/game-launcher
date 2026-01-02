<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { Login, Register, GetProfile, GetGameCredentials } from '../../../wailsjs/go/backend/AuthService';
  import { auth } from '../stores/auth';
  
  const dispatch = createEventDispatcher();
  
  export let username = '';
  export let password = '';
  
  let isRegistering = false;
  let email = '';
  let confirmPassword = '';
  let error = '';
  let loading = false;

  function getUserIdFromToken(token: string): string {
    try {
      const payload = token.split('.')[1];
      const decoded = JSON.parse(atob(payload));
      return decoded.sub || decoded.user_id || '';
    } catch {
      return '';
    }
  }

  async function handleSubmit() {
    error = '';
    loading = true;

    try {
      let response;
      if (isRegistering) {
        if (password !== confirmPassword) {
          error = 'Passwords do not match';
          return;
        }
        response = await Register(username, email, password);
      } else {
        response = await Login(username, password);
      }

      const userId = getUserIdFromToken(response.access_token);
      const profile = await GetProfile(userId);
      
      // Try to fetch game credentials
      let gameApiKey = '';
      let gameLinked = false;
      try {
        const gameCreds = await GetGameCredentials(response.access_token);
        gameApiKey = gameCreds.api_key;
        gameLinked = true;
      } catch (err) {
        // account_not_linked is expected for unverified users
        console.log('Game account not linked:', err);
      }
      
      auth.set({
        isLoggedIn: true,
        userId: profile.user_id,
        username: profile.username,
        role: profile.role,
        profileImage: profile.profile_image || '',
        accessToken: response.access_token,
        refreshToken: response.refresh_token,
        gameApiKey,
        gameLinked
      });
      
      dispatch(isRegistering ? 'register' : 'login');
    } catch (err) {
      error = err as string;
    } finally {
      loading = false;
    }
  }
</script>

<div class="h-full flex items-center justify-center p-8">
  <div class="w-full max-w-xs space-y-6">
    <div class="text-center mb-8">
      <h1 class="text-white text-lg font-medium">
        {isRegistering ? 'Create an account' : 'Sign in to continue'}
      </h1>
    </div>

    {#if error}
      <div class="bg-red-500/10 border border-red-500/20 rounded-xl px-4 py-3 text-red-400 text-sm">
        {error}
      </div>
    {/if}

    <div class="space-y-3">
      <input
        type="text"
        bind:value={username}
        placeholder="Username"
        class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
      />
      
      {#if isRegistering}
        <input
          type="email"
          bind:value={email}
          placeholder="Email"
          class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
        />
      {/if}
      
      <input
        type="password"
        bind:value={password}
        placeholder="Password"
        class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
      />
      
      {#if isRegistering}
        <input
          type="password"
          bind:value={confirmPassword}
          placeholder="Confirm Password"
          class="w-full bg-white/5 text-white px-4 py-3.5 rounded-xl border border-white/10 focus:border-yellow-500/50 focus:bg-white/10 focus:outline-none transition placeholder:text-neutral-500"
        />
      {/if}
    </div>

    <button
      on:click={handleSubmit}
      disabled={loading}
      class="w-full py-3.5 bg-gradient-to-r from-yellow-500 to-yellow-600 text-black font-semibold rounded-xl hover:opacity-90 transition shadow-lg shadow-yellow-500/20 disabled:opacity-50 disabled:cursor-not-allowed"
    >
      {#if loading}
        Loading...
      {:else}
        {isRegistering ? 'Create Account' : 'Sign In'}
      {/if}
    </button>

    <p class="text-center text-neutral-600 text-sm">
      {#if isRegistering}
        Already have an account?
        <button on:click={() => isRegistering = false} class="text-yellow-500 hover:text-yellow-400 transition">
          Sign in
        </button>
      {:else}
        Need an account?
        <button on:click={() => isRegistering = true} class="text-yellow-500 hover:text-yellow-400 transition">
          Register
        </button>
      {/if}
    </p>
  </div>
</div>