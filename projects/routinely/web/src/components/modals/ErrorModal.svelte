<script>
	import Center from "components/layouts/Center.svelte";
	import CloseButton from "components/buttons/CloseButton.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import CircleCross from "components/animicons/CircleCross.svelte";
	import { createEventDispatcher } from "svelte";
	import { onMount } from "svelte";
	import "animate.css";

	export let active = false;
	export let message = "Error Message";
	export let title = "Title";
	export let overlayclose = false;

	const dispatch = createEventDispatcher();
	onMount(() => {
		document.addEventListener("click", function (event) {
			if (event.target.closest(".modal_overlay")) {
				if (overlayclose) {
					dispatch("close");
				}
			}
		});
	});
	function okButtonClickHandler() {
		dispatch("close");
	}
</script>

{#if active}
	<div class="content_modal">
		<div class="modal_overlay"></div>
		<div class="content_modal_window">
			<div class="content_area animate__animated animate__fadeInUp">
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
			</div>
		</div>
	</div>
{/if}

<style>
	.content_modal {
		position: fixed;
		top: 0;
		bottom: 0;
		left: 0;
		right: 0;
		width: 100%;
		height: 100vh;
		z-index: 1001;
		font-family: monospace;
	}
	.modal_overlay {
		background: #3f51b5c9;
		height: inherit;
		width: inherit;
	}
	.content_modal_window {
		position: fixed !important;
		top: 50%;
		left: 50%;
		width: 500px;
		height: auto;
		transform: translate(-50%, -50%);
		border-radius: 8px;
		z-index: 10001;
		animation-delay: 200ms;
		background: transparent;
	}
	.content_area {
		box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
		background-color: #fff;
		border-radius: 8px;
		padding-bottom: 20px;
		display: flex;
		flex-direction: column;
		padding-top: 50px;
	}
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
