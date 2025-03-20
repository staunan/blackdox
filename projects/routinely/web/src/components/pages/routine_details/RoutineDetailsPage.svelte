<script>
	import { page } from "$app/stores";
	import { onMount } from "svelte";

	import { getRoutineDetails } from "apis/apis.js";
	import CarbonTab from "components/tabs/CarbonTab.svelte";

	import RoutineDetailsTab from "components/pages/routine_details/RoutineDetailsTab.svelte";

	import SectionAbout from "components/pages/routine_details/SectionAbout.svelte";
	import SectionSetting from "components/pages/routine_details/SectionSetting.svelte";

	let routine_slug = $page.params.routine_slug;

	let currentTabName = "about";
	let routine_details = null;

	onMount(async () => {
		let res = await getRoutineDetails({
			routine_slug: routine_slug,
		});
		if (res.HasError) {
			console.log(res);
		} else {
			routine_details = res.Data;
		}
	});

	function detailsTabChangedHandler(event) {
		currentTabName = event.detail;
	}
</script>

<div class="routine_details">
	{#if routine_details}
		<RoutineDetailsTab on:change={detailsTabChangedHandler}
		></RoutineDetailsTab>
		{#if currentTabName == "about"}
			<div class="about_tab">
				<SectionAbout routine={routine_details}></SectionAbout>
			</div>
		{:else if currentTabName == "progress"}
			<h1>Implementation Pending</h1>
		{:else if currentTabName == "history"}
			<h1>Implementation Pending</h1>
		{:else if currentTabName == "settings"}
			<div class="settings_tab">
				<SectionSetting routine={routine_details}></SectionSetting>
			</div>
		{/if}
	{/if}
</div>

<style>
	.about_tab {
		padding-top: 30px;
	}
</style>
