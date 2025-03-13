<script>
	import { user_details, store } from "store";
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	let currentPageUrl = "";
	let restricted_routes = ["/create-routine", "/inbox", "/me", "/routines"];

	user_details.subscribe((value) => {
		if (value && value.ID) {
			loggedInHandler();
		} else {
			// User is not Logged In server --
			// User should be redirect to login page --
			notLoggedInHandler();
		}
	});

	onMount(async () => {
		currentPageUrl = $page.url.pathname;
		if ($user_details && $user_details.ID) {
			// User Logged In --
			loggedInHandler();
		} else {
			// Check if user is logged in server --
			try {
				await store.getUser();
			} catch (err) {
				// User is not Logged In server --
				// User should be redirect to login page --
				notLoggedInHandler();
			}
		}
	});
	function notLoggedInHandler() {
		if (restricted_routes.includes(currentPageUrl)) {
			goto("/login");
		}
	}
	function loggedInHandler() {
		if (currentPageUrl == "/login") {
			// Logged in page should be blocked --
			goto("/inbox");
		}
	}
</script>

<div>
	<slot></slot>
</div>
