<script>
	import ContentModal from "components/modals/ContentModal.svelte";
	import FormHeadingTitle from "components/form/FormHeadingTitle.svelte";
	import CreateRoutineForm from "components/pages/create_routine/CreateRoutineForm.svelte";
	import { createEventDispatcher } from "svelte";

	export let active = false;
	export let overlayclose = false;

	const dispatch = createEventDispatcher();

	function closeEditModalHandler() {
		dispatch("close");
	}
</script>

<ContentModal {active} {overlayclose} on:close={closeEditModalHandler}>
	<FormHeadingTitle slot="header" title="Edit Routine"></FormHeadingTitle>
	<div class="routine_form">
		<CreateRoutineForm title="Edit Routine" disableadvancesettings={true}
		></CreateRoutineForm>
	</div>
</ContentModal>

<style>
	.routine_form {
		padding: 20px;
	}
</style>
