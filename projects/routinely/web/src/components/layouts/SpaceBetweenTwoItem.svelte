<script>
	/**
	 * @typedef {Object} Props
	 * @property {import('svelte').Snippet} [left]
	 * @property {import('svelte').Snippet} [right]
	 */

	/** @type {Props} */
	let { left, right } = $props();
</script>

<div class="row">
	<div class="column_left">
		{@render left?.()}
	</div>
	<div class="column_right">
		{@render right?.()}
	</div>
</div>

<style>
	.row {
		display: flex;
		flex-wrap: wrap;
		box-sizing: border-box;
		justify-content: space-between;
	}
	.column_left,
	.column_right {
		flex: 1;
	}
	.column_left {
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}
	.column_right {
		display: flex;
		justify-content: flex-end;
		align-items: center;
	}
</style>
