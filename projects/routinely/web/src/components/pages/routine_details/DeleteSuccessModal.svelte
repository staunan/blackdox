<script>
	import SuccessModal from "components/modals/SuccessModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Center from "components/layouts/Center.svelte";
	import { createEventDispatcher } from "svelte";

	export let active;

	const dispatch = createEventDispatcher();
	function closeModal() {
		dispatch("close");
	}
</script>

<SuccessModal
	{active}
	overlayclose={true}
	title="Routine Deleted"
	message="Routine has been moved to trash"
>
	<Center
		><SubmitButton title="Got it" on:tap={closeModal} color="blue"
		></SubmitButton>
	</Center>
</SuccessModal>
