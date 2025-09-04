<script>
	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [children]
	 */

	/** @type {Props} */
	let { children } = $props();
</script>

<div class="center_layout">
	{@render children?.()}
</div>

<style>
	.center_layout {
		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>
