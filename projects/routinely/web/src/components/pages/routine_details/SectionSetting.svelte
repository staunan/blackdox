<script>
	import ConfirmDeleteModal from "components/pages/routine_details/ConfirmDeleteModal.svelte";
	import EditRoutineModal from "components/pages/routine_details/EditRoutineModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import EditRoutineIcon from "components/svg/EditRoutineIcon.svelte";
	import TrashIcon from "components/svg/TrashIcon.svelte";
	import Right from "components/layouts/Right.svelte";
	import Card from "components/Card.svelte";
	import { createEventDispatcher } from "svelte";

	let { routine } = $props();

	const dispatch = createEventDispatcher();
	let is_confirm_delete_modal_active = $state(false);
	let editRoutineModalOverlayClose = true;
	let editRoutineModalActive = $state(false);

	function editRoutineHandler(event) {
		editRoutineModalActive = true;
	}
	function deleteRoutineHandler(event) {
		is_confirm_delete_modal_active = true;
	}
	function routineUpdatedHandler(event) {
		dispatch("updated", event.detail);
	}
	function closeEditRoutineMoalHandler(event) {
		editRoutineModalActive = false;
	}
</script>

<div class="edit_routine_section">
	<Card padding={4} background="#0000001a">
		<div class="section_title">
			<EditRoutineIcon></EditRoutineIcon>
			<div class="section_title_text">Edit Routine</div>
		</div>
		<p class="section_title_desc">
			Update information like title, description, time, days etc.
		</p>
		<Right>
			<SubmitButton title="Edit Routine" on:tap={editRoutineHandler}
			></SubmitButton>
		</Right>
	</Card>
</div>
<div class="delete_routine_section">
	<Card padding={4} background="#0000001a">
		<div class="section_title">
			<TrashIcon></TrashIcon>
			<div class="section_title_text">Move to Trash</div>
		</div>
		<p class="section_title_desc">
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
	.section_title {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		width: auto;
	}
	.section_title_text {
		padding-left: 20px;
		font-size: 24px;
		font-weight: bold;
	}
	.section_title_desc {
		font-size: 20px;
	}
</style>
