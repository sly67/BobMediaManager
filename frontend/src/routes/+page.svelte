<script lang="ts">
	import { onMount } from 'svelte';

	// --- stores (to be expanded in Phase 9) ---
	let activeTab = $state('rename');
	let theme = $state('dark');

	const TABS = [
		{ id: 'rename',    label: 'Rename & Sort' },
		{ id: 'metadata',  label: 'Metadata' },
		{ id: 'nfo',       label: 'NFO' },
		{ id: 'images',    label: 'Images' },
		{ id: 'subtitles', label: 'Subtitles' },
		{ id: 'verify',    label: 'Verify' },
		{ id: 'history',   label: 'History' },
		{ id: 'sources',   label: '⚙ Sources' },
	];

	function setTheme(t: string) {
		theme = t;
		document.documentElement.setAttribute('data-theme', t);
		localStorage.setItem('bmm-theme', t);
	}

	onMount(() => {
		const saved = localStorage.getItem('bmm-theme');
		if (saved) setTheme(saved);

		// Connect WebSocket
		const ws = new WebSocket(`ws://${location.host}/ws/progress`);
		ws.onopen = () => console.log('WS connected');
		ws.onclose = () => console.log('WS closed');
		return () => ws.close();
	});
</script>

<!-- Top bar -->
<header class="navbar bg-base-200 border-b border-base-300 px-4 gap-2 min-h-12">
	<span class="font-bold text-primary text-lg mr-4">BobMediaManager</span>

	<!-- Tabs -->
	<div class="tabs tabs-bordered flex-1">
		{#each TABS as tab}
			<button
				class="tab {activeTab === tab.id ? 'tab-active' : ''}"
				onclick={() => activeTab = tab.id}
			>
				{tab.label}
			</button>
		{/each}
	</div>

	<!-- Theme picker (placeholder — full picker in Phase 9) -->
	<div class="dropdown dropdown-end">
		<button tabindex="0" class="btn btn-ghost btn-sm">Theme ▾</button>
		<ul class="dropdown-content menu bg-base-100 rounded-box shadow z-50 p-2 w-40">
			{#each ['light','dark','cupcake','cyberpunk','dracula','forest','luxury','night','sunset'] as t}
				<li><button onclick={() => setTheme(t)}>{t}</button></li>
			{/each}
		</ul>
	</div>

	<button class="btn btn-primary btn-sm">Execute</button>
</header>

<!-- Main layout -->
<main class="flex flex-1 overflow-hidden" style="height: calc(100vh - 3.5rem);">
	<!-- Left panel: source folder tree -->
	<aside class="w-56 border-r border-base-300 bg-base-100 flex flex-col p-2 overflow-auto shrink-0">
		<div class="text-xs font-semibold uppercase text-base-content/50 mb-2">Source</div>
		<div class="text-sm text-base-content/40 italic">No folder selected</div>
	</aside>

	<!-- Center: workflow area -->
	<section class="flex flex-col flex-1 overflow-hidden">
		{#if activeTab === 'rename'}
			<!-- Toolbar row -->
			<div class="flex items-center gap-2 px-3 py-2 border-b border-base-300 bg-base-100 shrink-0">
				<select class="select select-bordered select-sm w-28">
					<option>All</option>
					<option>Movie</option>
					<option>Series</option>
					<option>Anime</option>
					<option>Music</option>
				</select>
				<span class="badge badge-outline badge-sm">TMDB</span>
				<span class="badge badge-outline badge-sm">TVDB</span>
				<span class="badge badge-outline badge-sm">AniDB</span>
				<span class="badge badge-outline badge-sm">IMDb</span>
				<div class="flex-1"></div>
				<input type="text" placeholder="Search / ID…" class="input input-bordered input-sm w-48" />
				<button class="btn btn-outline btn-sm">Pick ID</button>
			</div>

			<!-- Op strip -->
			<div class="flex items-center gap-3 px-3 py-2 border-b border-base-300 bg-base-200 shrink-0">
				<div class="join">
					<button class="btn btn-sm join-item btn-primary">Move</button>
					<button class="btn btn-sm join-item btn-outline">Copy</button>
					<button class="btn btn-sm join-item btn-outline">Rename in-place</button>
				</div>
				<input
					type="text"
					placeholder="/media/library/anime  (required)"
					class="input input-bordered input-sm flex-1 border-error"
				/>
				<label class="flex items-center gap-1 text-sm cursor-pointer">
					<input type="checkbox" class="checkbox checkbox-sm" />
					Rename folders too
				</label>
			</div>

			<!-- Format row -->
			<div class="flex items-center gap-2 px-3 py-2 border-b border-base-300 bg-base-100 shrink-0">
				<span class="text-sm font-medium shrink-0">Format:</span>
				<input
					type="text"
					value="{n} ({y})/Season {s00}/{n} - S{s00}E{e00} - {t}"
					class="input input-bordered input-sm flex-1 font-mono text-xs"
				/>
				<button class="btn btn-ghost btn-sm">Presets</button>
				<button class="btn btn-ghost btn-sm">Tokens</button>
			</div>

			<!-- Split view: source | arrow | renamed -->
			<div class="flex flex-1 overflow-hidden">
				<!-- Source column -->
				<div class="flex flex-col flex-1 overflow-hidden border-r border-base-300">
					<div class="px-2 py-1 bg-base-200 border-b border-base-300">
						<input type="text" placeholder="Filter source…" class="input input-bordered input-xs w-full" />
					</div>
					<div class="flex-1 overflow-auto text-sm p-2 text-base-content/40 italic">
						Scan a folder to begin
					</div>
				</div>

				<!-- Arrow column -->
				<div class="flex flex-col items-center justify-center w-8 bg-base-200 shrink-0 text-base-content/30 text-lg">
					→
				</div>

				<!-- Renamed column -->
				<div class="flex flex-col flex-1 overflow-hidden">
					<div class="px-2 py-1 bg-base-200 border-b border-base-300">
						<input type="text" placeholder="Filter renamed…" class="input input-bordered input-xs w-full" />
					</div>
					<div class="flex-1 overflow-auto text-sm p-2 text-base-content/40 italic">
						Previews will appear here
					</div>
				</div>
			</div>
		{:else}
			<div class="flex-1 flex items-center justify-center text-base-content/30 text-sm">
				{activeTab} — coming in Phase 9
			</div>
		{/if}
	</section>

	<!-- Right panel (collapsible) -->
	<aside class="w-72 border-l border-base-300 bg-base-100 flex flex-col shrink-0 overflow-hidden">
		<div class="tabs tabs-bordered px-2 pt-1 shrink-0">
			<button class="tab tab-active tab-sm">Metadata</button>
			<button class="tab tab-sm">Images</button>
			<button class="tab tab-sm">Companions</button>
			<button class="tab tab-sm">NFO</button>
		</div>
		<div class="flex-1 p-3 text-sm text-base-content/40 italic overflow-auto">
			Select a file to see metadata
		</div>
	</aside>
</main>

<!-- Status bar -->
<footer class="flex items-center gap-3 px-4 h-6 bg-base-300 text-xs text-base-content/60 border-t border-base-300 shrink-0">
	<span class="w-2 h-2 rounded-full bg-success inline-block"></span>
	<span>TMDB</span>
	<span class="w-2 h-2 rounded-full bg-warning inline-block"></span>
	<span>TVDB</span>
	<span class="w-2 h-2 rounded-full bg-success inline-block"></span>
	<span>AniDB</span>
	<div class="flex-1"></div>
	<span>:48040</span>
</footer>

<style>
	:global(html, body) {
		height: 100%;
		overflow: hidden;
		display: flex;
		flex-direction: column;
	}
	main {
		flex: 1;
	}
</style>
