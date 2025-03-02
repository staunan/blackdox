<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import DailyRoutineList from "components/DailyRoutineList.svelte";
	import WeeklyRoutineList from "components/WeeklyRoutineList.svelte";
	import MonthlyRoutineList from "components/MonthlyRoutineList.svelte";
	import YearlyRoutineList from "components/YearlyRoutineList.svelte";
	import { onMount } from "svelte";
	import { getAllRoutines } from "apis/apis.js";

	let daily_routines = [];
	let weekly_routines = [];
	let monthly_routines = [];
	let yearly_routines = [];
	let currentTabName = "daily";
	onMount(async () => {
		try {
			let response = await getAllRoutines({ user_id: 1 });
			daily_routines = response.Data.filter((r) => r.Mode === "Daily");
			weekly_routines = response.Data.filter((r) => r.Mode === "Weekly");
			monthly_routines = response.Data.filter(
				(r) => r.Mode === "Monthly"
			);
			yearly_routines = response.Data.filter((r) => r.Mode === "Yearly");
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
