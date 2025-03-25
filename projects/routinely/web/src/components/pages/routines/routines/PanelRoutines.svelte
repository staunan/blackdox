<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import DailyRoutineList from "components/pages/routines/routines/DailyRoutineList.svelte";
	import WeeklyRoutineList from "components/pages/routines/routines/WeeklyRoutineList.svelte";
	import MonthlyRoutineList from "components/pages/routines/routines/MonthlyRoutineList.svelte";
	import YearlyRoutineList from "components/pages/routines/routines/YearlyRoutineList.svelte";

	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
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
				await store.init();
			}
		} catch (error) {
			console.log(error);
		}
	});
	function routinesChanged() {
		daily_routines = all_routines.filter(
			(r) => r.Mode === "Daily" && r.IsTrash === 0
		);
		weekly_routines = all_routines.filter(
			(r) => r.Mode === "Weekly" && r.IsTrash === 0
		);
		monthly_routines = all_routines.filter(
			(r) => r.Mode === "Monthly" && r.IsTrash === 0
		);
		yearly_routines = all_routines.filter(
			(r) => r.Mode === "Yearly" && r.IsTrash === 0
		);
	}
	function tabModeChangedHandler(event) {
		currentTabName = event.detail;
		if (event.detail == "daily") {
		} else if (event.detail == "weekly") {
		} else if (event.detail == "monthly") {
		} else if (event.detail == "yearly") {
		}
	}
</script>

<div class="routines_container">
	<CarbonTab on:change={tabModeChangedHandler}></CarbonTab>
	<div class="tab_content">
		<DailyRoutineList
			active={currentTabName === "daily"}
			routines={daily_routines}
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
