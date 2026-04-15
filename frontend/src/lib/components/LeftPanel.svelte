<script lang="ts">
  import { sourceFiles, sourceFilter, mediaType } from '$lib/stores';
  import { scan } from '$lib/api';

  interface FolderEntry {
    path: string;
    label: string;
    children: string[];
    expanded: boolean;
  }

  let folders: FolderEntry[] = $state([
    { path: '/media/downloads', label: '/media/downloads', children: ['Anime', 'Movies', 'TV Shows', 'Music'], expanded: true },
    { path: '/media/series',    label: '/media/series',    children: ['Completed', 'Watching'], expanded: false },
  ]);
  let selectedPath = $state<string | null>(null);
  let loading = $state(false);

  async function scanFolder(path: string) {
    selectedPath = path;
    loading = true;
    try {
      const files = await scan({ path, recursive: true, mediaType: $mediaType });
      sourceFiles.set(files);
    } catch (e) {
      console.error('scan failed', e);
    } finally {
      loading = false;
    }
  }

  function addFolder() {
    const path = prompt('Enter folder path:');
    if (!path) return;
    folders = [...folders, { path, label: path, children: [], expanded: false }];
  }

  function toggleFolder(f: FolderEntry) {
    f.expanded = !f.expanded;
    folders = [...folders];
  }
</script>

<aside class="w-[220px] bg-base-100 border-r border-base-300 flex flex-col shrink-0 overflow-hidden">
  <div class="flex items-center justify-between px-3 py-2 border-b border-base-300">
    <span class="text-[10px] font-bold uppercase tracking-widest text-base-content/50">Source Folders</span>
    <button class="text-base-content/50 hover:text-base-content text-sm leading-none" onclick={addFolder} title="Add folder">＋</button>
  </div>

  <div class="overflow-y-auto flex-1">
    {#each folders as folder}
      <!-- Root folder row -->
      <button
        class="flex items-center gap-1.5 w-full px-3 py-[5px] text-left text-xs font-semibold text-base-content hover:bg-base-200 cursor-pointer"
        onclick={() => toggleFolder(folder)}
      >
        <span>{folder.expanded ? '📂' : '📁'}</span>
        <span class="truncate">{folder.label}</span>
      </button>

      {#if folder.expanded}
        {#each folder.children as child}
          <button
            class="flex items-center gap-1.5 w-full pl-6 pr-3 py-1 text-left text-xs border-l-2 cursor-pointer {selectedPath === folder.path + '/' + child ? 'border-l-primary bg-base-200 text-primary' : 'border-l-transparent text-base-content/60 hover:bg-base-200 hover:text-base-content'}"
            onclick={() => scanFolder(folder.path + '/' + child)}
          >
            <span>📂</span>
            <span class="truncate">{child}</span>
            {#if loading && selectedPath === folder.path + '/' + child}
              <span class="loading loading-spinner loading-xs ml-auto"></span>
            {/if}
          </button>
        {/each}
      {/if}
    {/each}
  </div>
</aside>
