<script>
	import { store } from "store";
	import { onMount } from "svelte";
	import { page } from "$app/stores";
	import { goto } from "$app/navigation";

	let currentPageUrl = "";
	let user_details = store.user_details;

	onMount(async () => {
		currentPageUrl = $page.url.pathname;
		if (store.initialized && store.user_details) {
			// User Logged In --
			validatePage();
		} else {
			// Check if user is logged in server --
			if (await store.initialize()) {
				// User logged in server --
				user_details = store.user_details;
				if (user_details && user_details.ID) {
					validatePage();
				}
			} else {
				// User is not Logged In server --
				// User should be redirect to login page --
				if (currentPageUrl == "/inbox") {
					goto("/login");
				}
			}
		}
	});
	function validatePage() {
		if (currentPageUrl == "/login") {
			// Logged in page should be blocked --
			goto("/inbox");
		}
	}
</script>

<div>
	<slot></slot>
</div>
