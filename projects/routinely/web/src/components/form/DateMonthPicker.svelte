<script>
	import Dropdown from "components/form/Dropdown.svelte";
	import FormLabel from "components/form/FormLabel.svelte";
	import { onMount } from "svelte";
	import { createEventDispatcher } from "svelte";

	export let label = "";
	export let value = null;

	const dispatch = createEventDispatcher();
	let Months = [
		{ label: "Anyday of the Year", value: 0 },
		{ label: "January", value: 1 },
		{ label: "February", value: 2 },
		{ label: "March", value: 3 },
		{ label: "April", value: 4 },
		{ label: "May", value: 5 },
		{ label: "June", value: 6 },
		{ label: "July", value: 7 },
		{ label: "August", value: 8 },
		{ label: "September", value: 9 },
		{ label: "October", value: 10 },
		{ label: "November", value: 11 },
		{ label: "December", value: 12 },
	];
	let Days = [];
	let dayDropdownDisabled = false;
	let selectedMonth = null;
	let selectedDay = null;

	$: {
		if (value) {
			let arr = value.split("-");
			let initialMonth = parseInt(arr[0]);
			let initialDay = parseInt(arr[1]);
			if (initialMonth === 0) {
				selectedMonth = Months[0];
			} else {
				selectedMonth = Months.filter(
					(m) => m.value === initialMonth
				)[0];
			}
			if (initialDay === 0) {
				selectedDay = null;
			} else {
				selectedDay = Days.filter((d) => d.value === initialDay)[0];
			}
		} else {
			changeMonth(Months[0]);
		}
	}

	onMount(() => {
		Days = calculateDays();
		if (selectedMonth == null || selectedMonth.value === 0) {
			dayDropdownDisabled = true;
		} else {
			dayDropdownDisabled = false;
		}
	});
	function monthChangeHandler(event) {
		selectedMonth = event.detail;
		changeMonth(selectedMonth);
	}
	function changeMonth(month) {
		if (
			[
				"January",
				"March",
				"May",
				"July",
				"August",
				"October",
				"December",
			].includes(month.label)
		) {
			Days = calculateDays(31);
		} else if (
			["April", "June", "September", "November"].includes(month.label)
		) {
			Days = calculateDays(30);
		} else if (["February"].includes(month.label)) {
			Days = calculateDays(29);
		}
		if (month.value === 0) {
			dayDropdownDisabled = true;
			selectedDay = null;
		} else {
			selectedDay = Days[0];
			dayDropdownDisabled = false;
		}
		emitValue();
	}
	function dayChangeHandler(event) {
		selectedDay = event.detail;
		emitValue();
	}
	function calculateDays(length) {
		if (!length) {
			length = 30;
		}
		let arr = [];
		// arr.push({label: "00", value: 0});
		for (let i = 1; i <= length; i++) {
			let label = i;
			if (i < 10) {
				label = "0" + i;
			}
			arr.push({ label: label, value: i });
		}
		return arr;
	}
	function emitValue() {
		let m = "";
		let d = "";
		if (selectedMonth) {
			if (selectedMonth.value < 10) {
				m = "0" + selectedMonth.value;
			} else {
				m = selectedMonth.value;
			}
		} else {
			m = "00";
		}
		if (selectedDay) {
			if (selectedDay.value < 10) {
				d = "0" + selectedDay.value;
			} else {
				d = selectedDay.value;
			}
		} else {
			d = "00";
		}
		dispatch("change", m + "-" + d);
	}
</script>

<div class="datemonthpicker_container">
	{#if label}
		<FormLabel {label}></FormLabel>
	{/if}
	<div class="datemonthpicker_input">
		<div class="month_input">
			<Dropdown
				placeholder="Month"
				items={Months}
				currentitem={selectedMonth}
				on:change={monthChangeHandler}
			/>
		</div>
		<div class="day_input" class:disabled={dayDropdownDisabled}>
			<Dropdown
				placeholder="Day"
				items={Days}
				currentitem={selectedDay}
				on:change={dayChangeHandler}
				disabled={dayDropdownDisabled}
			/>
		</div>
	</div>
</div>

<style>
	.datemonthpicker_input {
		display: flex;
		width: 100%;
	}
	.month_input {
		flex-basis: 70%;
	}
	.day_input {
		flex-basis: 30%;
	}
	.day_input.disabled {
		cursor: not-allowed;
	}
</style>
