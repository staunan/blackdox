<script>
	import ContentModal from "components/modals/ContentModal.svelte";
	import FormHeadingTitle from "components/form/FormHeadingTitle.svelte";
	import CreateRoutineForm from "components/pages/create_routine/CreateRoutineForm.svelte";
	import { createEventDispatcher } from "svelte";

	export let active = false;
	export let overlayclose = false;
	export let routine = null;

	const dispatch = createEventDispatcher();

	function closeEditModalHandler() {
		dispatch("close");
	}
	function routineUpdatedHandler(event) {
		dispatch("updated", event.detail);
	}
</script>

<ContentModal {active} {overlayclose} on:close={closeEditModalHandler}>
	<FormHeadingTitle slot="header" title="Edit Routine"></FormHeadingTitle>
	<div class="routine_form">
		<CreateRoutineForm
			disableadvancesettings={true}
			edit={true}
			{routine}
			on:updated={routineUpdatedHandler}
		></CreateRoutineForm>
	</div>
</ContentModal>

<style>
	.routine_form {
		padding: 20px;
	}
</style>
