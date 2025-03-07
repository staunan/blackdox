<script>
	import LoginForm from "components/LoginForm.svelte";
</script>

<svelte:head>
	<title>Login</title>
	<meta name="login" content="Login to routinely" />
</svelte:head>
<section>
	<LoginForm></LoginForm>
</section>

<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
</style>
