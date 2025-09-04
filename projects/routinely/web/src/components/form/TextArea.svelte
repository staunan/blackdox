<script>
	import FormLabel from "components/form/FormLabel.svelte";
	import FormErrorMessage from "components/form/FormErrorMessage.svelte";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();
	/**
	 * @typedef {Object} Props
	 * @property {string} [label]
	 * @property {string} [placeholder]
	 * @property {boolean} [hasError]
	 * @property {string} [errorMessage]
	 * @property {string} [value]
	 */

	/** @type {Props} */
	let {
		label = "Input Label",
		placeholder = "Input Placeholder...",
		hasError = false,
		errorMessage = "",
		value = ""
	} = $props();
	function onTextChange(event) {
		dispatch("change", event.target.value);
	}
</script>

<div class="form_field">
	<FormLabel {label}></FormLabel>
	<textarea rows="5" {placeholder} {value} onkeyup={onTextChange}></textarea>
	{#if hasError}
		<FormErrorMessage message={errorMessage}></FormErrorMessage>
	{/if}
</div>

<style>
	.form_field textarea {
		outline: none;
		border: 1px solid #b5b5b5;
		padding-left: 10px;
		box-shadow: inset 0 0.0625em 0.125em rgba(10, 10, 10, 0.05);
		border-radius: 0.375em;
		width: 100%;
		font-size: 18px;
		font-weight: 300px;
		box-sizing: border-box;
		font-family: monospace;
		border-radius: 0;
	}
	.form_field textarea:hover,
	.form_field textarea:focus {
		border-color: #3e8ed0;
	}
</style>
