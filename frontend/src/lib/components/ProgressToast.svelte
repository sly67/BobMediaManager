<script lang="ts">
  import { activeProgress } from '$lib/stores';

  const pct = $derived(
    $activeProgress && $activeProgress.total > 0
      ? Math.round(($activeProgress.done / $activeProgress.total) * 100)
      : 0
  );
</script>

{#if $activeProgress}
  <div class="fixed bottom-8 left-1/2 -translate-x-1/2 z-[400] bg-base-100 border border-base-300 rounded-xl shadow-2xl px-5 py-4 min-w-[360px] max-w-[500px]">
    <div class="flex items-center gap-3 mb-2">
      <span class="loading loading-spinner loading-sm text-primary"></span>
      <div class="flex-1">
        <div class="text-sm font-semibold capitalize">{$activeProgress.operation}…</div>
        <div class="text-[11px] text-base-content/50 mt-0.5">
          {$activeProgress.done} / {$activeProgress.total} files
        </div>
      </div>
      <span class="text-sm font-bold text-primary">{pct}%</span>
    </div>

    <!-- Progress bar -->
    <div class="h-1.5 bg-base-300 rounded-full overflow-hidden">
      <div
        class="h-full bg-primary rounded-full transition-[width] duration-300"
        style="width:{pct}%"
      ></div>
    </div>

    {#if $activeProgress.error}
      <div class="mt-2 text-[11px] text-error">{$activeProgress.error}</div>
    {/if}
  </div>
{/if}
