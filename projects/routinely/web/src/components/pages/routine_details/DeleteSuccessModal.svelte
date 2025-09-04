<script>
	import SuccessModal from "components/modals/SuccessModal.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Center from "components/layouts/Center.svelte";
	import { goto } from "$app/navigation";

	let { active } = $props();

	function goToList() {
		goto("/routines");
	}
</script>

<SuccessModal
	{active}
	overlayclose={false}
	title="Success"
	message="Routine has been successfully moved to trash"
>
	<Center
		><SubmitButton title="Go to Routines" on:tap={goToList} color="blue"
		></SubmitButton>
	</Center>
</SuccessModal>
