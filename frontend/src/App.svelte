<script lang="ts">
  import Sidebar from "./lib/Sidebar.svelte";
  import TitleBar from "./lib/TitleBar.svelte";
  import MainContent from "./lib/MainContent.svelte";
</script>

<div class="h-screen flex flex-col overflow-hidden bg-neutral-950">
  <TitleBar />

  <div class="flex-1 flex overflow-hidden">
    <Sidebar />
    <main class="flex-1 overflow-y-auto no-scrollbar">
      <MainContent />
    </main>
  </div>
</div>