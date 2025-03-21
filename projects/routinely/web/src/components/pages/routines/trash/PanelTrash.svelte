<script>
	import { routines } from "store";
	import { goto } from "$app/navigation";
	import TrashRoutineItem from "components/pages/routines/trash/TrashRoutineItem.svelte";
	import NoItemInTrash from "components/pages/routines/trash/NoItemInTrash.svelte";

	let deleted_routines = [];
	routines.subscribe((v) => {
		deleted_routines = v.filter((item) => item.IsTrash === 1);
	});
</script>

<div class="routine_container">
	{#if deleted_routines.length === 0}
		<NoItemInTrash></NoItemInTrash>
	{/if}
	{#each deleted_routines as routine}
		<TrashRoutineItem {routine}></TrashRoutineItem>
	{/each}
</div>

<style>
	.routine_container {
		padding-top: 20px;
	}
</style>
