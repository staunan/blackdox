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
</script>

<header>
	<div class="website_logo">
		<a href="https://svelte.dev/docs/kit">
			<img src={logo} alt="SvelteKit" />
		</a>
	</div>
	<HeaderNav></HeaderNav>
	<div class="profile_display_photo">
		<HeaderProfileIcon {user} on:click={goToProfilePageHandler}
		></HeaderProfileIcon>
	</div>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
		height: 50px;
	}

	.website_logo {
		width: 60px;
		height: inherit;
	}

	.website_logo a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	.website_logo img {
		width: 2em;
		height: 2em;
		object-fit: contain;
	}
	.profile_display_photo {
		display: flex;
		width: 45px;
		height: inherit;
		align-items: center;
	}
</style>
