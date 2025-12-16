<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import { auth } from '../stores/auth';
  import { getAvatarSrc } from '../utils/avatars';

  const dispatch = createEventDispatcher();
  
  export let username: string;
  
  // Get avatar source from profile image filename
  $: profileImageSrc = getAvatarSrc($auth.profileImage);
</script>

<div class="flex items-center justify-between px-6 py-4 border-b border-white/5">
  <div class="flex items-center gap-3">
    <div class="w-8 h-8 bg-white/10 rounded-full flex items-center justify-center text-sm text-white">
      <img src={profileImageSrc} alt="User Avatar" class="w-8 h-8 rounded-full border border-neutral-700 object-cover" />
    </div>
    <span class="text-white text-sm font-medium">{username}</span>
  </div>
  <button
    on:click={() => dispatch('logout')}
    class="text-neutral-500 text-sm hover:text-white transition"
  >
    Sign out
  </button>
</div>