<script>
	import Modal from "components/modals/Modal.svelte";
	import Center from "components/layouts/Center.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import WarningSkull from "components/svg/WarningSkull.svelte";
	import SpaceBetweenTwoItem from "components/layouts/SpaceBetweenTwoItem.svelte";
	import ErrorModal from "components/modals/ErrorModal.svelte";
	import RestoreSuccessModal from "components/pages/routines/trash/RestoreSuccessModal.svelte";
	import { createEventDispatcher } from "svelte";
	import { restoreFromTrash } from "apis/apis.js";

	export let active = false;
	export let routine = null;

	let is_error_modal_active = false;
	let is_restore_success_modal_active = false;
	let error_modal_message = "";
	let overlayclose = false;
	let title = "Confirm Restore";
	let message =
		"Are you sure you want to restore this routine from your trash?";

	const dispatch = createEventDispatcher();
	function closeModal() {
		dispatch("close");
	}
	async function confirmRestoreHandler() {
		try {
			let formData = {
				id: routine.ID,
			};
			let response = await restoreFromTrash(formData);
			if (response.HasError) {
				error_modal_message = response.Message;
				is_error_modal_active = true;
			} else {
				// Show Success --
				is_restore_success_modal_active = true;
				dispatch("restored");
			}
		} catch (error) {
			console.log(error);
		}
	}
	function onRestoreSuccessModalCloseHandler(event) {
		is_restore_success_modal_active = false;
		closeModal();
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
				<SubmitButton
					slot="left"
					title="Oops! My mistake"
					on:tap={closeModal}
					color="blue"
				></SubmitButton>
				<SubmitButton
					slot="right"
					title="Yes, Restore it"
					on:tap={confirmRestoreHandler}
					color="red"
				></SubmitButton>
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

<RestoreSuccessModal
	active={is_restore_success_modal_active}
	on:close={onRestoreSuccessModalCloseHandler}
></RestoreSuccessModal>

<style>
	.success_title {
		padding-top: 20px;
		font-size: 24px;
		font-weight: 700;
		letter-spacing: 10px;
	}
	.success_message {
		padding: 20px;
		font-size: 16px;
		text-align: center;
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
