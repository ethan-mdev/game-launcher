<script lang="ts">
  import { WindowMinimise, Quit } from '../../wailsjs/runtime/runtime'
</script>

<div class="flex items-center justify-end px-5 py-3 bg-neutral-950 border-b border-neutral-800" style="--wails-draggable: drag">
  
  <div class="flex" style="--wails-draggable: no-drag">
    <button 
      class="w-6 h-4 flex items-center justify-center text-gray-300 hover:bg-neutral-700"
      on:click={WindowMinimise}>─</button>
    <button 
      class="w-6 h-4 flex items-center justify-center text-gray-300 hover:bg-red-600"
      on:click={Quit}>✕</button>
  </div>
</div>