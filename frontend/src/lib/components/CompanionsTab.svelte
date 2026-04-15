<script lang="ts">
  interface CompGroup {
    title: string;
    items: { icon: string; label: string; default: boolean }[];
  }

  const GROUPS: CompGroup[] = [
    {
      title: 'NFO Files',
      items: [
        { icon: '📄', label: 'tvshow.nfo',  default: true },
        { icon: '📄', label: 'episode.nfo', default: true },
        { icon: '📄', label: 'movie.nfo',   default: false },
        { icon: '📄', label: 'artist.nfo',  default: false },
        { icon: '📄', label: 'album.nfo',   default: false },
      ],
    },
    {
      title: 'Series-Level Images',
      items: [
        { icon: '🖼️', label: 'poster.jpg',      default: true },
        { icon: '🖼️', label: 'fanart.jpg',       default: true },
        { icon: '🖼️', label: 'extrafanart/',     default: true },
        { icon: '🖼️', label: 'banner.jpg',       default: false },
        { icon: '🖼️', label: 'logo.png',         default: false },
        { icon: '🖼️', label: 'clearart.png',     default: false },
        { icon: '🖼️', label: 'landscape.jpg',    default: false },
        { icon: '🖼️', label: 'disc.png',         default: false },
      ],
    },
    {
      title: 'Season-Level Images',
      items: [
        { icon: '🖼️', label: 'season{N}-poster.jpg', default: true },
        { icon: '🖼️', label: 'season{N}-fanart.jpg',  default: true },
        { icon: '🖼️', label: 'season{N}-banner.jpg',  default: false },
        { icon: '🖼️', label: 'season-all-poster.jpg', default: false },
      ],
    },
    {
      title: 'Episode-Level Images',
      items: [
        { icon: '🖼️', label: '{ep}-thumb.jpg',     default: true },
        { icon: '🖼️', label: '{ep}-landscape.jpg', default: false },
      ],
    },
    {
      title: 'Music Companions',
      items: [
        { icon: '🖼️', label: 'folder.jpg',  default: false },
        { icon: '🖼️', label: 'artist.jpg',  default: false },
        { icon: '📄', label: 'artist.nfo',  default: false },
        { icon: '📄', label: 'album.nfo',   default: false },
      ],
    },
    {
      title: 'Subtitle Companions',
      items: [
        { icon: '💬', label: '{ep}.en.srt', default: true },
        { icon: '💬', label: '{ep}.fr.srt', default: false },
        { icon: '💬', label: '{ep}.ja.srt', default: false },
      ],
    },
  ];

  // Build reactive state from defaults
  let checked = $state<Record<string, boolean>>(
    Object.fromEntries(
      GROUPS.flatMap(g => g.items.map(item => [`${g.title}::${item.label}`, item.default]))
    )
  );

  function toggle(group: string, label: string) {
    const key = `${group}::${label}`;
    checked = { ...checked, [key]: !checked[key] };
  }
</script>

<div class="overflow-y-auto flex-1">
  {#each GROUPS as group, gi}
    <div class="px-3 py-2 {gi < GROUPS.length - 1 ? 'border-b border-base-300' : ''}">
      <div class="text-[10px] font-bold uppercase tracking-widest text-base-content/50 mb-1.5">{group.title}</div>
      <div class="grid grid-cols-2 gap-1">
        {#each group.items as item}
          {@const key = `${group.title}::${item.label}`}
          {@const on = checked[key]}
          <!-- svelte-ignore a11y_click_events_have_key_events -->
          <!-- svelte-ignore a11y_no_static_element_interactions -->
          <label
            class="flex items-center gap-1.5 px-2 py-1.5 rounded-md border text-[11px] cursor-pointer {on ? 'border-primary bg-primary/8 text-base-content' : 'border-base-300 bg-base-200 text-base-content/60'}"
            onclick={() => toggle(group.title, item.label)}
          >
            <input type="checkbox" class="checkbox checkbox-xs" checked={on} onchange={() => toggle(group.title, item.label)} />
            <span class="text-sm leading-none">{item.icon}</span>
            <span class="truncate">{item.label}</span>
          </label>
        {/each}
      </div>
    </div>
  {/each}
</div>
