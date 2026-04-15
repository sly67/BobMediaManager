<script lang="ts">
  import { formatTemplate } from '$lib/stores';

  const PRESETS = [
    { label: 'Anime',   value: '{n} ({y})/Season {s00}/{n} - S{s00}E{e00} - {t}' },
    { label: 'Series',  value: '{n} ({y})/Season {s00}/{n} - S{s00}E{e00} - {t}' },
    { label: 'Movie',   value: '{n} ({y})/{n} ({y})' },
    { label: 'Music',   value: '{artist}/{album}/{track00} - {n}' },
  ];

  const TOKENS = [
    '{n}', '{y}', '{s}', '{e}', '{s00}', '{e00}', '{t}',
    '{artist}', '{album}', '{track00}', '{ext}'
  ];

  let showPresets = $state(false);
  let showTokens = $state(false);

  function applyPreset(val: string) {
    formatTemplate.set(val);
    showPresets = false;
  }
</script>

<div class="relative flex items-center gap-2 px-3 py-[5px] bg-base-200 border-b border-base-300 shrink-0 text-xs">
  <span class="text-base-content/50 shrink-0">Format:</span>
  <input
    class="flex-1 bg-base-300 border border-base-300 rounded px-2 py-0.5 text-secondary font-mono text-[11px] focus:outline-primary"
    value={$formatTemplate}
    oninput={(e) => formatTemplate.set((e.target as HTMLInputElement).value)}
  />

  <!-- Presets dropdown -->
  <div class="relative">
    <button class="btn btn-ghost btn-xs" onclick={() => { showPresets = !showPresets; showTokens = false; }}>Presets</button>
    {#if showPresets}
      <div class="absolute right-0 top-full mt-1 bg-base-100 border border-base-300 rounded-lg shadow-xl z-50 py-1 min-w-[160px]">
        {#each PRESETS as p}
          <button class="block w-full text-left px-3 py-1.5 text-xs hover:bg-base-200" onclick={() => applyPreset(p.value)}>
            {p.label}
          </button>
        {/each}
      </div>
    {/if}
  </div>

  <!-- Token reference dropdown -->
  <div class="relative">
    <button class="btn btn-ghost btn-xs" onclick={() => { showTokens = !showTokens; showPresets = false; }}>Tokens</button>
    {#if showTokens}
      <div class="absolute right-0 top-full mt-1 bg-base-100 border border-base-300 rounded-lg shadow-xl z-50 py-2 px-3 min-w-[220px]">
        <div class="text-[9px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">Available tokens</div>
        <div class="flex flex-wrap gap-1">
          {#each TOKENS as tok}
            <button
              class="font-mono text-[10px] px-1.5 py-0.5 rounded bg-base-200 text-secondary border border-base-300 hover:border-primary cursor-pointer"
              onclick={() => { formatTemplate.update(v => v + tok); showTokens = false; }}
            >
              {tok}
            </button>
          {/each}
        </div>
        <div class="mt-2 text-[9px] text-base-content/40 space-y-0.5">
          <div><span class="text-secondary font-mono">{'{n}'}</span> title &nbsp; <span class="text-secondary font-mono">{'{y}'}</span> year</div>
          <div><span class="text-secondary font-mono">{'{s00}'}</span> season 0-padded &nbsp; <span class="text-secondary font-mono">{'{e00}'}</span> ep 0-padded</div>
          <div><span class="text-secondary font-mono">{'{t}'}</span> episode title &nbsp; <span class="text-secondary font-mono">{'{ext}'}</span> extension</div>
        </div>
      </div>
    {/if}
  </div>
</div>
