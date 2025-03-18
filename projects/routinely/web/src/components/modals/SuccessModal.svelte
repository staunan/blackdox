<script>
	import Modal from "components/modals/Modal.svelte";

	import SubmitSuccessTick from "components/animicons/SubmitSuccessTick.svelte";
	import Center from "components/layouts/Center.svelte";
	import { createEventDispatcher } from "svelte";
	import { onMount } from "svelte";
	import "animate.css";

	export let active = false;
	export let overlayclose = false;
	export let title = "";
	export let message = "";

	const dispatch = createEventDispatcher();
	function closeModal() {
		dispatch("close");
	}
</script>

<Modal {active} {overlayclose} on:close={closeModal}>
	<div class="success_content">
		<!-- Success Tick -->
		<SubmitSuccessTick></SubmitSuccessTick>
		<Center>
			<div class="success_title">{title}</div>
		</Center>
		<Center>
			<div class="success_message">
				{message}
			</div>
		</Center>
		<div class="button_container">
			<slot></slot>
		</div>
	</div>
</Modal>

<style>
	.success_title {
		padding-top: 30px;
		font-size: 24px;
		font-weight: 700;
		letter-spacing: 10px;
	}
	.success_message {
		padding-top: 20px;
	}
	.button_container {
		padding-bottom: 20px;
		padding-top: 20px;
	}
</style>
