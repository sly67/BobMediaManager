<script lang="ts">
  import { ctxMenu, closeCtxMenu, idPickerOpen, idPickerTarget, selectAllMatched } from '$lib/stores';
  import { onMount } from 'svelte';

  let menuEl: HTMLDivElement;

  onMount(() => {
    function onDoc(e: MouseEvent) {
      if (menuEl && !menuEl.contains(e.target as Node)) {
        closeCtxMenu();
      }
    }
    document.addEventListener('click', onDoc);
    return () => document.removeEventListener('click', onDoc);
  });

  function action(fn: () => void) {
    fn();
    closeCtxMenu();
  }

  function openPicker() {
    idPickerTarget.set($ctxMenu.file);
    idPickerOpen.set(true);
    closeCtxMenu();
  }
</script>

{#if $ctxMenu.open}
  <div
    bind:this={menuEl}
    class="fixed bg-base-100 border border-base-300 rounded-lg min-w-[210px] shadow-2xl z-[500] py-1"
    style="left:{$ctxMenu.x}px; top:{$ctxMenu.y}px;"
  >
    <div class="px-3 py-1 text-[10px] font-bold uppercase tracking-widest text-base-content/40">File / Folder</div>
    <button class="ctx-item w-full text-left" onclick={openPicker}>🔍 Search / Override ID…</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>🔁 Re-match</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>✏️ Edit Renamed Path</button>

    <div class="h-px bg-base-300 my-1"></div>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>📄 Generate NFO</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>📥 Download Images</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>💬 Grab Subtitles</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>🔐 Verify Signature</button>

    <div class="h-px bg-base-300 my-1"></div>
    <div class="px-3 py-1 text-[10px] font-bold uppercase tracking-widest text-base-content/40">Batch</div>
    <button class="ctx-item w-full text-left" onclick={() => action(selectAllMatched)}>✅ Select All Matched</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => {})}>📦 Execute Selected Only</button>

    <div class="h-px bg-base-300 my-1"></div>
    <button class="ctx-item w-full text-left" onclick={() => action(() => navigator.clipboard.writeText($ctxMenu.file?.path ?? ''))}>📋 Copy Source Path</button>
    <button class="ctx-item w-full text-left" onclick={() => action(() => navigator.clipboard.writeText($ctxMenu.file?.renamedPath ?? ''))}>📋 Copy Destination Path</button>

    <div class="h-px bg-base-300 my-1"></div>
    <button class="ctx-item w-full text-left text-error" onclick={() => action(() => {})}>🗑️ Remove from List</button>
  </div>
{/if}

<style>
  .ctx-item {
    padding: 7px 14px;
    font-size: 12px;
    cursor: pointer;
    display: flex;
    align-items: center;
    gap: 8px;
    transition: background 0.1s;
  }
  .ctx-item:hover {
    background: oklch(var(--b2, 0.2 0 0));
  }
</style>
