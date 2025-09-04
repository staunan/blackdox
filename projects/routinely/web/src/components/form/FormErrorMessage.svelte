<script>
	/**
	 * @typedef {Object} Props
	 * @property {string} [message]
	 */

	/** @type {Props} */
	let { message = "Error Message..." } = $props();
</script>

<div class="error_message">{message}</div>

<style>
	.error_message {
		color: red;
		font-size: 14px;
		font-family: monospace;
		margin-top: 5px;
	}
</style>
