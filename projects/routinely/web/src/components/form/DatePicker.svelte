<script>
	import FormLabel from "components/form/FormLabel.svelte";
	import { DateInput } from "date-picker-svelte";
	import {
		ConvertMySQLDateTimeToJSDateTime,
		ConvertJSDateToMySQLDate,
	} from "lib/js/datetime.js";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	export let label = "";
	export let value = 0;

	let initializing = false;
	let month_date = null;

	$: {
		if (value) {
			initializing = true;
			if (value.length == 10) {
				value = value + " 00:00:00";
			}
			month_date = ConvertMySQLDateTimeToJSDateTime(value);
		}
	}
	$: {
		if (month_date) {
			if (!initializing) {
				dispatch("change", ConvertJSDateToMySQLDate(month_date));
			} else {
				initializing = false;
			}
		}
	}
</script>

<div class="date_picker">
	{#if label}
		<FormLabel {label}></FormLabel>
	{/if}
	<DateInput
		placeholder="Select a date"
		format="yyyy-MM-dd"
		bind:value={month_date}
	/>
</div>

<style>
	:global(.date-time-field input) {
		font-family: monospace;
		height: 36px;
		font-size: 18px;
		border-radius: 0 !important;
		width: 100% !important;
	}
	:global(.picker) {
		font-family: monospace;
		font-size: 16px;
	}
</style>
