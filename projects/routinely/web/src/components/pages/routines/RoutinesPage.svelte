<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import DailyRoutineList from "components/pages/routines/DailyRoutineList.svelte";
	import WeeklyRoutineList from "components/pages/routines/WeeklyRoutineList.svelte";
	import MonthlyRoutineList from "components/pages/routines/MonthlyRoutineList.svelte";
	import YearlyRoutineList from "components/pages/routines/YearlyRoutineList.svelte";

	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import { getProgress } from "apis/apis.js";
	import { routines, store } from "store";

	let all_routines = [];
	let daily_routines = [];
	let weekly_routines = [];
	let monthly_routines = [];
	let yearly_routines = [];
	let progress = [];
	let currentTabName = "daily";

	routines.subscribe((v) => {
		if (v == null) {
			all_routines = [];
		} else {
			all_routines = v;
		}
		routinesChanged();
	});

	onMount(async () => {
		try {
			// Get Routines --
			if ($routines == null || $routines.length == 0) {
				await store.getRoutines();
			}

			let progress_response = await getProgress({ user_id: 1 });
			if (!progress_response.HasError) {
				if (progress_response.Data == null) {
					progress = [];
				} else {
					progress = progress_response.Data;
				}
			}
		} catch (error) {
			console.log(error);
		}
	});
	function routinesChanged() {
		let daily_routine_list = all_routines.filter((r) => r.Mode === "Daily");
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
		weekly_routines = all_routines.filter((r) => r.Mode === "Weekly");
		monthly_routines = all_routines.filter((r) => r.Mode === "Monthly");
		yearly_routines = all_routines.filter((r) => r.Mode === "Yearly");
	}
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
	function createRoutineHandler(event) {
		goto("/create-routine");
	}
</script>

<div class="routines_container">
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
		<YearlyRoutineList
			routines={yearly_routines}
			active={currentTabName === "yearly"}
		></YearlyRoutineList>
	</div>
</div>
