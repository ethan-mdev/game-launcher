<script lang="ts">
  import { createEventDispatcher } from 'svelte';
  import avatar from '../../assets/images/avatar-1.png';

  const dispatch = createEventDispatcher();
  
  export let username: string;
</script>

<div class="flex items-center justify-between px-6 py-4 border-b border-white/5">
  <div class="flex items-center gap-3">
    <div class="w-8 h-8 bg-white/10 rounded-full flex items-center justify-center text-sm text-white">
      <img src={avatar} alt="User Avatar" class="w-8 h-8 rounded-full" />
    </div>
    <span class="text-white text-sm font-medium">{username}</span>
    <span class="w-1.5 h-1.5 bg-green-500 rounded-full"></span>
  </div>
  <button
    on:click={() => dispatch('logout')}
    class="text-neutral-500 text-sm hover:text-white transition"
  >
    Sign out
  </button>
</div>