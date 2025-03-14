<script>
	export let data = null;
	import SubmitButton from "components/buttons/SubmitButton.svelte";

	function createRoutineHandler(event) {}
</script>

<div class="routine_details">
	{#if data}
		<div class="routine_title">{data.Title}</div>

		<SubmitButton title="Edit Routine" on:tap={createRoutineHandler}
		></SubmitButton>
	{/if}
</div>
