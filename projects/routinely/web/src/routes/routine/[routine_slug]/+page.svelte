<script>
	import { page } from "$app/stores";
	import { onMount } from "svelte";
	import { getRoutineDetails } from "apis/apis.js";
	import RoutineDetailsPage from "components/pages/routine_details/RoutineDetailsPage.svelte";
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";

	let routine_slug = $page.params.routine_slug;
	let routine_details = null;
	onMount(async () => {
		let res = await getRoutineDetails({
			routine_slug: routine_slug,
		});
		if (res.HasError) {
			console.log(res);
		} else {
			routine_details = res.Data;
			console.log(routine_details);
		}
	});
</script>

<svelte:head>
	<title>Routine Details</title>
	<meta name="routine details" content="Routine details page" />
</svelte:head>
<RoutinelyPageContainer title="Routine Details">
	<RoutineDetailsPage data={routine_details}></RoutineDetailsPage>
</RoutinelyPageContainer>
