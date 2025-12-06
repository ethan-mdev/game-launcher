<script lang="ts">
  import { currentPage } from './stores/navigation';
  import Play from './pages/Play.svelte';
  import News from './pages/News.svelte';
  import PatchNotes from './pages/PatchNotes.svelte';
  import Settings from './pages/Settings.svelte';
</script>

<div class="h-full">
  {#if $currentPage === 'play'}
    <Play />
  {:else if $currentPage === 'news'}
    <News />
  {:else if $currentPage === 'patchnotes'}
    <PatchNotes />
  {:else if $currentPage === 'settings'}
    <Settings />
  {/if}
</div>