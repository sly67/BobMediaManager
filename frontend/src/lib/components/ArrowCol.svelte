<script lang="ts">
  import { filteredSource, filteredRenamed } from '$lib/stores';

  // Show an arrow (→) or question mark (?) aligned with each source row.
  // Row height is 38px (7px padding × 2 + content).
  const ROW_H = 38;
  const HEADER_H = 34; // col-hdr + search bar

  const arrows = $derived(
    ($filteredSource.length > 0 ? $filteredSource : new Array(8).fill(null)).map((f, i) => {
      if (!f) return '→';
      return f.matchStatus === 'matched' ? '→' : f.matchStatus === 'pending' ? '?' : '✕';
    })
  );
</script>

<div class="flex flex-col items-center w-7 bg-base-300/40 border-x border-base-300 shrink-0 overflow-hidden">
  <!-- Spacer for header area (col-hdr + search bar height) -->
  <div style="height:{HEADER_H}px" class="shrink-0"></div>
  {#each arrows as arrow}
    <span
      class="text-[11px] text-base-content/40 shrink-0 flex items-center justify-center"
      style="height:{ROW_H}px; width:28px;"
    >
      {arrow}
    </span>
  {/each}
</div>
