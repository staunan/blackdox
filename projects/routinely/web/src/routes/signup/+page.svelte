<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import SignUpPage from "components/pages/signup/SignUpPage.svelte";
</script>

<svelte:head>
	<title>Login</title>
	<meta name="login" content="Login to routinely" />
</svelte:head>
<RoutinelyPageContainer>
	<section>
		<SignUpPage></SignUpPage>
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
