<script>
	import SignUpForm from "components/pages/signup/SignUpForm.svelte";
	import UploadProfilePictureForm from "components/pages/signup/UploadProfilePictureForm.svelte";
	import RegistrationSuccess from "components/pages/signup/RegistrationSuccess.svelte";
	import DotNavigation from "components/tabs/DotNavigation.svelte";
	import { getUser } from "apis/apis.js";
	import { onMount } from "svelte";

	let Dots = [
		{ id: 1, title: "User Details", completed: false },
		{ id: 2, title: "Profile Picture", completed: false },
		{ id: 3, title: "Registration Successful", completed: false },
	];
	let selected_dot = Dots[0];
	let user_details = null;

	onMount(async () => {
		let response = await getUser();
		if (response.HasError == false) {
			user_details = response.Data;
			Dots = Dots.map(function (d, index) {
				if (user_details.RegistrationStepsCompleted > index) {
					d.completed = true;
					return d;
				} else {
					d.completed = false;
					return d;
				}
			});
			if (user_details.RegistrationStepsCompleted == 1) {
				selected_dot = Dots[1];
			} else if (user_details.RegistrationStepsCompleted == 2) {
				selected_dot = Dots[2];
			}
		}
	});

	function dotChangedHandler(event) {
		selected_dot = event.detail;
	}
	function userCreatedHandler(event) {
		user_details = event.detail;
		if (user_details.ID > 0) {
			if (user_details.RegistrationStepsCompleted == 1) {
				selected_dot = Dots[1];
			}
		}
	}
	function userProfilePictureUploadHandler(event) {
		if (user_details.ID > 0) {
			user_details.RegistrationStepsCompleted = 2;
			user_details.RegistrationSuccessful = true;
			selected_dot = Dots[2];
		}
	}
	function onUploadDisplayPhotoSkippedHandler(event) {
		user_details.RegistrationStepsCompleted = 2;
		user_details.RegistrationSuccessful = true;
		selected_dot = dots[2];
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
				<UploadProfilePictureForm
					on:uploaded={userProfilePictureUploadHandler}
					on:skipped={onUploadDisplayPhotoSkippedHandler}
				></UploadProfilePictureForm>
			</div>
		{:else if selected_dot.id == 3}
			<div class="update_profile_picture">
				<RegistrationSuccess></RegistrationSuccess>
			</div>
		{/if}
	</div>
	<div class="navigation_section">
		<DotNavigation
			dots={Dots}
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
