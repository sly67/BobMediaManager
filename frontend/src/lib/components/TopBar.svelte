<script lang="ts">
  import ThemePicker from './ThemePicker.svelte';
  import { activeTab, sourcesModalOpen, fileStats } from '$lib/stores';

  const TABS = [
    { id: 'rename',    label: 'Rename & Sort' },
    { id: 'metadata',  label: 'Metadata' },
    { id: 'nfo',       label: 'NFO' },
    { id: 'images',    label: 'Images' },
    { id: 'subtitles', label: 'Subtitles' },
    { id: 'verify',    label: 'Verify' },
    { id: 'history',   label: 'History' },
  ];
</script>

<header class="flex items-center gap-2 px-3 h-[46px] bg-base-100 border-b border-base-300 shrink-0 overflow-x-auto">
  <!-- Logo -->
  <span class="font-bold text-[15px] tracking-tight shrink-0 mr-1">
    <span class="text-primary">BobMedia</span><span class="text-secondary">Manager</span>
  </span>
  <div class="w-px h-5 bg-base-300 shrink-0"></div>

  <!-- Tabs -->
  <nav class="flex gap-0.5 flex-1 overflow-x-auto">
    {#each TABS as tab}
      <button
        class="px-3 py-1 rounded-md text-xs font-medium whitespace-nowrap cursor-pointer transition-colors {$activeTab === tab.id ? 'bg-base-200 text-base-content' : 'text-base-content/50 hover:bg-base-200 hover:text-base-content'}"
        onclick={() => activeTab.set(tab.id)}
      >
        {tab.label}
      </button>
    {/each}
    <button
      class="px-3 py-1 rounded-md text-xs font-medium whitespace-nowrap cursor-pointer transition-colors text-secondary {$activeTab === 'sources' ? 'bg-base-200' : 'hover:bg-base-200'}"
      onclick={() => { activeTab.set('sources'); sourcesModalOpen.set(true); }}
    >
      ⚙ Sources
    </button>
  </nav>

  <div class="flex items-center gap-2 shrink-0">
    <ThemePicker />
    <div class="w-px h-5 bg-base-300"></div>
    <button class="btn btn-ghost btn-sm text-xs">Match All</button>
    <button class="btn btn-primary btn-sm text-xs">
      ▶ Execute {$fileStats.matched > 0 ? `(${$fileStats.matched})` : ''}
    </button>
  </div>
</header>
