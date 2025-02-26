<script>
	import Dropdown from "components/form/Dropdown.svelte";
	import FormLabel from "components/form/FormLabel.svelte";
	import { createEventDispatcher } from "svelte";

	export let label = "";
	export let value = 0;

	const dispatch = createEventDispatcher();
	let Days = [
		{ label: "Anyday", value: 0 },
		{ label: "End of the Month", value: 32 },
		{ label: "The day before end of the Month", value: 33 },
		{ label: "01", value: 1 },
		{ label: "02", value: 2 },
		{ label: "03", value: 3 },
		{ label: "04", value: 4 },
		{ label: "05", value: 5 },
		{ label: "06", value: 6 },
		{ label: "07", value: 7 },
		{ label: "08", value: 8 },
		{ label: "09", value: 9 },
		{ label: "10", value: 10 },
		{ label: "11", value: 11 },
		{ label: "12", value: 12 },
		{ label: "13", value: 13 },
		{ label: "14", value: 14 },
		{ label: "15", value: 15 },
		{ label: "16", value: 16 },
		{ label: "17", value: 17 },
		{ label: "18", value: 18 },
		{ label: "19", value: 19 },
		{ label: "20", value: 20 },
		{ label: "21", value: 21 },
		{ label: "22", value: 22 },
		{ label: "23", value: 23 },
		{ label: "24", value: 24 },
		{ label: "25", value: 25 },
		{ label: "26", value: 26 },
		{ label: "27", value: 27 },
		{ label: "28", value: 28 },
		{ label: "29", value: 29 },
		{ label: "30", value: 30 },
		{ label: "31", value: 31 },
	];
	let day = null;

	$: {
		if (value && value <= 33) {
			day = Days.filter((d) => d.value == value)[0];
		} else if (value === 0) {
			day = Days[0];
		}
	}

	function dayChangeHandler(event) {
		day = event.detail;
		dispatch("change", day.value);
	}
</script>

<div class="daypicker_container">
	{#if label}
		<FormLabel {label}></FormLabel>
	{/if}
	<div class="daypicker_input">
		<Dropdown
			placeholder="00"
			items={Days}
			currentitem={day}
			on:change={dayChangeHandler}
		/>
	</div>
</div>
