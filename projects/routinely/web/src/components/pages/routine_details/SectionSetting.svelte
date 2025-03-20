<script>
	import ConfirmDeleteModal from "components/pages/routine_details/ConfirmDeleteModal.svelte";
	import EditRoutineModal from "components/pages/routine_details/EditRoutineModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Right from "components/layouts/Right.svelte";
	import Card from "components/Card.svelte";
	import { createEventDispatcher } from "svelte";

	export let routine;

	const dispatch = createEventDispatcher();
	let is_confirm_delete_modal_active = false;
	let editRoutineModalOverlayClose = true;
	let editRoutineModalActive = false;

	function editRoutineHandler(event) {
		editRoutineModalActive = true;
	}
	function deleteRoutineHandler(event) {
		is_confirm_delete_modal_active = true;
	}
	function routineUpdatedHandler(event) {
		routine_details = event.detail;
		console.log(routine_details);
	}
	function closeEditRoutineMoalHandler(event) {
		editRoutineModalActive = false;
	}
</script>

<div class="edit_routine_section">
	<Card padding={4} background="#0000001a">
		<h3>Edit Routine</h3>

		<p>Update information like title, description, time, days etc.</p>
		<Right>
			<SubmitButton title="Edit Routine" on:tap={editRoutineHandler}
			></SubmitButton>
		</Right>
	</Card>
</div>
<div class="delete_routine_section">
	<Card padding={4} background="#0000001a">
		<h3>Move to Trash</h3>

		<p>
			Move this routine to trash. You can restore this routine from the
			trash later on or delete this routine forever.
		</p>
		<Right>
			<SubmitButton
				color="red"
				title="Move to Trash"
				on:tap={deleteRoutineHandler}
			></SubmitButton>
		</Right>
	</Card>

	<EditRoutineModal
		active={editRoutineModalActive}
		overlayclose={editRoutineModalOverlayClose}
		{routine}
		on:close={closeEditRoutineMoalHandler}
		on:updated={routineUpdatedHandler}
	></EditRoutineModal>

	<ConfirmDeleteModal
		active={is_confirm_delete_modal_active}
		{routine}
		on:close={() => {
			is_confirm_delete_modal_active = false;
		}}
	></ConfirmDeleteModal>
</div>

<style>
	.edit_routine_section,
	.delete_routine_section {
		padding-top: 30px;
	}
</style>
