<script>
	import SignUpForm from "components/SignUpForm.svelte";
</script>

<svelte:head>
	<title>Sign Up</title>
	<meta name="login" content="Create new account" />
</svelte:head>
<section>
	<SignUpForm></SignUpForm>
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
