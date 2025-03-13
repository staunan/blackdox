<script>
	import RoutineItem from "components/pages/inbox/RoutineItem.svelte";
	import { goto } from "$app/navigation";

	export let item = null;

	function routineClickedHandler(event) {
		let slug = event.detail.Slug;
		goto("/routine/" + slug);
	}
</script>

<div class="routine_container">
	<RoutineItem routine={item} on:click={routineClickedHandler}></RoutineItem>
</div>

<style>
	.routine_container {
		padding-top: 20px;
	}
</style>
