<script>
	import SignUpForm from "components/SignUpForm.svelte";
	import UploadProfilePictureForm from "components/UploadProfilePictureForm.svelte";
	import DotNavigation from "components/DotNavigation.svelte";

	let dots = [
		{ id: 1, title: "User Details" },
		{ id: 2, title: "Profile Picture" },
	];
	let selected_dot = dots[1];
	let user_details = null;

	function dotChangedHandler(event) {
		selected_dot = event.detail;
	}
	function userCreatedHandler(event) {
		user_details = event.detail;
		if (user_details.ID > 0) {
			selected_dot = dots[1];
		}
	}
</script>

<svelte:head>
	<title>Sign Up</title>
	<meta name="login" content="Create new account" />
</svelte:head>
<section>
	<div class="navigation_content_area">
		{#if selected_dot.id == 1}
			<SignUpForm on:created={userCreatedHandler}></SignUpForm>
		{:else if selected_dot.id == 2}
			<div class="update_profile_picture">
				<UploadProfilePictureForm></UploadProfilePictureForm>
			</div>
		{/if}
	</div>
	<div class="navigation_section">
		<DotNavigation
			{dots}
			selected={selected_dot}
			disabled={true}
			on:change={dotChangedHandler}
		></DotNavigation>
	</div>
</section>

<style>
	.navigation_content_area {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
	.navigation_section {
		padding-top: 25px;
	}
	.update_profile_picture {
		width: 80%;
	}
</style>
