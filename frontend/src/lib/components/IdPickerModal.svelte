<script lang="ts">
  import { idPickerOpen, idPickerTarget, searchResults, searchQuery, pickedResult } from '$lib/stores';
  import { search } from '$lib/api';
  import type { SearchResult } from '$lib/api';

  const DB_TABS = [
    { id: 'anidb',       label: 'AniDB',       cls: 'bg-[#1a0a2e] text-[#a78bfa] border-[#a78bfa]' },
    { id: 'tvdb',        label: 'TVDB',         cls: 'bg-[#0a1628] text-[#4fa3e0] border-[#4fa3e0]' },
    { id: 'tmdb',        label: 'TMDB',         cls: 'bg-[#0d253f] text-[#01d277] border-[#01d277]' },
    { id: 'imdb',        label: 'IMDb',         cls: 'bg-[#1a1400] text-[#f5c518] border-[#f5c518]' },
    { id: 'musicbrainz', label: 'MusicBrainz',  cls: 'bg-[#0a1e1a] text-[#ba478f] border-[#ba478f]' },
  ];

  const ID_COLORS: Record<string, string> = {
    anidb: '#a78bfa', tvdb: '#4fa3e0', tmdb: '#01d277', imdb: '#f5c518',
  };

  let activeDBs = $state(new Set(['anidb', 'tvdb', 'tmdb', 'imdb']));
  let manualId = $state('');
  let loading = $state(false);
  let localQuery = $state($idPickerTarget?.name ?? '');

  // Demo results
  const DEMO_RESULTS: SearchResult[] = [
    { id: '17243', sourceSlug: 'anidb', title: 'Delicious in Dungeon', year: 2024, overview: 'When Laios and his party are forced to flee the dungeon after a dragon attack, they must resort to eating the monsters they defeat to survive…', posterUrl: '', ids: { anidb: '17243', tvdb: '431514', tmdb: '209867', imdb: 'tt14491254' }, confidence: 0.98 },
    { id: '11757', sourceSlug: 'anidb', title: 'Dungeon Meshi', year: 2024, overview: 'Alternate title entry — same series in AniDB old system', posterUrl: '', ids: { anidb: '11757', tmdb: '209867' }, confidence: 0.71 },
    { id: '18301', sourceSlug: 'anidb', title: 'Delicious in Dungeon: Side Stories', year: 2024, overview: 'OVA / Special · 3 eps · Bonus content', posterUrl: '', ids: { anidb: '18301' }, confidence: 0.34 },
    { id: '921636', sourceSlug: 'tmdb', title: 'Dungeon', year: 2023, overview: 'Movie · Unrelated title match', posterUrl: '', ids: { tmdb: '921636', imdb: 'tt12345678' }, confidence: 0.12 },
  ];

  let results = $state<SearchResult[]>(DEMO_RESULTS);
  let selected = $state<SearchResult | null>(DEMO_RESULTS[0]);

  async function doSearch() {
    loading = true;
    try {
      const res = await search({ query: localQuery, sources: [...activeDBs], mediaType: 'anime' });
      results = res.length > 0 ? res : DEMO_RESULTS;
    } catch {
      results = DEMO_RESULTS;
    } finally {
      loading = false;
    }
  }

  function confirm() {
    pickedResult.set(selected);
    idPickerOpen.set(false);
  }

  function closeOutside(e: MouseEvent) {
    if ((e.target as HTMLElement).classList.contains('modal-backdrop')) {
      idPickerOpen.set(false);
    }
  }

  const selectedDisplay = $derived(
    selected
      ? Object.entries(selected.ids ?? {}).map(([k, v]) => `${k.toUpperCase()}:${v}`).join(' · ') + ` — ${selected.title} (${selected.year})`
      : manualId || 'None selected'
  );
</script>

{#if $idPickerOpen}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="modal-backdrop fixed inset-0 flex items-center justify-center z-50"
    style="background: var(--overlay, rgba(0,0,0,.72)); backdrop-filter:blur(4px);"
    onclick={closeOutside}
  >
    <div class="bg-base-100 border border-base-300 rounded-xl shadow-2xl flex flex-col overflow-hidden" style="width:min(92vw,1080px); height:min(88vh,700px);">

      <!-- Header -->
      <div class="flex items-start gap-2.5 px-4 py-3 border-b border-base-300 shrink-0">
        <div class="flex-1">
          <div class="text-sm font-bold">Select / Override Media ID</div>
          <div class="text-[11px] text-base-content/50 mt-0.5">
            File: <strong class="text-primary">{$idPickerTarget?.name ?? 'No file selected'}</strong>
            &nbsp;·&nbsp; Type directly: <code class="text-secondary">TMDB:12345</code> or <code class="text-secondary">ANIDB:17243</code>
          </div>
        </div>
        <button class="text-base-content/50 hover:text-base-content text-lg leading-none mt-0.5" onclick={() => idPickerOpen.set(false)}>✕</button>
      </div>

      <!-- Search bar -->
      <div class="flex items-center gap-2 px-4 py-2 border-b border-base-300 shrink-0 flex-wrap">
        <div class="flex gap-1">
          {#each DB_TABS as db}
            <button
              class="px-2 py-1 rounded-md border text-[10px] font-bold cursor-pointer {activeDBs.has(db.id) ? db.cls + ' border' : 'border-base-300 text-base-content/40 bg-base-200'}"
              onclick={() => {
                if (activeDBs.has(db.id)) {
                  activeDBs = new Set([...activeDBs].filter(x => x !== db.id));
                } else {
                  activeDBs = new Set([...activeDBs, db.id]);
                }
              }}
            >{db.label}</button>
          {/each}
        </div>
        <input
          class="input input-bordered input-sm font-mono text-xs w-36"
          placeholder="ANIDB:17243"
          bind:value={manualId}
        />
        <input
          class="input input-bordered input-sm flex-1 text-xs min-w-[180px]"
          placeholder="Search title…"
          bind:value={localQuery}
          onkeydown={(e) => { if (e.key === 'Enter') doSearch(); }}
        />
        <button class="btn btn-primary btn-sm text-xs" onclick={doSearch}>
          {#if loading}<span class="loading loading-spinner loading-xs"></span>{:else}Search{/if}
        </button>
      </div>

      <!-- Results -->
      <div class="overflow-y-auto flex-1">
        {#each results as r}
          <!-- svelte-ignore a11y_click_events_have_key_events -->
          <!-- svelte-ignore a11y_no_static_element_interactions -->
          <div
            class="flex items-center gap-2.5 px-4 py-2.5 border-b border-base-300 cursor-pointer {selected?.id === r.id ? 'bg-base-200 outline outline-1 outline-primary' : 'hover:bg-base-200'}"
            onclick={() => selected = r}
          >
            <!-- Poster placeholder -->
            <div class="w-[52px] h-[76px] bg-base-200 border border-base-300 rounded flex items-center justify-center text-2xl shrink-0">
              🎌
            </div>

            <!-- Info -->
            <div class="flex-1 min-w-0">
              <div class="text-sm font-semibold">
                {r.title}
                <span class="text-base-content/50 font-normal text-xs">({r.year})</span>
              </div>
              <div class="text-[11px] text-base-content/50 mt-0.5 line-clamp-1">{r.overview}</div>
              <div class="text-[10px] text-base-content/40 mt-0.5 line-clamp-2">{r.overview}</div>
              <div class="flex flex-wrap gap-1 mt-1">
                {#each Object.entries(r.ids ?? {}) as [src, id]}
                  <span class="px-1.5 py-0.5 rounded-full bg-base-200 border border-base-300 text-[10px]" style="color:{ID_COLORS[src] ?? 'inherit'}">
                    {src.toUpperCase()}:{id}
                  </span>
                {/each}
              </div>
            </div>

            <!-- Confidence -->
            <div class="text-[11px] font-bold w-9 text-right shrink-0 {r.confidence > 0.9 ? 'text-success' : r.confidence > 0.5 ? 'text-warning' : 'text-error'}">
              {Math.round(r.confidence * 100)}%
            </div>

            <!-- Link-back buttons -->
            <div class="flex flex-col gap-1 shrink-0">
              {#each Object.keys(r.ids ?? {}) as src}
                <a
                  class="text-[10px] text-secondary border border-base-300 rounded px-2 py-1 whitespace-nowrap hover:bg-base-200 text-center block"
                  href={src === 'anidb' ? `https://anidb.net/anime/${r.ids?.[src]}` : src === 'tvdb' ? `https://thetvdb.com/?id=${r.ids?.[src]}&tab=series` : src === 'tmdb' ? `https://www.themoviedb.org/tv/${r.ids?.[src]}` : `https://www.imdb.com/title/${r.ids?.[src]}`}
                  target="_blank"
                  rel="noopener noreferrer"
                  onclick={(e) => e.stopPropagation()}
                >↗ {src.toUpperCase()}</a>
              {/each}
            </div>
          </div>
        {/each}
      </div>

      <!-- Footer -->
      <div class="flex items-center gap-2.5 px-5 py-3 border-t border-base-300 shrink-0">
        <span class="text-[11px] text-base-content/50 shrink-0">Selected:</span>
        <div class="flex-1 font-mono text-[11px] px-2.5 py-1.5 rounded-md bg-base-200 border border-base-300 text-success truncate">
          {selectedDisplay}
        </div>
        <button class="btn btn-ghost btn-sm" onclick={() => idPickerOpen.set(false)}>Cancel</button>
        <button class="btn btn-success btn-sm" onclick={confirm}>✓ Confirm</button>
      </div>
    </div>
  </div>
{/if}
