<script>
	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [children]
	 */

	/** @type {Props} */
	let { children } = $props();
</script>

<div class="form_input">
	{@render children?.()}
</div>

<style>
	.form_input {
		padding-bottom: 30px;
	}
</style>
