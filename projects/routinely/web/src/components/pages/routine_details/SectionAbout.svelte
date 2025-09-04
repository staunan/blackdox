<script>
	import { run } from 'svelte/legacy';

	import { onMount } from "svelte";
	import RoutineModeDisplayString from "components/pages/create_routine/RoutineModeDisplayString.svelte";
	import { ParseTimeToHumanReadableFormat } from "lib/js/datetime.js";

	/**
	 * @typedef {Object} Props
	 * @property {any} [routine]
	 */

	/** @type {Props} */
	let { routine = null } = $props();
	let node_routine_mode_display_string = $state("");


	function generateRoutineModeDisplayString(routine) {
		if (routine == null) {
			return "";
		}
		let str = "";
		if (routine.Mode == "Daily") {
			let selected_days = routine.DailyBasisDays.split(",");
			let days_str = "";
			if (selected_days.length === 7) {
				days_str = "<b>Everyday</b> ";
			} else {
				days_str = "Every <b>" + selected_days.join(", ") + "</b> ";
			}
			str = days_str;
		} else if (routine.Mode == "Weekly") {
			let selected_week_day = routine.WeeklyBasisWeekDays;
			let days_str = "";
			if (selected_week_day) {
				days_str = "Every <b>" + selected_week_day + "</b> ";
			} else {
				days_str = "Anyday of the week ";
			}
			str = days_str;
		} else if (routine.Mode == "Monthly") {
			let selected_month_day = routine.MonthlyBasisDate;

			let days_str = "";
			if (selected_month_day) {
				let monthday = parseInt(selected_month_day);
				if (monthday == 0) {
					days_str = "Anyday of the month ";
				} else if (monthday == 1) {
					days_str = "Every 1<sup>st</sup> of the month ";
				} else if (monthday == 2) {
					days_str = "Every 2<sup>nd</sup> of the month ";
				} else if (monthday == 3) {
					days_str = "Every 3<sup>rd</sup> of the month ";
				} else if (monthday == 32) {
					days_str = "Every end of the month ";
				} else if (monthday == 33) {
					days_str = "Yestereday of every end of the month ";
				} else {
					days_str =
						"Every " +
						selected_month_day +
						"<sup>th</sup> of the month ";
				}
			}
			str = days_str;
		} else if (routine.Mode == "Yearly") {
			let selected_month_and_date = routine.YearlyBasisMonthDate;
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
			let days_str = "";
			if (selected_month_and_date) {
				let monthday = selected_month_and_date.split("-");
				let m = parseInt(monthday[0]);
				let d = parseInt(monthday[1]);
				if (m === 0) {
					days_str = "Anyday of the year ";
				} else {
					let month_str = Months.filter((tm) => tm.value === m)[0]
						.label;
					if (d == 0) {
						days_str = "Anyday of the year";
					} else if (d == 1) {
						days_str = "Every 1<sup>st</sup> of " + month_str + " ";
					} else if (d == 2) {
						days_str = "Every 2<sup>nd</sup> of " + month_str + " ";
					} else if (d == 3) {
						days_str = "Every 3<sup>rd</sup> of " + month_str + " ";
					} else {
						days_str =
							"Every " +
							d +
							"<sup>th</sup> of " +
							month_str +
							" ";
					}
				}
			} else {
				days_str = "Anyday of the year ";
			}
			str = days_str;
		}
		return str;
	}
	run(() => {
		if (routine) {
			node_routine_mode_display_string =
				generateRoutineModeDisplayString(routine);
		}
	});
</script>

{#if routine}
	<div class="routine_title">{routine.Title}</div>
	{#if routine.Description !== "" || routine.Description !== null}
		<div class="routine_description">{routine.Description}</div>
	{/if}
	<div class="routine_mode">
		<div class="routine_mode_text">{routine.Mode}</div>
		<div class="mode_string">
			<RoutineModeDisplayString message={node_routine_mode_display_string}
			></RoutineModeDisplayString>
			<div class="time">
				at {ParseTimeToHumanReadableFormat(routine.Time)}
			</div>
		</div>
	</div>
{/if}

<style>
	.routine_title {
		font-size: 24px;
		font-weight: bold;
		padding-bottom: 20px;
	}
	.routine_description {
		font-size: 16px;
		font-weight: 500;
		padding-bottom: 20px;
	}
	.routine_mode_text {
		font-size: 30px;
		font-weight: bold;
	}
	.mode_string {
		display: flex;
		flex-wrap: nowrap;
	}
	.time {
		padding-top: 10px;
		padding-bottom: 10px;
		font-size: 14px;
		font-weight: bold;
		font-family: monospace;
		color: green;
		padding-left: 10px;
	}
</style>
