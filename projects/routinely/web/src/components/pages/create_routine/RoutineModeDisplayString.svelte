<script>
	import RichText from "components/typography/RichText.svelte";
	/**
	 * @typedef {Object} Props
	 * @property {string} [message]
	 */

	/** @type {Props} */
	let { message = "" } = $props();
</script>

<div class="routine_mode_display_string">
	<RichText content={message}></RichText>
</div>

<style>
	.routine_mode_display_string {
		padding-top: 10px;
		padding-bottom: 10px;
		font-size: 14px;
		font-family: monospace;
		color: green;
	}
</style>
