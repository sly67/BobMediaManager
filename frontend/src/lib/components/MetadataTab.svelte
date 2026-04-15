<script lang="ts">
  import { selectedFile, idPickerOpen, idPickerTarget } from '$lib/stores';

  // Demo metadata when a file is selected
  const DEMO_META = {
    title: 'Delicious in Dungeon',
    origTitle: 'ダンジョン飯',
    year: 2024,
    status: 'Completed · 24 eps',
    studio: 'Trigger',
    ids: { anidb: '17243', tvdb: '431514', tmdb: '209867', imdb: 'tt14491254' },
    episode: { num: 1, title: 'Road to the Golden Country', aired: '2024-01-04', runtime: '24 min', rating: 9.1 },
    genres: ['Adventure', 'Comedy', 'Fantasy', 'Cooking', 'Dungeon'],
  };

  const idColors: Record<string, string> = {
    anidb: '#a78bfa', tvdb: '#4fa3e0', tmdb: '#01d277', imdb: '#f5c518',
  };

  function openPicker() {
    idPickerTarget.set($selectedFile);
    idPickerOpen.set(true);
  }
</script>

<div class="overflow-y-auto flex-1">
  <!-- Poster / banner area -->
  <div class="w-full bg-base-200 border-b border-base-300 flex items-center justify-center relative overflow-hidden" style="aspect-ratio:16/9">
    <div class="text-center text-base-content/40">
      <div class="text-5xl mb-1">🎌</div>
      <div class="text-[11px]">Delicious in Dungeon</div>
    </div>
    <div class="absolute bottom-0 left-0 right-0 px-2.5 py-2 bg-gradient-to-t from-black/85 to-transparent text-xs font-semibold text-white">
      Delicious in Dungeon · 2024
    </div>
  </div>

  <!-- Series info -->
  <div class="px-3 py-2 border-b border-base-300">
    <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Series</div>
    {#each [
      ['Title',  DEMO_META.title,       true],
      ['Orig.',  DEMO_META.origTitle,   false],
      ['Year',   DEMO_META.year,        false],
      ['Status', DEMO_META.status,      false],
      ['Studio', DEMO_META.studio,      false],
    ] as [k, v, accent]}
      <div class="flex gap-1.5 mb-1">
        <span class="text-[11px] text-base-content/50 w-[72px] shrink-0">{k}</span>
        <span class="text-[11px] {accent ? 'text-primary' : 'text-base-content'} flex-1">{v}</span>
      </div>
    {/each}
    <!-- IDs row -->
    <div class="flex gap-1.5 mt-1">
      <span class="text-[11px] text-base-content/50 w-[72px] shrink-0">IDs</span>
      <div class="flex flex-wrap gap-1 flex-1">
        {#each Object.entries(DEMO_META.ids) as [src, id]}
          <span class="px-1.5 py-0.5 rounded-lg bg-base-200 border border-base-300 text-[10px]" style="color:{idColors[src] ?? 'inherit'}">
            {src.toUpperCase()}:{id}
          </span>
        {/each}
      </div>
    </div>
  </div>

  <!-- Episode info -->
  <div class="px-3 py-2 border-b border-base-300">
    <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Episode {DEMO_META.episode.num}</div>
    {#each [
      ['Title',   DEMO_META.episode.title],
      ['Aired',   DEMO_META.episode.aired],
      ['Runtime', DEMO_META.episode.runtime],
    ] as [k, v]}
      <div class="flex gap-1.5 mb-1">
        <span class="text-[11px] text-base-content/50 w-[72px] shrink-0">{k}</span>
        <span class="text-[11px] text-base-content">{v}</span>
      </div>
    {/each}
    <!-- Rating bar -->
    <div class="flex items-center gap-1.5">
      <span class="text-[11px] text-base-content/50 w-[72px] shrink-0">Rating</span>
      <div class="h-1 bg-base-300 rounded flex-1 overflow-hidden">
        <div class="h-full bg-warning rounded" style="width:{DEMO_META.episode.rating * 10}%"></div>
      </div>
      <span class="text-[10px] text-warning">{DEMO_META.episode.rating}</span>
    </div>
  </div>

  <!-- Genres -->
  <div class="px-3 py-2 border-b border-base-300">
    <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Genres</div>
    <div class="flex flex-wrap gap-1">
      {#each DEMO_META.genres as g}
        <span class="px-1.5 py-0.5 rounded-full bg-base-200 border border-base-300 text-[10px] text-base-content/60">{g}</span>
      {/each}
    </div>
  </div>

  <!-- Actions -->
  <div class="px-3 py-2">
    <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Actions</div>
    <div class="flex flex-col gap-1">
      <button class="btn btn-ghost btn-xs justify-start text-[11px]">💬 Grab Subtitles</button>
      <button class="btn btn-ghost btn-xs justify-start text-[11px]">🔐 Verify File Signatures</button>
      <button class="btn btn-ghost btn-xs justify-start text-[11px]">🔁 Re-match Episode</button>
      <button class="btn btn-ghost btn-xs justify-start text-[11px] text-primary" onclick={openPicker}>
        🔍 Override ID (e.g. TMDB:xxxx)
      </button>
    </div>
  </div>
</div>
