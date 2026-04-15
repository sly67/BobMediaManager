<script lang="ts">
  let editing = $state(false);

  const DEMO_NFO = `<episodedetails>
  <title>Road to the Golden Country</title>
  <showtitle>Delicious in Dungeon</showtitle>
  <season>1</season>
  <episode>1</episode>
  <aired>2024-01-04</aired>
  <rating>9.1</rating>
  <plot>When Laios and his party are forced to flee...</plot>
  <anidbid>17243</anidbid>
  <tmdbid>209867</tmdbid>
  <tvdbid>431514</tvdbid>
  <imdbid>tt14491254</imdbid>
</episodedetails>`;

  let nfoContent = $state(DEMO_NFO);

  // Simple XML syntax highlighting
  function highlight(xml: string): string {
    return xml
      .replace(/&/g, '&amp;')
      .replace(/</g, '&lt;')
      .replace(/>/g, '&gt;')
      .replace(/(&lt;\/?[\w]+&gt;)/g, '<span class="nfo-tag">$1</span>')
      .replace(/(&gt;)([^&<]+)(&lt;)/g, '$1<span class="nfo-val">$2</span>$3');
  }
</script>

<div class="overflow-y-auto flex-1">
  <div class="px-3 py-2 border-b border-base-300">
    <div class="flex items-center justify-between mb-1.5">
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50">NFO Template</div>
      <button class="btn btn-outline btn-xs text-[10px]" onclick={() => editing = !editing}>
        {editing ? 'Preview' : 'Edit'}
      </button>
    </div>

    {#if editing}
      <textarea
        class="w-full bg-base-200 border border-base-300 rounded-md p-2 font-mono text-[10px] text-base-content/80 leading-relaxed resize-none focus:outline-primary"
        rows="16"
        bind:value={nfoContent}
      ></textarea>
    {:else}
      <!-- Syntax-highlighted preview -->
      <div class="bg-base-200 border border-base-300 rounded-md p-2 font-mono text-[10px] leading-relaxed overflow-y-auto max-h-64">
        <!-- eslint-disable-next-line svelte/no-at-html-tags -->
        {@html highlight(nfoContent)}
      </div>
    {/if}
  </div>

  <div class="px-3 py-2">
    <button class="btn btn-primary btn-xs w-full text-[11px]">📄 Generate NFO Files</button>
  </div>
</div>

<style>
  :global(.nfo-tag) { color: var(--color-primary, #661ae4); }
  :global(.nfo-val) { color: var(--color-success, #36d399); }
</style>
