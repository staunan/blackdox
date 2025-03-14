<script>
	import CloseButton from "components/buttons/CloseButton.svelte";
	import { createEventDispatcher } from "svelte";
	import { onMount } from "svelte";
	import "animate.css";

	export let active = false;
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
	function closeModal() {
		dispatch("close");
	}
</script>

{#if active}
	<div class="content_modal">
		<div class="modal_overlay"></div>
		<div class="content_modal_window">
			<div class="content_area">
				<div class="modal_header">
					<slot name="header"></slot>
				</div>
				<div class="modal_body">
					<slot></slot>
				</div>
				<div class="modal_footer">
					<slot name="footer"></slot>
				</div>
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
		display: flex;
		flex-direction: column;
		height: calc(100vh - 100px);
	}
	.modal_body {
		flex: 1;
		overflow-y: auto;
	}
	.modal_header {
		display: flex;
		justify-content: center;
		align-items: center;
		height: 50px;
		background-color: #e91e63;
		color: #fff;
		border-top-left-radius: 8px;
		border-top-right-radius: 8px;
	}
	.modal_footer {
		padding: 10px;
	}
</style>
