<script>
	import { run } from 'svelte/legacy';

	/**
	 * @typedef {Object} Props
	 * @property {string} [content]
	 */

	/** @type {Props} */
	let { content = "" } = $props();
	let node = $state();
	run(() => {
		if (content) {
			if (node) {
				node.innerHTML = content;
			} else {
				setTimeout(function () {
					if (node) {
						node.innerHTML = content;
					}
				}, 200);
			}
		}
	});
</script>

<div class="rich_text" bind:this={node}></div>
