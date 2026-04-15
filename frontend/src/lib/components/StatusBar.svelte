<script lang="ts">
  import { fileStats } from '$lib/stores';

  // Source connection status — will be wired to real health checks in Phase 2
  const SOURCES = [
    { name: 'AniDB',       color: 'bg-success' },
    { name: 'TVDB',        color: 'bg-success' },
    { name: 'TMDB',        color: 'bg-success' },
    { name: 'IMDb',        color: 'bg-success' },
    { name: 'OpenSub',     color: 'bg-warning' },
    { name: 'MusicBrainz', color: 'bg-base-content/30' },
  ];
</script>

<footer class="flex items-center gap-3 px-3 h-[26px] bg-base-100 border-t border-base-300 text-[11px] text-base-content/50 shrink-0 overflow-x-auto">
  {#each SOURCES as src}
    <div class="flex items-center gap-1 shrink-0">
      <span class="w-1.5 h-1.5 rounded-full {src.color}"></span>
      <span>{src.name}</span>
    </div>
  {/each}

  <div class="w-px h-3 bg-base-300 shrink-0"></div>

  <div class="shrink-0">
    {$fileStats.total} files · {$fileStats.matched} matched · {$fileStats.pending} pending · {$fileStats.unmatched} unmatched
  </div>

  <div class="flex-1"></div>

  <div class="shrink-0 text-base-content/30">Port :48040</div>
</footer>
