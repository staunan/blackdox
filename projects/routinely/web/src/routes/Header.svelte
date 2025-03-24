<script>
	import HeaderProfileIcon from "components/pages/profile/HeaderProfileIcon.svelte";
	import HeaderNav from "components/pages/layouts/HeaderNav.svelte";
	import logo from "$lib/images/svelte-logo.svg";
	import { user_details, store } from "store";
	import { goto } from "$app/navigation";

	let user = null;
	user_details.subscribe((newValue) => {
		if (newValue && newValue.DisplayPictureName !== "") {
			user = newValue;
		}
	});

	function goToProfilePageHandler() {
		goto("/me");
	}
	function goToHomeClickHandler() {
		goto("/");
	}
</script>

<header>
	<div class="website_logo">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
		<img on:click={goToHomeClickHandler} src={logo} alt="SvelteKit" />
	</div>
	<div class="routine_nav">
		<HeaderNav></HeaderNav>
	</div>
	<div class="profile_display_photo">
		<HeaderProfileIcon {user} on:click={goToProfilePageHandler}
		></HeaderProfileIcon>
	</div>
</header>

<style>
	header {
		display: flex;
		height: 50px;
	}
	.website_logo {
		width: 60px;
		height: inherit;
		padding-left: 10px;
		padding-top: 5px;
	}
	.website_logo img {
		width: 45px;
		height: 45px;
		object-fit: contain;
		cursor: pointer;
	}
	.routine_nav {
		flex: 1;
		display: flex;
		justify-content: center;
	}
	.profile_display_photo {
		display: flex;
		height: inherit;
		align-items: center;
		padding-right: 20px;
		padding-top: 10px;
	}
</style>
