<script>
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

<div>
	<h1>My Account Page</h1>
</div>
