<script>
	import SuccessModal from "components/modals/SuccessModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Center from "components/layouts/Center.svelte";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	let { active } = $props();

	function modalClose() {
		dispatch("close");
	}
</script>

<SuccessModal
	{active}
	overlayclose={false}
	title="Success"
	message="Display Image has been updated successfully!"
>
	<Center
		><SubmitButton title="Got It" on:tap={modalClose} color="blue"
		></SubmitButton>
	</Center>
</SuccessModal>
