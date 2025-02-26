<script>
	import FormLabel from "components/form/FormLabel.svelte";
	import FormErrorMessage from "components/form/FormErrorMessage.svelte";
	export let label = "Input Label";
	export let placeholder = "Input Placeholder...";
	export let hasError = false;
	export let errorMessage = "";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();
	export let value = "";
	function onTextChange(event) {
		dispatch("change", event.target.value);
	}
</script>

<div class="form_field">
	<FormLabel {label}></FormLabel>
	<input type="text" {placeholder} {value} on:keyup={onTextChange} />
	{#if hasError}
		<FormErrorMessage message={errorMessage}></FormErrorMessage>
	{/if}
</div>

<style>
	.form_field input {
		height: 36px;
		width: 100%;
		outline: none;
		border: 1px solid #b5b5b5;
		box-sizing: border-box;
		padding-left: 10px;
		box-shadow: inset 0 0.0625em 0.125em rgba(10, 10, 10, 0.05);
		border-radius: 0.375em;
		font-size: 16px;
		font-weight: 300px;
		font-family: monospace;
		font-size: 18px;
		border-radius: 0;
	}
	.form_field input:hover,
	.form_field input:focus {
		border-color: #3e8ed0;
	}
</style>
