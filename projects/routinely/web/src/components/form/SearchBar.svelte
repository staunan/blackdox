<script>
	import FormLabel from "components/form/FormLabel.svelte";
	import FormErrorMessage from "components/form/FormErrorMessage.svelte";
	export let placeholder = "Input Placeholder...";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();
	export let value = "";

	let typingTimer = null;
	const typingDelay = 1000;

	function onTextChange(event) {
		clearTimeout(typingTimer); // Clear the previous timer
		typingTimer = setTimeout(() => {
			dispatch("change", event.target.value);
		}, typingDelay);
	}
</script>

<div class="search_box">
	<input type="text" {placeholder} {value} on:keyup={onTextChange} />
</div>

<style>
	.search_box {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.search_box input {
		height: 36px;
		width: 300px;
		outline: none;
		border: 1px solid #b5b5b5;
		box-sizing: border-box;
		padding-left: 10px;
		box-shadow: inset 0 0.0625em 0.125em rgba(10, 10, 10, 0.05);
		font-size: 16px;
		font-weight: 500;
		font-family: monospace;
		border-radius: 30px;
	}
	.search_box input:hover,
	.search_box input:focus {
		border-color: #3e8ed0;
	}
</style>
