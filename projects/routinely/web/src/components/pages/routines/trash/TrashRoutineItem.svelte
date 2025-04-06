<script>
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import ConfirmRestoreModal from "components/pages/routines/trash/ConfirmRestoreModal.svelte";
	import ConfirmDeleteModal from "components/pages/routines/trash/ConfirmDeleteModal.svelte";
	import { TodayDate } from "lib/js/datetime.js";
	import { createEventDispatcher } from "svelte";

	export let routine = {};
	const dispatch = createEventDispatcher();
	let confirmRestoreModalActive = false;
	let confirmDeleteModalActive = false;

	function restoreRoutineHandler(event) {
		confirmRestoreModalActive = true;
	}
	function closeConfirmRestoreModalHandler(event) {
		confirmRestoreModalActive = false;
	}
	function deleteForeverRoutineHandler(event) {
		confirmDeleteModalActive = true;
	}
	function closeConfirmDeleteModalHandler(event) {
		confirmDeleteModalActive = false;
	}
	function routineRestoredHandler(event) {
		dispatch("restored");
	}
	function routineDeletedHandler(event) {
		dispatch("deleted");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="routine_item" title={routine.Title}>
	<div class="routine_item_left">
		<div class="routine_title">
			{routine.Title}
		</div>
		<div class="routine_time">
			{routine.Description}
		</div>
		<div class="button_container">
			<div class="restore_button">
				<SubmitButton
					title="Restore"
					color="green"
					on:tap={restoreRoutineHandler}
				></SubmitButton>
			</div>
			<div class="delete_forever_button">
				<SubmitButton
					title="Delete Forever"
					color="red"
					on:tap={deleteForeverRoutineHandler}
				></SubmitButton>
			</div>
		</div>
	</div>

	<ConfirmRestoreModal
		active={confirmRestoreModalActive}
		{routine}
		on:close={closeConfirmRestoreModalHandler}
		on:restored={routineRestoredHandler}
	></ConfirmRestoreModal>

	<ConfirmDeleteModal
		active={confirmDeleteModalActive}
		{routine}
		on:close={closeConfirmDeleteModalHandler}
		on:deleted={routineDeletedHandler}
	></ConfirmDeleteModal>
</div>

<style>
	.routine_item {
		display: flex;
		border-radius: 4px;
		padding: 15px;
		margin-bottom: 15px;
		font-family: monospace;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		background-color: #ff980069;
	}
	.routine_item_left {
		flex: 1;
	}
	.routine_title {
		font-size: 26px;
		font-weight: bold;
		cursor: pointer;
	}
	.routine_time {
		font-size: 18px;
	}
	.button_container {
		display: flex;
		padding-top: 20px;
	}
	.delete_forever_button {
		padding-left: 20px;
	}
</style>
