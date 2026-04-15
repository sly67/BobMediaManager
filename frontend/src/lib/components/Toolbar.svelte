<script lang="ts">
  import { mediaType, enabledSources, idPickerOpen, idPickerTarget, selectedFile } from '$lib/stores';

  const BADGES = [
    { id: 'anidb',       label: 'AniDB',       cls: 'badge-anidb' },
    { id: 'tvdb',        label: 'TVDB',         cls: 'badge-tvdb' },
    { id: 'tmdb',        label: 'TMDB',         cls: 'badge-tmdb' },
    { id: 'imdb',        label: 'IMDb',         cls: 'badge-imdb' },
    { id: 'musicbrainz', label: 'MusicBrainz',  cls: 'badge-mb' },
  ];

  function toggleSource(id: string) {
    enabledSources.update(s => ({ ...s, [id]: !s[id] }));
  }

  function openPicker() {
    idPickerTarget.set($selectedFile);
    idPickerOpen.set(true);
  }
</script>

<div class="flex items-center gap-2 px-3 py-1.5 bg-base-100 border-b border-base-300 shrink-0 flex-wrap">
  <!-- Media type selector -->
  <select
    class="select select-bordered select-xs text-xs"
    value={$mediaType}
    onchange={(e) => mediaType.set((e.target as HTMLSelectElement).value)}
  >
    <option value="anime">🎌 Anime</option>
    <option value="series">📺 TV Series</option>
    <option value="movie">🎬 Movie</option>
    <option value="music">🎵 Music</option>
  </select>

  <!-- DB source badges -->
  {#each BADGES as badge}
    <button
      class="db-badge db-{badge.id} {$enabledSources[badge.id] ? '' : 'db-off'}"
      onclick={() => toggleSource(badge.id)}
      title="Toggle {badge.label}"
    >
      {badge.label}
    </button>
  {/each}

  <div class="w-px h-4 bg-base-300"></div>
  <button class="btn btn-ghost btn-xs text-xs">Auto-detect</button>
  <input
    class="input input-bordered input-xs flex-1 min-w-[120px] text-xs"
    placeholder="Override search query…"
  />
  <button class="btn btn-outline btn-xs text-xs" onclick={openPicker}>🔍 Pick ID</button>
</div>

<style>
  .db-badge {
    padding: 2px 8px;
    border-radius: 10px;
    font-size: 10px;
    font-weight: 700;
    cursor: pointer;
    user-select: none;
    border: 1px solid;
    transition: opacity 0.15s;
  }
  .db-badge.db-off { opacity: 0.3; }

  .db-anidb  { background: #1a0a2e; color: #a78bfa; border-color: #a78bfa; }
  .db-tvdb   { background: #0a1628; color: #4fa3e0; border-color: #4fa3e0; }
  .db-tmdb   { background: #0d253f; color: #01d277; border-color: #01d277; }
  .db-imdb   { background: #1a1400; color: #f5c518; border-color: #f5c518; }
  .db-musicbrainz { background: #0a1e1a; color: #ba478f; border-color: #ba478f; }
</style>
