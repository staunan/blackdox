<script>
	import CarbonTab from "components/tabs/CarbonTab.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import RoutineDetailsTab from "components/pages/routine_details/RoutineDetailsTab.svelte";
	import EditRoutineModal from "components/pages/routine_details/EditRoutineModal.svelte";
	import AboutSection from "components/pages/routine_details/AboutSection.svelte";
	import Right from "components/layouts/Right.svelte";
	import Card from "components/Card.svelte";

	export let data = null;
	let editRoutineModalActive = false;
	let editRoutineModalOverlayClose = true;
	let currentTabName = "about";

	function editRoutineHandler(event) {
		editRoutineModalActive = true;
	}

	function closeEditRoutineMoalHandler(event) {
		editRoutineModalActive = false;
	}
	function detailsTabChangedHandler(event) {
		currentTabName = event.detail;
		if (currentTabName == "about") {
			generateRoutineModeDisplayString(data);
		}
	}
</script>

<div class="routine_details">
	{#if data}
		<RoutineDetailsTab on:change={detailsTabChangedHandler}
		></RoutineDetailsTab>
		{#if currentTabName == "about"}
			<div class="about_tab">
				<AboutSection routine={data}></AboutSection>
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
								on:tap={editRoutineHandler}
							></SubmitButton>
						</Right>
					</Card>
				</div>
			</div>
		{/if}
		<EditRoutineModal
			active={editRoutineModalActive}
			overlayclose={editRoutineModalOverlayClose}
			routine={data}
			on:close={closeEditRoutineMoalHandler}
		></EditRoutineModal>
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
