<script lang="ts">
  import { resolution, windowed, masterVolume, musicVolume, sfxVolume } from '../stores/settings';
  
  const resolutions = [
    '1280x720',
    '1366x768',
    '1600x900',
    '1920x1080',
    '2560x1440',
    '3840x2160'
  ];
</script>

<div class="p-6 flex flex-col gap-6">
  <div>
    <div class="flex items-center justify-between mb-1">
    <h1 class="text-white text-2xl font-bold">Settings</h1>
    <button class="px-4 py-1 bg-yellow-500 text-black font-semibold rounded-lg hover:bg-yellow-600 transition">Save</button>
    </div>
    <p class="text-neutral-500 text-sm mt-1">Configure your game preferences</p>
  </div>
  
  <!-- Display Settings -->
  <div class="bg-neutral-900 rounded-xl p-6 border border-neutral-800">
    <div class="flex items-center gap-3 mb-6">
      <h3 class="text-white font-semibold">Display</h3>
    </div>
    
    <div class="flex flex-col gap-5">
      <!-- Resolution -->
      <div class="flex items-center justify-between">
        <div>
          <span class="text-white text-sm">Resolution</span>
          <p class="text-neutral-500 text-xs mt-0.5">Set your preferred screen resolution</p>
        </div>
        <select 
          bind:value={$resolution}
          class="bg-neutral-800 text-white px-4 py-2.5 rounded-lg border border-neutral-700 hover:border-neutral-600 focus:border-yellow-500 focus:outline-none cursor-pointer min-w-[140px]"
        >
          {#each resolutions as res}
            <option value={res}>{res}</option>
          {/each}
        </select>
      </div>
      
      <div class="border-t border-neutral-800"></div>
      
      <!-- Windowed Mode -->
      <label class="flex items-center justify-between cursor-pointer group">
        <div>
          <span class="text-white text-sm">Windowed mode</span>
          <p class="text-neutral-500 text-xs mt-0.5">Run the game in a window instead of fullscreen</p>
        </div>
        <div class="relative">
          <input 
            type="checkbox" 
            bind:checked={$windowed}
            class="sr-only peer" 
          />
          <div class="w-11 h-6 bg-neutral-700 rounded-full peer peer-checked:bg-yellow-500 transition-colors"></div>
          <div class="absolute left-1 top-1 w-4 h-4 bg-white rounded-full peer-checked:translate-x-5 transition-transform"></div>
        </div>
      </label>
    </div>
  </div>

  <!-- Sound Settings -->
  <div class="bg-neutral-900 rounded-xl p-6 border border-neutral-800">
    <div class="flex items-center gap-3 mb-6">
      <h3 class="text-white font-semibold">Sound</h3>
    </div>
    
    <div class="flex flex-col gap-5">
      <!-- Master Volume -->
      <div class="flex flex-col gap-2">
        <div class="flex items-center justify-between">
          <span class="text-white text-sm">Master Volume</span>
          <span class="text-neutral-400 text-sm">{$masterVolume}%</span>
        </div>
        <input 
          type="range" 
          min="0" 
          max="100" 
          bind:value={$masterVolume}
          class="w-full h-2 bg-neutral-700 rounded-full appearance-none cursor-pointer accent-yellow-500"
        />
      </div>
      
      <div class="border-t border-neutral-800"></div>
      
      <!-- Music Volume -->
      <div class="flex flex-col gap-2">
        <div class="flex items-center justify-between">
          <span class="text-white text-sm">Music</span>
          <span class="text-neutral-400 text-sm">{$musicVolume}%</span>
        </div>
        <input 
          type="range" 
          min="0" 
          max="100" 
          bind:value={$musicVolume}
          class="w-full h-2 bg-neutral-700 rounded-full appearance-none cursor-pointer accent-yellow-500"
        />
      </div>
      
      <div class="border-t border-neutral-800"></div>
      
      <!-- SFX Volume -->
      <div class="flex flex-col gap-2">
        <div class="flex items-center justify-between">
          <span class="text-white text-sm">Sound Effects</span>
          <span class="text-neutral-400 text-sm">{$sfxVolume}%</span>
        </div>
        <input 
          type="range" 
          min="0" 
          max="100" 
          bind:value={$sfxVolume}
          class="w-full h-2 bg-neutral-700 rounded-full appearance-none cursor-pointer accent-yellow-500"
        />
      </div>
    </div>
  </div>
</div>