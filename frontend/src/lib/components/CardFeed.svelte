<script lang="ts">
  export let title: string;
  export let items: Array<{
    title: string;
    date: string;
    featured?: boolean;
    content: string | string[];
  }> = [];
</script>

<div class="p-6 flex flex-col gap-4">
  <h1 class="text-white text-2xl font-bold">{title}</h1>

  {#each items as item}
    <div class="bg-neutral-900 rounded-lg p-5 border border-neutral-800">
      <div class="flex justify-between mb-2">
        {#if item.featured}
          <span class="text-yellow-500 text-xs font-medium">FEATURED</span>
        {/if}
        <span class="text-neutral-500 text-xs ml-auto">{item.date}</span>
      </div>
      <h3 class="text-white font-semibold text-lg">{item.title}</h3>
      
      {#if Array.isArray(item.content)}
        <ul class="mt-3 space-y-1.5">
          {#each item.content as note}
            <li class="text-neutral-400 text-sm flex items-start gap-2">
              <span class="text-neutral-600">•</span>
              {note}
            </li>
          {/each}
        </ul>
      {:else}
        <p class="text-neutral-400 text-sm mt-2">{item.content}</p>
      {/if}
    </div>
  {/each}
</div>