<script lang="ts">
  import { renameOp, destination, renameFolders } from '$lib/stores';

  const OPS: { id: 'move' | 'copy' | 'rename'; label: string }[] = [
    { id: 'move',   label: 'Move' },
    { id: 'copy',   label: 'Copy' },
    { id: 'rename', label: 'Rename in-place' },
  ];
</script>

<div class="flex items-center gap-2 px-3 py-1.5 bg-base-200 border-b border-base-300 shrink-0 flex-wrap">
  <span class="text-[10px] font-bold uppercase tracking-wide text-base-content/50 shrink-0">Operation</span>

  <!-- Segmented control -->
  <div class="join">
    {#each OPS as op}
      <button
        class="btn btn-xs join-item {$renameOp === op.id ? 'btn-primary' : 'btn-outline'}"
        onclick={() => renameOp.set(op.id)}
      >
        {op.label}
      </button>
    {/each}
  </div>

  <span class="text-[10px] font-bold uppercase tracking-wide text-base-content/50 shrink-0">
    Destination <span class="text-error text-sm leading-none">*</span>
  </span>

  <input
    class="input input-bordered input-xs flex-1 min-w-[200px] font-mono text-xs {$renameOp === 'rename' ? 'opacity-40 pointer-events-none' : $destination === '' ? 'border-error' : ''}"
    placeholder="/media/library/anime  (required)"
    value={$destination}
    oninput={(e) => destination.set((e.target as HTMLInputElement).value)}
    disabled={$renameOp === 'rename'}
  />
  <button
    class="btn btn-outline btn-xs shrink-0"
    disabled={$renameOp === 'rename'}
  >Browse</button>

  <div class="w-px h-4 bg-base-300"></div>
  <label class="flex items-center gap-1.5 text-xs cursor-pointer shrink-0 whitespace-nowrap">
    <input
      type="checkbox"
      class="checkbox checkbox-xs"
      checked={$renameFolders}
      onchange={(e) => renameFolders.set((e.target as HTMLInputElement).checked)}
    />
    Rename folders too
  </label>
</div>
