<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import LoginForm from "components/pages/login/LoginForm.svelte";
</script>

<svelte:head>
	<title>Login</title>
	<meta name="login" content="Login to routinely" />
</svelte:head>
<RoutinelyPageContainer>
	<section>
		<LoginForm></LoginForm>
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
