<script>
	/**
	 * @typedef {Object} Props
	 * @property {string} [title]
	 */

	/** @type {Props} */
	let { title = "" } = $props();
</script>

<div class="header_title">{title}</div>

<style>
	.header_title {
		display: flex;
		font-size: 24px;
		font-weight: bold;
	}
</style>
