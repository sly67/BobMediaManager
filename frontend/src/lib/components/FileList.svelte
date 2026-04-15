<script lang="ts">
  import FileRow from './FileRow.svelte';
  import { filteredSource, filteredRenamed, sourceFilter, renamedFilter, selection, selectFile, fileStats } from '$lib/stores';
  import type { MediaFile } from '$lib/api';

  let { mode }: { mode: 'source' | 'renamed' } = $props();

  const files = $derived(mode === 'source' ? $filteredSource : $filteredRenamed);
  const filter = mode === 'source' ? sourceFilter : renamedFilter;
  const filterVal = $derived(mode === 'source' ? $sourceFilter : $renamedFilter);

  const headerStats = $derived(
    mode === 'source'
      ? `${$fileStats.total} files`
      : `${$fileStats.matched} matched · ${$fileStats.pending} pending · ${$fileStats.unmatched} unmatched`
  );

  const statsColor = $derived(mode === 'renamed' ? 'text-success' : 'text-base-content');

  function handleSelect(e: MouseEvent, file: MediaFile, index: number) {
    selectFile(file, index, e.ctrlKey || e.metaKey, e.shiftKey, files);
  }

  // Sample demo data when no files are loaded
  const DEMO_SOURCE: MediaFile[] = [
    { id: '1', name: '[SubsPlease] Delicious in Dungeon - 01 [1080p].mkv', path: '/media/anime/1', size: 2.1e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched',   confidence: 0.98, resolvedIds: { anidb: '17243', tvdb: '431514', tmdb: '209867' } },
    { id: '2', name: '[SubsPlease] Delicious in Dungeon - 02 [1080p].mkv', path: '/media/anime/2', size: 2.2e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched',   confidence: 0.98, resolvedIds: { anidb: '17243' } },
    { id: '3', name: '[SubsPlease] Delicious in Dungeon - 03 [1080p].mkv', path: '/media/anime/3', size: 2.0e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched',   confidence: 0.97, resolvedIds: { anidb: '17243' } },
    { id: '4', name: '[SubsPlease] Dungeon.Meshi Season 1/',               path: '/media/anime/4', size: 0,    isFolder: true,  mediaType: 'anime', matchStatus: 'pending',   confidence: 0.71, resolvedIds: {} },
    { id: '5', name: 'unknown_anime_ep05_final.mkv',                       path: '/media/anime/5', size: 1.8e9, isFolder: false, mediaType: 'anime', matchStatus: 'unmatched', confidence: 0,    resolvedIds: {} },
    { id: '6', name: 'Sousou.no.Frieren.E01.1080p.mkv',                   path: '/media/anime/6', size: 1.5e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched',   confidence: 0.95, resolvedIds: { anidb: '17617' } },
    { id: '7', name: 'The.Bear.S03E01.WEB-DL.mkv',                        path: '/media/anime/7', size: 1.3e9, isFolder: false, mediaType: 'series',matchStatus: 'matched',   confidence: 0.99, resolvedIds: { tvdb: '401015' } },
    { id: '8', name: 'movie.2024.2160p.HDR.mkv',                          path: '/media/anime/8', size: 18e9, isFolder: false, mediaType: 'movie', matchStatus: 'pending',   confidence: 0.44, resolvedIds: {} },
  ];

  const DEMO_RENAMED: MediaFile[] = [
    { id: '1', name: 'Delicious in Dungeon - S01E01.mkv', path: '/media/anime/1', renamedPath: 'Delicious in Dungeon (2024)/Season 01/Delicious in Dungeon - S01E01 - Road to the Golden Country.mkv', size: 2.1e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched', confidence: 0.98, resolvedIds: { anidb: '17243' } },
    { id: '2', name: 'ep2.mkv', path: '/media/anime/2', renamedPath: 'Delicious in Dungeon (2024)/Season 01/Delicious in Dungeon - S01E02 - Omelet.mkv', size: 2.2e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched', confidence: 0.98, resolvedIds: { anidb: '17243' } },
    { id: '3', name: 'ep3.mkv', path: '/media/anime/3', renamedPath: 'Delicious in Dungeon (2024)/Season 01/Delicious in Dungeon - S01E03 - Kakiage.mkv', size: 2.0e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched', confidence: 0.97, resolvedIds: { anidb: '17243' } },
    { id: '4', name: 'folder/', path: '/media/anime/4', renamedPath: 'Multiple matches — click to pick', size: 0, isFolder: true, mediaType: 'anime', matchStatus: 'pending', confidence: 0.71, resolvedIds: {} },
    { id: '5', name: 'unknown.mkv', path: '/media/anime/5', renamedPath: 'No match — manual search', size: 1.8e9, isFolder: false, mediaType: 'anime', matchStatus: 'unmatched', confidence: 0, resolvedIds: {} },
    { id: '6', name: 'frieren.mkv', path: '/media/anime/6', renamedPath: "Frieren Beyond Journey's End (2023)/Season 01/Frieren - S01E01 - The Journey's End.mkv", size: 1.5e9, isFolder: false, mediaType: 'anime', matchStatus: 'matched', confidence: 0.95, resolvedIds: { anidb: '17617' } },
    { id: '7', name: 'bear.mkv', path: '/media/anime/7', renamedPath: 'The Bear (2022)/Season 3/The Bear - S03E01 - Tomorrow.mkv', size: 1.3e9, isFolder: false, mediaType: 'series', matchStatus: 'matched', confidence: 0.99, resolvedIds: { tvdb: '401015' } },
    { id: '8', name: 'movie.mkv', path: '/media/anime/8', renamedPath: 'Low confidence — confirm', size: 18e9, isFolder: false, mediaType: 'movie', matchStatus: 'pending', confidence: 0.44, resolvedIds: { tmdb: '693134' } },
  ];

  const displayFiles = $derived(files.length > 0 ? files : (mode === 'source' ? DEMO_SOURCE : DEMO_RENAMED));
</script>

<div class="flex flex-col flex-1 overflow-hidden {mode === 'source' ? 'border-r border-base-300' : ''}">
  <!-- Column header -->
  <div class="flex items-center justify-between px-3 py-[7px] bg-base-200 border-b border-base-300 text-[10px] font-bold uppercase tracking-widest text-base-content/50 shrink-0">
    <span>{mode === 'source' ? 'Source Files' : 'Renamed Preview'}</span>
    <span class={statsColor}>{headerStats}</span>
  </div>

  <!-- Search bar -->
  <div class="px-2 py-1 bg-base-100 border-b border-base-300 shrink-0">
    <input
      class="input input-xs w-full text-xs"
      placeholder="🔍  Filter {mode === 'source' ? 'files' : 'renamed paths'}…"
      value={filterVal}
      oninput={(e) => filter.set((e.target as HTMLInputElement).value)}
    />
  </div>

  <!-- File list -->
  <div class="overflow-y-auto flex-1">
    {#if displayFiles.length === 0}
      <div class="flex items-center justify-center h-full text-xs text-base-content/30 italic p-4 text-center">
        {mode === 'source' ? 'Scan a folder to begin' : 'Previews will appear here'}
      </div>
    {:else}
      {#each displayFiles as file, i (file.id)}
        <FileRow {file} index={i} renamed={mode === 'renamed'} onselect={handleSelect} />
      {/each}
    {/if}
  </div>
</div>
