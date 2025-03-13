<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import RoutinesPage from "components/pages/routines/RoutinesPage.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { goto } from "$app/navigation";

	function createRoutineHandler(event) {
		goto("/create-routine");
	}
</script>

<svelte:head>
	<title>Routines</title>
	<meta
		name="routines"
		content="List of routines. Explore your routines by daily, weely, monthly or yearly, create groups like 'Morning Routines', 'Evening Routine' in Daily tabs, create group like 'Christmas Routine' in Yearly tab. Check documentation for more info."
	/>
</svelte:head>
<RoutinelyPageContainer title="Routines">
	<div slot="title_right">
		<SubmitButton title="Create Routine" on:tap={createRoutineHandler}
		></SubmitButton>
	</div>
	<RoutinesPage></RoutinesPage>
</RoutinelyPageContainer>
