<script lang="ts">
  import { onMount } from 'svelte';
  import TopBar from '$lib/components/TopBar.svelte';
  import LeftPanel from '$lib/components/LeftPanel.svelte';
  import Toolbar from '$lib/components/Toolbar.svelte';
  import OpStrip from '$lib/components/OpStrip.svelte';
  import FormatRow from '$lib/components/FormatRow.svelte';
  import FileList from '$lib/components/FileList.svelte';
  import ArrowCol from '$lib/components/ArrowCol.svelte';
  import RightPanel from '$lib/components/RightPanel.svelte';
  import StatusBar from '$lib/components/StatusBar.svelte';
  import ContextMenu from '$lib/components/ContextMenu.svelte';
  import IdPickerModal from '$lib/components/IdPickerModal.svelte';
  import SourcesModal from '$lib/components/SourcesModal.svelte';
  import ProgressToast from '$lib/components/ProgressToast.svelte';
  import { activeTab, theme, connectWS } from '$lib/stores';

  const NON_RENAME_TABS = ['metadata', 'nfo', 'images', 'subtitles', 'verify', 'history'];

  onMount(() => {
    // Restore saved theme
    const saved = localStorage.getItem('bmm-theme');
    if (saved) {
      theme.set(saved);
      document.documentElement.setAttribute('data-theme', saved);
    }
    // Connect WebSocket for progress events
    connectWS();
  });
</script>

<div class="flex flex-col" style="height:100vh; overflow:hidden;">
  <TopBar />

  <!-- Main layout area -->
  <div class="flex flex-1 overflow-hidden">
    {#if $activeTab === 'rename'}
      <!-- Left: source folder tree -->
      <LeftPanel />

      <!-- Center: workflow -->
      <section class="flex flex-col flex-1 overflow-hidden">
        <Toolbar />
        <OpStrip />
        <FormatRow />

        <!-- Split view: source | arrows | renamed -->
        <div class="flex flex-1 overflow-hidden">
          <FileList mode="source" />
          <ArrowCol />
          <FileList mode="renamed" />
        </div>
      </section>

      <!-- Right: inspector panel -->
      <RightPanel />

    {:else if NON_RENAME_TABS.includes($activeTab)}
      <div class="flex-1 flex items-center justify-center text-base-content/30 text-sm">
        <div class="text-center">
          <div class="text-4xl mb-3">🚧</div>
          <div class="font-medium capitalize">{$activeTab}</div>
          <div class="text-xs mt-1">Coming in Phase {$activeTab === 'metadata' ? '2–3' : $activeTab === 'nfo' ? '5' : $activeTab === 'images' ? '6' : $activeTab === 'subtitles' ? '7' : $activeTab === 'verify' ? '8' : '4'}</div>
        </div>
      </div>
    {/if}
  </div>

  <StatusBar />
</div>

<!-- Modals & floating overlays -->
<ContextMenu />
<IdPickerModal />
<SourcesModal />
<ProgressToast />
