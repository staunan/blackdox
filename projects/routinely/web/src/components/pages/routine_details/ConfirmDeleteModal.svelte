<script>
	import Modal from "components/modals/Modal.svelte";
	import Center from "components/layouts/Center.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import WarningSkull from "components/svg/WarningSkull.svelte";
	import SpaceBetweenTwoItem from "components/layouts/SpaceBetweenTwoItem.svelte";
	import ErrorModal from "components/modals/ErrorModal.svelte";
	import DeleteSuccessModal from "components/pages/routine_details/DeleteSuccessModal.svelte";
	import { createEventDispatcher } from "svelte";
	import { moveToTrash } from "apis/apis.js";
	import { store } from "store";

	/**
	 * @typedef {Object} Props
	 * @property {boolean} [active]
	 * @property {any} [routine]
	 */

	/** @type {Props} */
	let { active = false, routine = null } = $props();

	let is_error_modal_active = $state(false);
	let is_delete_success_modal_active = $state(false);
	let error_modal_message = $state("");
	let overlayclose = false;
	let title = "Confirm Delete";
	let message = "Are you sure you want to delete this routine?";

	const dispatch = createEventDispatcher();
	function closeModal() {
		dispatch("close");
	}
	async function moveToTrashHandler() {
		try {
			let formData = {
				id: routine.ID,
			};
			let response = await moveToTrash(formData);
			if (response.HasError) {
				error_modal_message = response.Message;
				is_error_modal_active = true;
			} else {
				// Show Success --
				is_delete_success_modal_active = true;
				// Update Routines --
				store.moveRoutineToTrash(routine);
			}
		} catch (error) {
			console.log(error);
		}
	}
</script>

<Modal {active} , {overlayclose} on:close={closeModal}>
	<div class="modal_body">
		<Center>
			<WarningSkull></WarningSkull>
		</Center>
		<Center>
			<div class="success_title">{title}</div>
		</Center>
		<Center>
			<div class="success_message">
				{message}
			</div>
		</Center>
		<div class="button_group">
			<SpaceBetweenTwoItem>
				{#snippet left()}
								<SubmitButton
						
						title="Oops! My mistake"
						on:tap={closeModal}
						color="blue"
					></SubmitButton>
							{/snippet}
				{#snippet right()}
								<SubmitButton
						
						title="Yes, Move to trash"
						on:tap={moveToTrashHandler}
						color="red"
					></SubmitButton>
							{/snippet}
			</SpaceBetweenTwoItem>
		</div>
	</div>
</Modal>
<ErrorModal
	active={is_error_modal_active}
	overlayclose={false}
	title="Error"
	message={error_modal_message}
	buttonname="Got it"
	on:close={() => {
		is_error_modal_active = false;
		closeModal();
	}}
></ErrorModal>

<DeleteSuccessModal active={is_delete_success_modal_active}
></DeleteSuccessModal>

<style>
	.success_title {
		padding-top: 20px;
		font-size: 24px;
		font-weight: 700;
		letter-spacing: 10px;
	}
	.success_message {
		padding-top: 20px;
		font-size: 16px;
	}
	.button_group {
		padding-left: 20px;
		padding-right: 20px;
		padding-top: 30px;
		width: 100%;
	}

	.modal_body {
		padding-top: 30px;
		padding-bottom: 30px;
	}
</style>
