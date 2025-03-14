<script>
	import Modal from "components/modals/Modal.svelte";
	import Center from "components/layouts/Center.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import CircleCross from "components/animicons/CircleCross.svelte";
	import { createEventDispatcher } from "svelte";
	import "animate.css";

	export let active = false;
	export let message = "Error Message";
	export let title = "Title";
	export let overlayclose = false;

	const dispatch = createEventDispatcher();
	function okButtonClickHandler() {
		dispatch("close");
	}
</script>

<Modal {active} , {overlayclose} on:close={okButtonClickHandler}>
	<!-- Success Tick -->
	<Center>
		<CircleCross></CircleCross>
	</Center>
	<Center>
		<div class="success_title">{title}</div>
	</Center>
	<Center>
		<div class="success_message">
			{message}
		</div>
	</Center>
	<Center>
		<div class="create_another">
			<SubmitButton
				title="Got It"
				on:tap={okButtonClickHandler}
				color="blue"
			></SubmitButton>
		</div>
	</Center>
</Modal>

<style>
	.success_title {
		padding-top: 20px;
		font-size: 24px;
		font-weight: 700;
		letter-spacing: 10px;
	}
	.create_another {
		padding-top: 30px;
		padding-bottom: 20px;
	}
	.success_message {
		padding-top: 20px;
	}
</style>
