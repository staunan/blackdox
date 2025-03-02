<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import { onMount } from "svelte";
	import { getAllRoutines } from "apis/apis.js";

	onMount(async () => {
		try {
			let response = await getAllRoutines({ user_id: 1 });
			console.log(response);
		} catch (error) {
			console.log(error);
		}
	});
</script>

<svelte:head>
	<title>Routines</title>
	<meta
		name="List of routines. Explore your routines by daily, weely, monthly or yearly, create groups like 'Morning Routines', 'Evening Routine' in Daily tabs, create group like 'Christmas Routine' in Yearly tab. Check documentation for more info."
		content="List of routines"
	/>
</svelte:head>

<div class="page">
	<CarbonTab></CarbonTab>
</div>

<style>
	.page {
		width: 100%;
	}
</style>
