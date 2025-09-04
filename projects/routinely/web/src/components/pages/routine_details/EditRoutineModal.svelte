<script>
	import ContentModal from "components/modals/ContentModal.svelte";
	import FormHeadingTitle from "components/form/FormHeadingTitle.svelte";
	import CreateRoutineForm from "components/pages/create_routine/CreateRoutineForm.svelte";
	import { createEventDispatcher } from "svelte";

	/**
	 * @typedef {Object} Props
	 * @property {boolean} [active]
	 * @property {boolean} [overlayclose]
	 * @property {any} [routine]
	 */

	/** @type {Props} */
	let { active = false, overlayclose = false, routine = null } = $props();

	const dispatch = createEventDispatcher();

	function closeEditModalHandler() {
		dispatch("close");
	}
	function routineUpdatedHandler(event) {
		dispatch("updated", event.detail);
	}
</script>

<ContentModal {active} {overlayclose} on:close={closeEditModalHandler}>
	{#snippet header()}
		<FormHeadingTitle  title="Edit Routine"></FormHeadingTitle>
	{/snippet}
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
