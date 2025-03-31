<script>
	import Header from "./Header.svelte";
	import "../app.css";

	let { children } = $props();
</script>

<div class="app">
	<Header />

	<main>
		{@render children()}
	</main>
</div>

<style>
	.app {
		min-height: 100vh;
	}

	main {
		width: 70%;
		margin: 0 auto;
	}
</style>
