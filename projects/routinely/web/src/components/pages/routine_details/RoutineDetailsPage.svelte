<script>
	import { page } from "$app/stores";
	import { onMount } from "svelte";
	import { createEventDispatcher } from "svelte";
	import { getRoutineDetails } from "apis/apis.js";
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import RoutineDetailsTab from "components/pages/routine_details/RoutineDetailsTab.svelte";
	import EditRoutineModal from "components/pages/routine_details/EditRoutineModal.svelte";
	import AboutSection from "components/pages/routine_details/AboutSection.svelte";
	import ConfirmDeleteModal from "components/pages/routine_details/ConfirmDeleteModal.svelte";
	import Right from "components/layouts/Right.svelte";
	import Card from "components/Card.svelte";

	const dispatch = createEventDispatcher();
	let routine_slug = $page.params.routine_slug;
	let editRoutineModalActive = false;
	let is_confirm_delete_modal_active = true;
	let editRoutineModalOverlayClose = true;
	let currentTabName = "about";
	let routine_details = null;

	onMount(async () => {
		let res = await getRoutineDetails({
			routine_slug: routine_slug,
		});
		if (res.HasError) {
			console.log(res);
		} else {
			routine_details = res.Data;
			console.log(routine_details);
		}
	});

	function editRoutineHandler(event) {
		editRoutineModalActive = true;
	}
	function closeEditRoutineMoalHandler(event) {
		editRoutineModalActive = false;
	}
	function deleteRoutineHandler(event) {
		is_confirm_delete_modal_active = true;
	}
	function detailsTabChangedHandler(event) {
		currentTabName = event.detail;
	}
	function routineUpdatedHandler(event) {
		routine_details = event.detail;
		console.log(routine_details);
	}
</script>

<div class="routine_details">
	{#if routine_details}
		<RoutineDetailsTab on:change={detailsTabChangedHandler}
		></RoutineDetailsTab>
		{#if currentTabName == "about"}
			<div class="about_tab">
				<AboutSection routine={routine_details}></AboutSection>
			</div>
		{:else if currentTabName == "progress"}
			<h1>Implementation Pending</h1>
		{:else if currentTabName == "history"}
			<h1>Implementation Pending</h1>
		{:else if currentTabName == "settings"}
			<div class="settings_tab">
				<div class="edit_routine_section">
					<Card padding={4} background="#0000001a">
						<h3>Edit Routine</h3>

						<p>
							Update information like title, description, time,
							days etc.
						</p>
						<Right>
							<SubmitButton
								title="Edit Routine"
								on:tap={editRoutineHandler}
							></SubmitButton>
						</Right>
					</Card>
				</div>
				<div class="delete_routine_section">
					<Card padding={4} background="#0000001a">
						<h3>Delete Routine</h3>

						<p>
							Delete this routine parmanently. Deleting routine
							will also delete all entries and histories
						</p>
						<Right>
							<SubmitButton
								color="red"
								title="Delete Routine"
								on:tap={deleteRoutineHandler}
							></SubmitButton>
						</Right>
					</Card>
				</div>
			</div>
		{/if}
		<EditRoutineModal
			active={editRoutineModalActive}
			overlayclose={editRoutineModalOverlayClose}
			routine={routine_details}
			on:close={closeEditRoutineMoalHandler}
			on:updated={routineUpdatedHandler}
		></EditRoutineModal>

		<ConfirmDeleteModal
			active={is_confirm_delete_modal_active}
			routine={routine_details}
			on:close={() => {
				is_confirm_delete_modal_active = false;
			}}
		></ConfirmDeleteModal>
	{/if}
</div>

<style>
	.about_tab {
		padding-top: 30px;
	}

	.edit_routine_section,
	.delete_routine_section {
		padding-top: 30px;
	}
</style>
