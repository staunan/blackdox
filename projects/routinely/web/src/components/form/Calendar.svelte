<script>
	import { onMount } from "svelte";
	import ArrowDown from "components/svg/ArrowDown.svelte";
	import {
		ConvertMySQLDateTimeToJSDateTime,
		ConvertJSDateToMySQLDate,
		TodayDate,
		ParseDateToHumanReadableFormat,
	} from "lib/js/datetime.js";
	import { createEventDispatcher } from "svelte";

	export let value = null;
	export let user = null;

	const dispatch = createEventDispatcher();
	let active = false;
	let date = new Date();
	let calendar_days = [];
	let selected_date = null;
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

	$: {
		if (value) {
			if (value.length == 10) {
				value = value + " 00:00:00";
			}
			selected_date = ConvertMySQLDateTimeToJSDateTime(value);
			generateDays(selected_date);
		} else {
			generateDays();
		}
	}

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

	function generateDays(dateArg) {
		if (dateArg) {
			date = dateArg;
		} else {
			date = new Date();
		}
		let year = date.getFullYear();
		let month = date.getMonth();
		// Get the first day of week of the previous month
		let dayone = new Date(year, month, 1).getDay();

		// Get the last date of the current month
		let lastdate = new Date(year, month + 1, 0).getDate();

		// Get the day of the last date of the month
		let dayend = new Date(year, month, lastdate).getDay();

		// Get the last date of the previous month
		let monthlastdate = new Date(year, month, 0).getDate();

		let user_created_at = null;
		if (user) {
			user_created_at = ConvertMySQLDateTimeToJSDateTime(user.CreatedAt);
			user_created_at.setDate(user_created_at.getDate() - 1);
		}

		// Loop to add the last dates of the previous month
		let month_days = [];
		for (let i = dayone; i > 0; i--) {
			let cell_date = new Date(year, month - 1, monthlastdate - i + 1);
			let disabled = false;
			if (cell_date > new Date() || cell_date < user_created_at) {
				disabled = true;
			}
			month_days.push({
				month: "previous",
				value: monthlastdate - i + 1,
				active: false,
				date: cell_date,
				disabled: disabled,
			});
		}

		// Loop to add the dates of the current month
		for (let i = 1; i <= lastdate; i++) {
			let cell_date = new Date(year, month, i);
			let disabled = false;
			if (cell_date > new Date() || cell_date < user_created_at) {
				disabled = true;
			}
			if (
				i == new Date().getDate() &&
				month == new Date().getMonth() &&
				year == new Date().getFullYear()
			) {
				month_days.push({
					month: "current",
					value: i,
					active: false,
					today: true,
					date: cell_date,
					disabled: disabled,
				});
			} else {
				month_days.push({
					month: "current",
					value: i,
					active: false,
					date: cell_date,
					disabled: disabled,
				});
			}
		}

		// Loop to add the first dates of the next month
		for (let i = dayend; i < 6; i++) {
			let cell_date = new Date(year, month + 1, i - dayend + 1);
			let disabled = false;
			if (cell_date > new Date() || cell_date < user_created_at) {
				disabled = true;
			}
			month_days.push({
				month: "next",
				value: i - dayend + 1,
				active: false,
				date: cell_date,
				disabled: disabled,
			});
		}

		calendar_days = month_days;
	}
	function handleDropdownItemClick() {
		active = !active;
	}
	function dayClickedHandler(day) {
		let date_str = ConvertJSDateToMySQLDate(day.date);
		selected_date = day.date;
		dispatch("change", date_str);
	}
	function goPrevHandler() {
		let month = date.getMonth() - 1;
		let year = date.getFullYear();
		let day = 1;
		generateDays(new Date(year, month, day));
	}
	function goNextHandler() {
		let month = date.getMonth() + 1;
		let year = date.getFullYear();
		let day = 1;
		generateDays(new Date(year, month, day));
	}
	function goToTodayHandler() {
		selected_date = new Date();
		dispatch("change", TodayDate());
	}
	function getPrevProgressHandler() {
		selected_date.setDate(selected_date.getDate() - 1);
		let date_str = ConvertJSDateToMySQLDate(selected_date);
		dispatch("change", date_str);
	}
	function getNextProgressHandler() {
		selected_date.setDate(selected_date.getDate() + 1);
		let date_str = ConvertJSDateToMySQLDate(selected_date);
		dispatch("change", date_str);
	}
</script>

<div class="calendar_wrapper">
	<!-- svelte-ignore a11y-click-events-have-key-events -->
	<!-- svelte-ignore a11y-no-static-element-interactions -->
	<div class="calendar_navigation">
		<div
			class="calendar_nagivation_left"
			on:click={() => getPrevProgressHandler()}
			class:disabled={user &&
				selected_date <
					ConvertMySQLDateTimeToJSDateTime(user.CreatedAt)}
		>
			<ArrowDown
				disabled={user &&
					selected_date <
						ConvertMySQLDateTimeToJSDateTime(user.CreatedAt)}
			></ArrowDown>
		</div>
		<div
			class="calendar_dropdown_trigger"
			on:click={() => handleDropdownItemClick()}
		>
			{ParseDateToHumanReadableFormat(selected_date)}
		</div>
		<div
			class="calendar_nagivation_right"
			on:click={() => getNextProgressHandler()}
			class:disabled={ConvertJSDateToMySQLDate(selected_date) ===
				TodayDate()}
		>
			<ArrowDown
				disabled={ConvertJSDateToMySQLDate(selected_date) ===
					TodayDate()}
			></ArrowDown>
		</div>
	</div>

	<div class="calendar_dropdown" class:show={active}>
		<div class="calendar_dropdown_content">
			<div class="month_navigation">
				<div class="month_info">
					{months[date.getMonth()]}, {date.getFullYear()}
				</div>
				<div class="navigation_container">
					<div class="today_container">
						<!-- svelte-ignore a11y-click-events-have-key-events -->
						<!-- svelte-ignore a11y-no-static-element-interactions -->
						<div
							class="today_button"
							on:click={() => goToTodayHandler()}
						>
							Today
						</div>
					</div>
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
						class:prev_month={day.month === "previous"}
						class:next_month={day.month === "next"}
					>
						<div
							class="day"
							class:disabled={day.disabled === true}
							class:today={day.today === true}
							class:selected={selected_date &&
								selected_date.getDate() ===
									day.date.getDate() &&
								selected_date.getMonth() ==
									day.date.getMonth() &&
								selected_date.getFullYear() ==
									day.date.getFullYear()}
							on:click={() => dayClickedHandler(day)}
						>
							{day.value}
						</div>
					</div>
				{/each}
			</div>
		</div>
		<div class="dropdown_triangle"></div>
	</div>
</div>

<style>
	.calendar_navigation {
		display: flex;
		height: 60px;
		justify-content: center;
		align-items: center;
	}
	.calendar_nagivation_left {
		width: 60px;
		height: inherit;
		display: flex;
		justify-content: center;
		align-items: center;
		transform: rotate(90deg);
		cursor: pointer;
	}
	.calendar_nagivation_right {
		width: 60px;
		height: inherit;
		display: flex;
		justify-content: center;
		align-items: center;
		transform: rotate(270deg);
		cursor: pointer;
	}
	.calendar_nagivation_right.disabled {
		pointer-events: none;
		cursor: not-allowed;
	}
	.calendar_nagivation_left.disabled {
		pointer-events: none;
		cursor: not-allowed;
	}
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
		display: flex;
		justify-content: center;
		flex-wrap: nowrap;
		min-width: 400px;
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
	.calendar_days_grid {
		padding-bottom: 30px;
	}
	.calendar_week_days_grid {
		padding-top: 30px;
	}
	.day_cell {
		height: 50px;
		width: 80px;
		font-size: 16px;
		font-weight: bold;
		display: flex;
		justify-content: center;
		align-items: center;
		margin-bottom: 5px;
	}
	.weekday_cell {
		width: 80px;
		font-size: 20px;
		font-weight: bold;
		display: flex;
		justify-content: center;
		align-items: center;
		padding-bottom: 5px;
	}
	.day {
		border-radius: 50%;
		height: 50px;
		width: 50px;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.day.today {
		background-color: #3f51b5;
		color: #fff;
	}
	.day:hover {
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
	.day.selected {
		color: #fff;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		background: #0f2027;
	}
	.day.disabled {
		background-color: #ccc !important;
		color: #333 !important;
		cursor: not-allowed !important;
		pointer-events: none;
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
	.today_container {
		display: flex;
		align-items: center;
		padding-right: 20px;
	}
	.today_button {
		font-size: 14px;
		font-weight: bold;
		padding-left: 10px;
		padding-right: 10px;
		padding-top: 5px;
		padding-bottom: 5px;
		cursor: pointer;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		color: #fff;
		background-color: #673ab7ed;
		border-radius: 3px;
	}
</style>
