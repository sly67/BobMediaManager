<script lang="ts">
  import type { MediaFile } from '$lib/api';
  import { selection, selectedFile, openCtxMenu } from '$lib/stores';

  let {
    file,
    index,
    renamed = false,
    onselect,
  }: {
    file: MediaFile;
    index: number;
    renamed?: boolean;
    onselect: (e: MouseEvent, file: MediaFile, index: number) => void;
  } = $props();

  function handleClick(e: MouseEvent) {
    selectedFile.set(file);
    onselect(e, file, index);
  }

  function handleCtx(e: MouseEvent) {
    e.preventDefault();
    selectedFile.set(file);
    openCtxMenu(e, file);
  }

  const statusColors = {
    matched:   'border-l-success  bg-success/10',
    pending:   'border-l-warning  bg-warning/10',
    unmatched: 'border-l-error    bg-error/10',
  } as Record<string, string>;

  const dotColors = {
    matched:   'bg-success',
    pending:   'bg-warning',
    unmatched: 'bg-error',
  } as Record<string, string>;

  const displayName = $derived(renamed ? (file.renamedPath ?? file.name) : file.name);
  const isSelected = $derived($selection.has(file.id));
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
  class="file-row flex items-center gap-1.5 px-3 py-[7px] border-b border-base-300 cursor-pointer border-l-[3px] {statusColors[file.matchStatus] ?? ''} {isSelected ? 'file-selected' : 'hover:bg-base-200'}"
  onclick={handleClick}
  oncontextmenu={handleCtx}
>
  <span class="w-[7px] h-[7px] rounded-full shrink-0 {dotColors[file.matchStatus] ?? 'bg-base-content/30'}"></span>
  <span class="text-sm shrink-0">{file.isFolder ? '📁' : '🎞️'}</span>
  <div class="flex-1 min-w-0">
    <div class="file-name text-xs font-medium truncate {renamed && file.matchStatus === 'pending' ? 'text-warning' : ''} {renamed && file.matchStatus === 'unmatched' ? 'text-error' : ''}">
      {displayName}
    </div>
    {#if !renamed}
      <div class="text-[10px] text-base-content/50 mt-0.5">
        {#if file.size}
          {(file.size / 1e9).toFixed(1)} GB ·
        {/if}
        {file.name.split('.').pop()?.toLowerCase() ?? ''}
        {#if file.resolvedIds && Object.keys(file.resolvedIds).length > 0}
          · {Object.entries(file.resolvedIds).map(([k, v]) => `${k.toUpperCase()}:${v}`).join(' ')}
        {/if}
      </div>
    {:else}
      <div class="text-[10px] text-base-content/50 mt-0.5">
        {#if file.resolvedIds && Object.keys(file.resolvedIds).length > 0}
          {Object.entries(file.resolvedIds).slice(0, 2).map(([k, v]) => `${k.toUpperCase()}:${v}`).join(' · ')}
        {:else if file.matchStatus === 'pending'}
          Multiple matches — click to pick
        {:else if file.matchStatus === 'unmatched'}
          Click to open ID picker
        {/if}
      </div>
    {/if}
  </div>
  {#if !renamed}
    <span class="text-[10px] text-base-content/40 shrink-0">
      {file.confidence > 0 ? Math.round(file.confidence * 100) + '%' : '—'}
    </span>
  {/if}
</div>

<style>
  .file-selected {
    background: color-mix(in srgb, var(--color-primary, #661ae4) 15%, transparent) !important;
    outline: 1px solid var(--color-primary, #661ae4);
  }
</style>
