<script lang="ts">
  import { sourcesModalOpen } from '$lib/stores';

  const BUILTIN = [
    { icon: '🎌', name: 'AniDB',        url: 'api.anidb.net · UDP API',        types: ['Anime'],            enabled: true },
    { icon: '📺', name: 'TheTVDB',      url: 'api4.thetvdb.com/v4',            types: ['TV Series','Anime'], enabled: true },
    { icon: '🎬', name: 'TMDB',         url: 'api.themoviedb.org/3',           types: ['Movie','TV Series','Anime'], enabled: true },
    { icon: '⭐', name: 'IMDb',         url: 'imdb.com · scrape / Cinemagoer', types: ['Movie','TV Series'], enabled: true },
    { icon: '🎵', name: 'MusicBrainz',  url: 'musicbrainz.org/ws/2',           types: ['Music'],             enabled: false },
    { icon: '🎵', name: 'Last.fm',      url: 'ws.audioscrobbler.com/2.0',      types: ['Music'],             enabled: false },
    { icon: '💬', name: 'OpenSubtitles',url: 'api.opensubtitles.com/api/v1',   types: ['All'],               enabled: true },
  ];

  interface CustomSrc { icon: string; name: string; url: string; types: string[]; enabled: boolean; }
  let custom = $state<CustomSrc[]>([
    { icon: '🌐', name: 'MyAnimeList (Jikan)', url: 'api.jikan.moe/v4', types: ['Anime'], enabled: true },
  ]);

  let builtin = $state(BUILTIN.map(s => ({ ...s })));

  // Add form
  let newName = $state('');
  let newUrl = $state('');
  let newKey = $state('');
  let newTypes = $state<Set<string>>(new Set());

  const TYPE_OPTS = ['Movie', 'TV Series', 'Anime', 'Music'];
  const TYPE_COLORS: Record<string, string> = {
    'Movie': '#c084fc', 'TV Series': '#60a5fa', 'Anime': '#a78bfa', 'Music': '#34d399', 'All': '#94a3b8',
  };

  function toggleBuiltin(i: number) {
    builtin[i].enabled = !builtin[i].enabled;
  }

  function removeCustom(i: number) {
    custom = custom.filter((_, idx) => idx !== i);
  }

  function addCustom() {
    if (!newName || !newUrl) return;
    custom = [...custom, { icon: '🌐', name: newName, url: newUrl, types: [...newTypes], enabled: true }];
    newName = ''; newUrl = ''; newKey = ''; newTypes = new Set();
  }

  function closeOutside(e: MouseEvent) {
    if ((e.target as HTMLElement).classList.contains('modal-backdrop')) {
      sourcesModalOpen.set(false);
    }
  }
</script>

{#if $sourcesModalOpen}
  <!-- svelte-ignore a11y_click_events_have_key_events -->
  <!-- svelte-ignore a11y_no_static_element_interactions -->
  <div
    class="modal-backdrop fixed inset-0 flex items-center justify-center z-50"
    style="background: var(--overlay, rgba(0,0,0,.72)); backdrop-filter:blur(4px);"
    onclick={closeOutside}
  >
    <div class="bg-base-100 border border-base-300 rounded-xl shadow-2xl flex flex-col overflow-hidden" style="width:720px; max-height:85vh;">

      <!-- Header -->
      <div class="flex items-start gap-2.5 px-4 py-3 border-b border-base-300 shrink-0">
        <div class="flex-1">
          <div class="text-sm font-bold">Metadata Sources</div>
          <div class="text-[11px] text-base-content/50 mt-0.5">Manage built-in and custom API sources. Specify which media types each source applies to.</div>
        </div>
        <button class="text-base-content/50 hover:text-base-content text-lg leading-none" onclick={() => sourcesModalOpen.set(false)}>✕</button>
      </div>

      <!-- Source list -->
      <div class="overflow-y-auto flex-1">
        <div class="px-4 py-1.5 text-[10px] font-bold uppercase tracking-widest text-base-content/50">Built-in</div>

        {#each builtin as src, i}
          <div class="flex items-center gap-2.5 px-4 py-2.5 border-b border-base-300 hover:bg-base-200">
            <div class="w-9 h-9 rounded-lg bg-base-200 border border-base-300 flex items-center justify-center text-lg shrink-0">{src.icon}</div>
            <div class="flex-1 min-w-0">
              <div class="text-sm font-semibold">{src.name}</div>
              <div class="text-[10px] text-base-content/40 font-mono mt-0.5">{src.url}</div>
              <div class="flex gap-1 mt-1 flex-wrap">
                {#each src.types as t}
                  <span class="px-1.5 py-0.5 rounded-full text-[9px] font-bold border" style="color:{TYPE_COLORS[t] ?? '#888'}; border-color:{TYPE_COLORS[t] ?? '#888'}; background: color-mix(in srgb, {TYPE_COLORS[t] ?? '#888'} 10%, transparent)">
                    {t}
                  </span>
                {/each}
              </div>
            </div>
            <div class="flex items-center gap-2 shrink-0">
              <!-- Toggle switch -->
              <button
                class="relative w-9 h-5 rounded-full cursor-pointer border-none transition-colors {src.enabled ? 'bg-success' : 'bg-base-300'}"
                onclick={() => toggleBuiltin(i)}
              >
                <span class="absolute top-0.5 w-4 h-4 rounded-full bg-white shadow transition-all {src.enabled ? 'left-[18px]' : 'left-0.5'}"></span>
              </button>
              <button class="btn btn-outline btn-xs">Configure</button>
            </div>
          </div>
        {/each}

        <div class="px-4 py-1.5 text-[10px] font-bold uppercase tracking-widest text-base-content/50 mt-1">Custom Sources</div>

        {#each custom as src, i}
          <div class="flex items-center gap-2.5 px-4 py-2.5 border-b border-base-300 hover:bg-base-200">
            <div class="w-9 h-9 rounded-lg bg-base-200 border border-base-300 flex items-center justify-center text-lg shrink-0">{src.icon}</div>
            <div class="flex-1 min-w-0">
              <div class="text-sm font-semibold">{src.name}</div>
              <div class="text-[10px] text-base-content/40 font-mono mt-0.5">{src.url}</div>
              <div class="flex gap-1 mt-1">
                {#each src.types as t}
                  <span class="px-1.5 py-0.5 rounded-full text-[9px] font-bold border" style="color:{TYPE_COLORS[t] ?? '#888'}; border-color:{TYPE_COLORS[t] ?? '#888'}">
                    {t}
                  </span>
                {/each}
              </div>
            </div>
            <div class="flex items-center gap-2 shrink-0">
              <button
                class="relative w-9 h-5 rounded-full cursor-pointer border-none transition-colors {src.enabled ? 'bg-success' : 'bg-base-300'}"
                onclick={() => { custom[i].enabled = !custom[i].enabled; custom = [...custom]; }}
              >
                <span class="absolute top-0.5 w-4 h-4 rounded-full bg-white shadow transition-all {src.enabled ? 'left-[18px]' : 'left-0.5'}"></span>
              </button>
              <button class="btn btn-outline btn-xs">Edit</button>
              <button class="btn btn-outline btn-xs text-error" onclick={() => removeCustom(i)}>✕</button>
            </div>
          </div>
        {/each}
      </div>

      <!-- Add form -->
      <div class="px-4 py-3 border-t border-base-300 bg-base-200 shrink-0">
        <div class="text-[11px] font-bold uppercase tracking-widest text-base-content/50 mb-2">Add Custom Source</div>
        <div class="grid grid-cols-2 gap-2 mb-2">
          <div class="flex items-center gap-2">
            <span class="text-[11px] text-base-content/50 w-16 shrink-0">Name</span>
            <input class="input input-bordered input-xs flex-1" placeholder="e.g. Kitsu" bind:value={newName} />
          </div>
          <div class="flex items-center gap-2">
            <span class="text-[11px] text-base-content/50 w-16 shrink-0">Base URL</span>
            <input class="input input-bordered input-xs flex-1" placeholder="https://kitsu.io/api/edge" bind:value={newUrl} />
          </div>
          <div class="flex items-center gap-2">
            <span class="text-[11px] text-base-content/50 w-16 shrink-0">API Key</span>
            <input class="input input-bordered input-xs flex-1" placeholder="Optional" bind:value={newKey} />
          </div>
          <div class="flex items-center gap-2">
            <span class="text-[11px] text-base-content/50 w-16 shrink-0">Types</span>
            <div class="flex gap-1 flex-wrap">
              {#each TYPE_OPTS as t}
                <label class="flex items-center gap-0.5 text-[10px] cursor-pointer">
                  <input
                    type="checkbox"
                    class="checkbox checkbox-xs"
                    checked={newTypes.has(t)}
                    onchange={() => {
                      if (newTypes.has(t)) { newTypes.delete(t); } else { newTypes.add(t); }
                      newTypes = new Set(newTypes);
                    }}
                  />
                  {t}
                </label>
              {/each}
            </div>
          </div>
        </div>
        <div class="flex justify-end">
          <button class="btn btn-primary btn-xs" onclick={addCustom}>+ Add Source</button>
        </div>
      </div>
    </div>
  </div>
{/if}
