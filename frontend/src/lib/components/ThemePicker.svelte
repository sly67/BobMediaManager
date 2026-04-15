<script lang="ts">
  import { LIGHT_THEMES, DARK_THEMES, getThemeDef } from '$lib/themes';
  import { theme } from '$lib/stores';

  let open = $state(false);

  function setTheme(name: string) {
    theme.set(name);
    document.documentElement.setAttribute('data-theme', name);
    localStorage.setItem('bmm-theme', name);
    open = false;
  }

  function toggle(e: MouseEvent) {
    e.stopPropagation();
    open = !open;
  }

  function closeOnOutside(e: MouseEvent) {
    open = false;
  }

  $effect(() => {
    if (open) {
      document.addEventListener('click', closeOnOutside, { once: true });
    }
  });

  const currentColors = $derived(
    getThemeDef($theme)?.colors ?? ['#1d232a', '#661ae4', '#d926a9', '#1fb2a6']
  );
</script>

<div class="relative">
  <!-- Trigger button -->
  <button
    class="flex items-center gap-1.5 px-3 py-1 rounded-md border border-base-300 bg-base-200 text-base-content text-xs cursor-pointer hover:opacity-85"
    onclick={toggle}
  >
    <!-- 4-color swatch of current theme -->
    <span class="flex rounded-full overflow-hidden w-3 h-3 shrink-0 border border-base-300">
      {#each currentColors as c}
        <span class="flex-1" style="background:{c}"></span>
      {/each}
    </span>
    <span>{$theme}</span>
    <span class="opacity-50">▾</span>
  </button>

  <!-- Dropdown -->
  {#if open}
    <!-- svelte-ignore a11y_click_events_have_key_events -->
    <!-- svelte-ignore a11y_no_static_element_interactions -->
    <div
      class="absolute right-0 top-[calc(100%+6px)] bg-base-100 border border-base-300 rounded-xl p-3 w-[370px] shadow-2xl z-50"
      onclick={(e) => e.stopPropagation()}
    >
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">☀️ Light</div>
      <div class="grid grid-cols-5 gap-1 mb-3">
        {#each LIGHT_THEMES as t}
          <button
            class="flex flex-col items-center gap-0.5 cursor-pointer p-1 rounded-md border {$theme === t.name ? 'border-primary bg-base-200' : 'border-transparent hover:bg-base-200'}"
            title={t.name}
            onclick={() => setTheme(t.name)}
          >
            <span class="flex rounded overflow-hidden border border-black/10" style="width:52px;height:20px;">
              {#each t.colors as c}
                <span class="flex-1" style="background:{c}"></span>
              {/each}
            </span>
            <span class="text-[8px] text-base-content/50 w-full text-center truncate">{t.name}</span>
          </button>
        {/each}
      </div>

      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">🌙 Dark</div>
      <div class="grid grid-cols-5 gap-1">
        {#each DARK_THEMES as t}
          <button
            class="flex flex-col items-center gap-0.5 cursor-pointer p-1 rounded-md border {$theme === t.name ? 'border-primary bg-base-200' : 'border-transparent hover:bg-base-200'}"
            title={t.name}
            onclick={() => setTheme(t.name)}
          >
            <span class="flex rounded overflow-hidden border border-black/20" style="width:52px;height:20px;">
              {#each t.colors as c}
                <span class="flex-1" style="background:{c}"></span>
              {/each}
            </span>
            <span class="text-[8px] text-base-content/50 w-full text-center truncate">{t.name}</span>
          </button>
        {/each}
      </div>
    </div>
  {/if}
</div>
