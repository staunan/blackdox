<script>
	import { page } from "$app/stores";
	import { onMount } from "svelte";
	import { getRoutineDetails } from "apis/apis.js";
	let routine_slug = $page.params.routine_slug;
	let routine_details = null;
	onMount(async () => {
		routine_details = await getRoutineDetails({
			routine_slug: routine_slug,
		});
		console.log(routine_details);
	});
</script>

<h1>Routine Details</h1>
