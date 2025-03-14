<script>
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import EditRoutineModal from "components/pages/routine_details/EditRoutineModal.svelte";

	export let data = null;
	let editRoutineModalActive = false;
	let editRoutineModalOverlayClose = true;

	function editRoutineHandler(event) {
		editRoutineModalActive = true;
	}

	function closeEditRoutineMoalHandler(event) {
		editRoutineModalActive = false;
	}
</script>

<div class="routine_details">
	{#if data}
		<div class="routine_title">{data.Title}</div>

		<SubmitButton title="Edit Routine" on:tap={editRoutineHandler}
		></SubmitButton>

		<EditRoutineModal
			active={editRoutineModalActive}
			overlayclose={editRoutineModalOverlayClose}
			on:close={closeEditRoutineMoalHandler}
		></EditRoutineModal>
	{/if}
</div>
