<script>
	/**
	 * @typedef {Object} Props
	 * @property {number} [padding]
	 * @property {string} [background]
	 * @property {import('svelte').Snippet} [children]
	 */

	/** @type {Props} */
	let { padding = 0, background = "", children } = $props();
</script>

<div
	class="shadow card"
	style={"padding:" +
		(padding ? padding * 5 : 0) +
		"px;" +
		"background:" +
		(background ? background : "transparent") +
		";"}
>
	{@render children?.()}
</div>

<style>
	.card {
		width: 100%;
		height: auto;
		border-radius: 4px;
		font-family: monospace;
	}
	.shadow {
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
	}
</style>
