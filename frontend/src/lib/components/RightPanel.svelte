<script lang="ts">
  import MetadataTab from './MetadataTab.svelte';
  import ImagesTab from './ImagesTab.svelte';
  import CompanionsTab from './CompanionsTab.svelte';
  import NfoTab from './NfoTab.svelte';
  import { rightTab, rightCollapsed } from '$lib/stores';

  const TABS = ['Metadata', 'Images', 'Companions', 'NFO'];
</script>

<aside
  class="bg-base-100 border-l border-base-300 flex shrink-0 overflow-hidden transition-[width] duration-200"
  style="width: {$rightCollapsed ? '28px' : '330px'}; min-width: {$rightCollapsed ? '28px' : '330px'};"
>
  <!-- Collapsed state: rotated label only -->
  {#if $rightCollapsed}
    <button
      class="w-full h-full bg-base-200 border-none cursor-pointer text-base-content/50 hover:text-base-content flex flex-col items-center pt-3 gap-1.5 border-l border-base-300"
      onclick={() => rightCollapsed.set(false)}
      title="Expand inspector"
    >
      ▶
      <span class="text-[9px] font-bold uppercase tracking-widest" style="writing-mode:vertical-rl;">Inspector</span>
    </button>
  {:else}
    <!-- Full panel -->
    <div class="flex flex-col flex-1 overflow-hidden">
      <!-- Tabs -->
      <div class="flex border-b border-base-300 bg-base-200 shrink-0">
        {#each TABS as tab}
          <button
            class="flex-1 py-[7px] px-1 text-center text-[10px] font-semibold border-b-2 cursor-pointer transition-colors {$rightTab === tab.toLowerCase() ? 'text-primary border-b-primary' : 'text-base-content/50 border-b-transparent hover:text-base-content'}"
            onclick={() => rightTab.set(tab.toLowerCase())}
          >
            {tab}
          </button>
        {/each}
      </div>

      <!-- Tab content -->
      {#if $rightTab === 'metadata'}
        <MetadataTab />
      {:else if $rightTab === 'images'}
        <ImagesTab />
      {:else if $rightTab === 'companions'}
        <CompanionsTab />
      {:else if $rightTab === 'nfo'}
        <NfoTab />
      {/if}
    </div>

    <!-- Collapse button strip (right edge) -->
    <button
      class="w-7 h-full bg-base-200 border-none cursor-pointer text-base-content/50 hover:text-base-content flex flex-col items-center pt-3 gap-1.5 border-l border-base-300 shrink-0"
      onclick={() => rightCollapsed.set(true)}
      title="Collapse inspector"
    >
      ▶
      <span class="text-[9px] font-bold uppercase tracking-widest" style="writing-mode:vertical-rl;">Inspector</span>
    </button>
  {/if}
</aside>
