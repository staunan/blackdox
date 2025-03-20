<script>
	import { routines } from "store";
	import { goto } from "$app/navigation";
	import TrashRoutineItem from "components/pages/routines/TrashRoutineItem.svelte";

	let deleted_routines = [];
	routines.subscribe((v) => {
		deleted_routines = v.filter((item) => item.Status == "deleted");
	});
</script>

<div class="routine_container">
	{#each deleted_routines as routine}
		<TrashRoutineItem {routine}></TrashRoutineItem>
	{/each}
</div>

<style>
	.routine_container {
		padding-top: 20px;
	}
</style>
