<script>
	import RoutineItem from "components/pages/routines/routines/RoutineItem.svelte";
	import NoItemRoutineList from "components/pages/routines/routines/NoItemRoutineList.svelte";
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
		{#if routines && routines.length == 0}
			<NoItemRoutineList></NoItemRoutineList>
		{:else}
			{#each routines as routine}
				<RoutineItem {routine} on:click={routineClickedHandler}
				></RoutineItem>
			{/each}
		{/if}
	</div>
{/if}

<style>
	.routine_container {
		padding-top: 20px;
	}
</style>
