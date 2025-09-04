<script>
	/**
	 * @typedef {Object} Props
	 * @property {string} [label]
	 */

	/** @type {Props} */
	let { label = "Input Label" } = $props();
</script>

<div class="form_label">{label}</div>

<style>
	.form_label {
		font-family: monospace;
		font-size: 14px;
		font-weight: bold;
		padding-bottom: 5px;
	}
</style>
