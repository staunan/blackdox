<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import DailyRoutineList from "components/DailyRoutineList.svelte";
	import WeeklyRoutineList from "components/WeeklyRoutineList.svelte";
	import MonthlyRoutineList from "components/MonthlyRoutineList.svelte";
	import YearlyRoutineList from "components/YearlyRoutineList.svelte";
	import { onMount } from "svelte";
	import { getAllRoutines, getProgress } from "apis/apis.js";

	let daily_routines = [];
	let weekly_routines = [];
	let monthly_routines = [];
	let yearly_routines = [];
	let progress = [];
	let currentTabName = "daily";
	onMount(async () => {
		try {
			let all_routines_response = await getAllRoutines({ user_id: 1 });
			let progress_response = await getProgress({ user_id: 1 });
			if (!progress_response.HasError) {
				progress = progress_response.Data;
			}

			let daily_routine_list = all_routines_response.Data.filter(
				(r) => r.Mode === "Daily"
			);
			daily_routine_list.forEach((routine) => {
				let entry = null;
				if (progress) {
					for (let i = 0; i < progress.length; i++) {
						if (progress[i].RoutineID === routine.ID) {
							entry = progress[i];
							break;
						}
					}
				}
				if (entry) {
					routine.Done = true;
					routine.DoneData = entry;
				} else {
					routine.Done = false;
					routine.DoneData = null;
				}
			});
			daily_routines = daily_routine_list;
			weekly_routines = all_routines_response.Data.filter(
				(r) => r.Mode === "Weekly"
			);
			monthly_routines = all_routines_response.Data.filter(
				(r) => r.Mode === "Monthly"
			);
			yearly_routines = all_routines_response.Data.filter(
				(r) => r.Mode === "Yearly"
			);
		} catch (error) {
			console.log(error);
		}
	});
	function tabModeChangedHandler(event) {
		currentTabName = event.detail;
		if (event.detail == "daily") {
		} else if (event.detail == "weekly") {
		} else if (event.detail == "monthly") {
		} else if (event.detail == "yearly") {
		}
	}
	function entryAddedHandler(event) {
		let entry = event.detail;
		daily_routines = daily_routines.map(function (dr) {
			if (dr.ID === entry.RoutineID) {
				return { ...dr, Done: true, DoneData: entry };
			} else {
				return dr;
			}
		});
		progress.push(entry);
	}
	function entryRemovedHandler(event) {
		let entry = event.detail;
		daily_routines = daily_routines.map(function (dr) {
			if (dr.ID === entry.RoutineID) {
				return { ...dr, Done: false, DoneData: null };
			} else {
				return dr;
			}
		});
		progress = progress.filter((p) => p.ID !== entry.ID);
	}
</script>

<svelte:head>
	<title>Routines</title>
	<meta
		name="List of routines. Explore your routines by daily, weely, monthly or yearly, create groups like 'Morning Routines', 'Evening Routine' in Daily tabs, create group like 'Christmas Routine' in Yearly tab. Check documentation for more info."
		content="List of routines"
	/>
</svelte:head>

<div class="page">
	<CarbonTab on:change={tabModeChangedHandler}></CarbonTab>
	<div class="tab_content">
		<DailyRoutineList
			active={currentTabName === "daily"}
			routines={daily_routines}
			on:entryadded={entryAddedHandler}
			on:entryremoved={entryRemovedHandler}
		></DailyRoutineList>
		<WeeklyRoutineList active={currentTabName === "weekly"}
		></WeeklyRoutineList>
		<MonthlyRoutineList active={currentTabName === "monthly"}
		></MonthlyRoutineList>
		<YearlyRoutineList active={currentTabName === "yearly"}
		></YearlyRoutineList>
	</div>
</div>

<style>
	.page {
		width: 100%;
	}
</style>
