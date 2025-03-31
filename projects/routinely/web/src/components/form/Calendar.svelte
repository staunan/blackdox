<script>
	import { onMount } from "svelte";
	import ArrowDown from "components/svg/ArrowDown.svelte";

	let active = false;
	let date = new Date();
	let year = date.getFullYear();
	let month = date.getMonth();
	let calendar_days = [];

	onMount(() => {
		const funcRef = (event) => {
			if (event.target.closest(".calendar_dropdown_trigger")) {
				active = true;
			} else if (!event.target.closest(".calendar_dropdown")) {
				active = false;
			}
		};
		window.addEventListener("click", funcRef);

		generateDays();

		// Called when component is destroyed --
		return () => {
			window.removeEventListener("click", funcRef);
		};
	});

	const months = [
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	];

	function generateDays() {
		// Get the first day of week of the previous month
		let dayone = new Date(year, month, 1).getDay();

		// Get the last date of the current month
		let lastdate = new Date(year, month + 1, 0).getDate();

		// Get the day of the last date of the month
		let dayend = new Date(year, month, lastdate).getDay();

		// Get the last date of the previous month
		let monthlastdate = new Date(year, month, 0).getDate();

		// Loop to add the last dates of the previous month
		let month_days = [];
		for (let i = dayone; i > 0; i--) {
			month_days.push({
				month: "previous",
				value: monthlastdate - i + 1,
				active: false,
			});
		}

		// Loop to add the dates of the current month
		for (let i = 1; i <= lastdate; i++) {
			month_days.push({
				month: "current",
				value: i,
				active: false,
			});
		}

		// Loop to add the first dates of the next month
		for (let i = dayend; i < 6; i++) {
			month_days.push({
				month: "next",
				value: i - dayend + 1,
				active: false,
			});
		}

		calendar_days = month_days;
	}
	function handleDropdownItemClick() {
		active = !active;
	}
	function dayClickedHandler() {}
	function goPrevHandler() {
		month = month - 1;

		// Check if the month is out of range
		if (month < 0) {
			// Set the date to the first day of the
			// month with the new year
			date = new Date(year, month, new Date().getDate());

			// Set the year to the new year
			year = date.getFullYear();

			// Set the month to the new month
			month = date.getMonth();
		} else {
			// Set the date to the current date
			date = new Date();
		}
		generateDays();
	}
	function goNextHandler() {
		month = month + 1;

		// Check if the month is out of range
		if (month > 11) {
			// Set the date to the first day of the
			// month with the new year
			date = new Date(year, month, new Date().getDate());

			// Set the year to the new year
			year = date.getFullYear();

			// Set the month to the new month
			month = date.getMonth();
		} else {
			// Set the date to the current date
			date = new Date();
		}
		generateDays();
	}

	generateDays();
</script>

<div class="calendar_wrapper">
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div
		class="calendar_dropdown_trigger"
		on:click={() => handleDropdownItemClick()}
	>
		Today
	</div>
	<div class="calendar_dropdown" class:show={active}>
		<div class="calendar_dropdown_content">
			<div class="month_navigation">
				<div class="month_info">{months[month]}, {year}</div>
				<div class="navigation_container">
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<div class="arrow_down" on:click={() => goPrevHandler()}>
						<ArrowDown></ArrowDown>
					</div>
					<!-- svelte-ignore a11y-click-events-have-key-events -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<div class="arrow_up" on:click={() => goNextHandler()}>
						<ArrowDown></ArrowDown>
					</div>
				</div>
			</div>
			<div class="calendar_week_days_grid">
				<div class="weekday_cell">Sun</div>
				<div class="weekday_cell">Mon</div>
				<div class="weekday_cell">Tue</div>
				<div class="weekday_cell">Wed</div>
				<div class="weekday_cell">Thu</div>
				<div class="weekday_cell">Fri</div>
				<div class="weekday_cell">Sat</div>
			</div>
			<div class="calendar_days_grid">
				{#each calendar_days as day}
					<!-- svelte-ignore a11y_click_events_have_key_events -->
					<!-- svelte-ignore a11y-no-static-element-interactions -->
					<div
						class="day_cell"
						class:selected={day.active === true}
						class:prev_month={day.month === "previous"}
						class:next_month={day.month === "next"}
						on:click={() => dayClickedHandler(day)}
						title={day.value}
					>
						{day.value}
					</div>
				{/each}
			</div>
		</div>
		<div class="dropdown_triangle"></div>
	</div>
</div>

<style>
	.dropdown_triangle {
		width: 20px;
		height: 20px;
		background-color: #ddd;
		position: absolute;
		left: 50%;
		top: -2px;
		z-index: -1;
		transform: rotate(45deg) translateX(-50%);
	}
	.calendar_wrapper {
		display: flex;
		justify-content: center;
		align-items: center;
		position: relative;
	}
	.calendar_dropdown_trigger {
		font-size: 24px;
		font-weight: bold;
		cursor: pointer;
	}
	.calendar_dropdown {
		display: none;
		position: absolute;
		left: 50%;
		top: 150%;
		transform: translateX(-50%);
		background-color: #ddd;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}
	.show {
		display: block !important;
	}
	.calendar_days_grid,
	.calendar_week_days_grid {
		display: grid;
		grid-template-columns: auto auto auto auto auto auto auto;
	}
	.calendar_week_days_grid {
		padding-top: 30px;
	}
	.day_cell {
		height: 50px;
		width: 80px;
		font-size: 16px;
		font-weight: bold;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: 300ms all;
	}
	.weekday_cell {
		width: 80px;
		font-size: 16px;
		font-weight: bold;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
		transition: 300ms all;
		padding-bottom: 5px;
	}
	.day_cell:hover {
		color: #fff;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		background: #0f2027; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#2c5364,
			#203a43,
			#0f2027
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#2c5364,
			#203a43,
			#0f2027
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
	}
	.prev_month,
	.next_month {
		color: #3f51b5;
	}
	.month_info {
		font-size: 18px;
		font-weight: bold;
		padding: 20px;
		padding-left: 25px;
		flex: 1;
	}
	.month_navigation {
		display: flex;
		align-items: center;
		border-bottom: 1px solid #ccc;
		height: 50px;
	}
	.arrow_down {
		padding-right: 20px;
		padding-left: 20px;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.arrow_up {
		transform: rotate(180deg);
		padding-right: 20px;
		padding-left: 20px;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.arrow_down:hover {
		background-color: #ccc;
	}
	.arrow_up:hover {
		background-color: #ccc;
	}
	.navigation_container {
		display: flex;
		padding-left: 20px;
		height: inherit;
	}
</style>
