<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";

	import { logoutUser } from "apis/apis.js";
	async function logoutHandler() {
		let response = await logoutUser();
		if (!response.HasError) {
			goto("/login");
		} else {
			console.log("Error while logging out");
			console.log(response);
		}
	}
</script>

<RoutinelyPageContainer>
	<h1>My Account Page</h1>

	<SubmitButton title="Log out" color="red" on:tap={logoutHandler}
	></SubmitButton>
</RoutinelyPageContainer>
