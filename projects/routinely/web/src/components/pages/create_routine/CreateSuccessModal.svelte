<script>
	import SuccessModal from "components/modals/SuccessModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import SpaceBetweenTwoItem from "components/layouts/SpaceBetweenTwoItem.svelte";
	import { createEventDispatcher } from "svelte";
	import { goto } from "$app/navigation";

	export let active;

	const dispatch = createEventDispatcher();
	function closeModal() {
		goto("/routines");
	}
	function createAnotherHandler() {
		dispatch("close");
	}
</script>

<SuccessModal
	{active}
	overlayclose={true}
	title="Awesome!"
	message="Routine has been created successfully"
>
	<div class="button_group">
		<SpaceBetweenTwoItem>
			<SubmitButton
				slot="left"
				title="Go to list"
				on:tap={closeModal}
				color="red"
			></SubmitButton>
			<SubmitButton
				slot="right"
				title="Create Another"
				on:tap={createAnotherHandler}
				color="blue"
			></SubmitButton>
		</SpaceBetweenTwoItem>
	</div>
</SuccessModal>

<style>
	.button_group {
		padding-left: 20px;
		padding-right: 20px;
		width: 100%;
	}
</style>
