<script lang="ts">
  const IMG_TYPES = ['Poster', 'Fanart', 'ExtraFanart', 'Thumb', 'Banner', 'Logo', 'ClearArt', 'Disc', 'CharArt'];
  let activeTypes = $state(new Set(['Poster', 'Fanart', 'ExtraFanart', 'Thumb']));
  let maxExtrafanart = $state(5);
  let selectedImages = $state(new Set<string>());

  function toggleType(t: string) {
    activeTypes = activeTypes.has(t)
      ? new Set([...activeTypes].filter(x => x !== t))
      : new Set([...activeTypes, t]);
  }

  function toggleImg(id: string) {
    selectedImages = selectedImages.has(id)
      ? new Set([...selectedImages].filter(x => x !== id))
      : new Set([...selectedImages, id]);
  }

  const DEMO_POSTERS  = ['2000×3000', '1400×2100', '680×1000'];
  const DEMO_FANARTS  = ['1920×1080', '1920×1080', '1280×720', '1920×1080'];
  const DEMO_EXTRAFAN = ['ef/1', 'ef/2', 'ef/3', 'ef/4', 'ef/5', 'ef/6'];
  const DEMO_THUMBS   = ['thumb1', 'thumb2'];
</script>

<div class="overflow-y-auto flex-1">
  <!-- Type selector -->
  <div class="px-3 py-2 border-b border-base-300">
    <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Image Types</div>
    <div class="flex flex-wrap gap-1">
      {#each IMG_TYPES as t}
        <button
          class="px-2 py-0.5 rounded-full border text-[10px] cursor-pointer {activeTypes.has(t) ? 'border-primary text-primary' : 'border-base-300 text-base-content/50'}"
          onclick={() => toggleType(t)}
        >{t}</button>
      {/each}
    </div>
  </div>

  <!-- Poster grid -->
  {#if activeTypes.has('Poster')}
    <div class="px-3 py-2 border-b border-base-300">
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">
        Poster <span class="text-base-content/40 font-normal normal-case tracking-normal">{DEMO_POSTERS.length} available</span>
      </div>
      <div class="grid gap-1" style="grid-template-columns:repeat(3,1fr)">
        {#each DEMO_POSTERS as dim, i}
          {@const id = 'poster-' + i}
          <button
            class="relative bg-base-200 rounded border cursor-pointer flex items-center justify-center text-base-content/30 text-xs overflow-hidden {selectedImages.has(id) ? 'border-primary' : 'border-base-300'}"
            style="aspect-ratio:2/3"
            onclick={() => toggleImg(id)}
          >
            🎌
            <span class="absolute bottom-0.5 left-1 text-[8px] text-white/70">{dim}</span>
            {#if selectedImages.has(id)}
              <span class="absolute bottom-0.5 right-0.5 bg-primary text-primary-content text-[8px] px-1 rounded">✓</span>
            {/if}
          </button>
        {/each}
      </div>
    </div>
  {/if}

  <!-- Fanart grid -->
  {#if activeTypes.has('Fanart')}
    <div class="px-3 py-2 border-b border-base-300">
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">
        Fanart <span class="text-base-content/40 font-normal normal-case tracking-normal">{DEMO_FANARTS.length} available</span>
      </div>
      <div class="grid grid-cols-4 gap-1">
        {#each DEMO_FANARTS as dim, i}
          {@const id = 'fanart-' + i}
          <button
            class="relative bg-base-200 rounded border cursor-pointer flex items-center justify-center text-base-content/30 overflow-hidden {selectedImages.has(id) ? 'border-primary' : 'border-base-300'}"
            style="aspect-ratio:16/9"
            onclick={() => toggleImg(id)}
          >
            🖼️
            <span class="absolute bottom-0.5 left-1 text-[8px] text-white/70">{dim}</span>
          </button>
        {/each}
      </div>
    </div>
  {/if}

  <!-- ExtraFanart -->
  {#if activeTypes.has('ExtraFanart')}
    <div class="px-3 py-2 border-b border-base-300">
      <div class="flex items-center justify-between mb-1">
        <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50">ExtraFanart</div>
        <div class="flex items-center gap-1 text-[10px] text-base-content">
          Max: <input
            type="number"
            class="input input-xs w-10 text-center"
            bind:value={maxExtrafanart}
            min="0" max="20"
          />
        </div>
      </div>
      <div class="text-[10px] text-base-content/40 mb-1">Saved to <code class="text-[9px]">extrafanart/</code> subfolder</div>
      <div class="grid grid-cols-4 gap-1">
        {#each DEMO_EXTRAFAN.slice(0, maxExtrafanart + 1) as lbl, i}
          {@const id = 'ef-' + i}
          <button
            class="relative bg-base-200 rounded border cursor-pointer flex items-center justify-center text-base-content/30 overflow-hidden {selectedImages.has(id) ? 'border-primary' : 'border-base-300'}"
            style="aspect-ratio:16/9"
            onclick={() => toggleImg(id)}
          >
            🖼️
            <span class="absolute bottom-0.5 left-1 text-[8px] text-white/70">{lbl}</span>
          </button>
        {/each}
      </div>
    </div>
  {/if}

  <!-- Thumb -->
  {#if activeTypes.has('Thumb')}
    <div class="px-3 py-2 border-b border-base-300">
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Thumb / Episode Still</div>
      <div class="grid grid-cols-4 gap-1">
        {#each DEMO_THUMBS as lbl, i}
          {@const id = 'thumb-' + i}
          <button
            class="relative bg-base-200 rounded border cursor-pointer flex items-center justify-center text-base-content/30 overflow-hidden {selectedImages.has(id) ? 'border-primary' : 'border-base-300'}"
            style="aspect-ratio:16/9"
            onclick={() => toggleImg(id)}
          >🖼️</button>
        {/each}
      </div>
    </div>
  {/if}

  <!-- Download buttons -->
  <div class="px-3 py-2">
    <div class="flex gap-1.5">
      <button class="btn btn-primary btn-xs flex-1 text-[11px]">↓ Download Selected ({selectedImages.size})</button>
      <button class="btn btn-outline btn-xs text-[11px]">↓ All</button>
    </div>
  </div>
</div>
