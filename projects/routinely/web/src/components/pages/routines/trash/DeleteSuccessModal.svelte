<script>
	import SuccessModal from "components/modals/SuccessModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Center from "components/layouts/Center.svelte";
	import { createEventDispatcher } from "svelte";
	import { goto } from "$app/navigation";

	export let active;
	const dispatch = createEventDispatcher();

	function closeModal() {
		dispatch("close");
	}
</script>

<SuccessModal
	{active}
	overlayclose={false}
	title="Success"
	message="Routine has been successfully removed"
>
	<Center
		><SubmitButton title="Close" on:tap={closeModal} color="blue"
		></SubmitButton>
	</Center>
</SuccessModal>
