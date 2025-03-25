<script>
	import RoutineItem from "components/pages/inbox/RoutineItem.svelte";
	import { goto } from "$app/navigation";
	import { createEventDispatcher } from "svelte";

	export let item = null;
	const dispatch = createEventDispatcher();

	function routineClickedHandler(event) {
		let slug = event.detail.Slug;
		goto("/routine/" + slug);
	}
	function entryAddedHandler(event) {
		dispatch("entryadded", event.detail);
	}
	function entryRemovedHandler(event) {
		dispatch("entryremoved", event.detail);
	}
</script>

<div class="routine_container">
	<RoutineItem
		routine={item}
		on:click={routineClickedHandler}
		on:entryadded={entryAddedHandler}
		on:entryremoved={entryRemovedHandler}
	></RoutineItem>
</div>

<style>
	.routine_container {
		padding-bottom: 20px;
	}
</style>
