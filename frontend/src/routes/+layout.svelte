<script lang="ts">
	import '../app.css';
	import { onMount } from 'svelte';

	// Persist theme across reloads
	onMount(() => {
		const saved = localStorage.getItem('bmm-theme');
		if (saved) {
			document.documentElement.setAttribute('data-theme', saved);
		}
	});
</script>

<slot />
