<script>
	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [children]
	 */

	/** @type {Props} */
	let { children } = $props();
</script>

<div class="right_layout">
	{@render children?.()}
</div>

<style>
	.right_layout {
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}
</style>
