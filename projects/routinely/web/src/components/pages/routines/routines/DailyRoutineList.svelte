<script>
	import RoutineItem from "components/pages/routines/routines/RoutineItem.svelte";
	import { goto } from "$app/navigation";

	export let active = false;
	export let routines = [];

	function routineClickedHandler(event) {
		let slug = event.detail.Slug;
		goto("/routine/" + slug);
	}
</script>

{#if active}
	<div class="routine_container">
		{#each routines as routine}
			<RoutineItem {routine} on:click={routineClickedHandler}
			></RoutineItem>
		{/each}
	</div>
{/if}

<style>
	.routine_container {
		padding-top: 20px;
	}
</style>
