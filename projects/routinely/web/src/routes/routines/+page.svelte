<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import RoutinesPage from "components/pages/routines/RoutinesPage.svelte";
</script>

<svelte:head>
	<title>Routines</title>
	<meta name="routines" content="All Routines" />
</svelte:head>
<RoutinelyPageContainer>
	<section>
		<RoutinesPage></RoutinesPage>
	</section>
</RoutinelyPageContainer>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
</style>
