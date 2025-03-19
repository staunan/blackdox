<script>
	import Modal from "components/modals/Modal.svelte";
	import Center from "components/layouts/Center.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import WarningSkull from "components/svg/WarningSkull.svelte";
	import SpaceBetweenTwoItem from "components/layouts/SpaceBetweenTwoItem.svelte";
	import { createEventDispatcher } from "svelte";
	import "animate.css";

	export let active = false;
	export let routine = null;
	let overlayclose = false;
	let title = "Confirm Delete";
	let message = "Are you sure you want to delete this routine?";

	const dispatch = createEventDispatcher();
	function closeModal() {
		dispatch("close");
	}
	async function updateRoutineStatus() {
		try {
			let formData = {
				id: routine.ID,
				status: "deleted",
			};
			let response = await updateRoutine(formData);
			if (response.HasError) {
				error_modal_message = response.Message;
				is_error_modal_active = true;
			} else {
				is_update_success_modal_active = true;
				store.updateRoutine(response.Data);
				dispatch("updated", response.Data);
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
				<SubmitButton
					slot="left"
					title="Oops! My mistake"
					on:tap={closeModal}
					color="blue"
				></SubmitButton>
				<SubmitButton
					slot="right"
					title="Yes Delete it"
					on:tap={deleteRoutine}
					color="red"
				></SubmitButton>
			</SpaceBetweenTwoItem>
		</div>
	</div>
</Modal>

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
