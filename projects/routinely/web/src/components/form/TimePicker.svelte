<script>
	import { onMount } from "svelte";
	import Dropdown from "components/form/Dropdown.svelte";
	import FormLabel from "components/form/FormLabel.svelte";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();
	export let label = "";
	export let value = "00:00";
	export let format = "24Hours";

	let Hours = [];
	let Minutes = [
		{ label: "00", value: 0 },
		{ label: "05", value: 5 },
		{ label: "10", value: 10 },
		{ label: "15", value: 15 },
		{ label: "20", value: 20 },
		{ label: "25", value: 25 },
		{ label: "30", value: 30 },
		{ label: "35", value: 35 },
		{ label: "40", value: 40 },
		{ label: "45", value: 45 },
		{ label: "50", value: 50 },
		{ label: "55", value: 55 },
	];

	let selectedHour = null;
	let selectedMinute = null;

	onMount(() => {
		if (format === "24Hours") {
			Hours = [
				{ label: "00", value: 0 },
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
			];
		} else {
			Hours = [
				{ label: "00", value: 0 },
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
			];
		}
	});

	$: {
		if (value) {
			let arr = value.split(":");
			try {
				selectedHour = Hours.filter((h) => h.label == arr[0])[0];
				selectedMinute = Minutes.filter((m) => m.label == arr[1])[0];
			} catch (err) {
				console.log("Error while parsing time input");
				console.log(err);
			}
		}
	}

	function hourChangeHandler(event) {
		selectedHour = event.detail;
		calculateTime(event.detail.label, selectedMinute.label);
	}
	function minuteChangeHandler(event) {
		selectedMinute = event.detail;
		calculateTime(selectedHour.label, event.detail.label);
	}
	function calculateTime(hour, minute) {
		let time = hour + ":" + minute;
		dispatch("change", time);
	}
</script>

<div class="timepicker_container">
	{#if label}
		<FormLabel {label}></FormLabel>
	{/if}
	<div class="timepicker_input">
		<div class="hour_input">
			<Dropdown
				placeholder="Hours"
				items={Hours}
				currentitem={selectedHour}
				on:change={hourChangeHandler}
			/>
		</div>
		<div class="minute_input">
			<Dropdown
				placeholder="Minute"
				items={Minutes}
				currentitem={selectedMinute}
				on:change={minuteChangeHandler}
			/>
		</div>
	</div>
</div>

<style>
	.timepicker_input {
		display: flex;
		width: 100%;
	}
	.hour_input,
	.minute_input {
		flex: 1;
	}
</style>
