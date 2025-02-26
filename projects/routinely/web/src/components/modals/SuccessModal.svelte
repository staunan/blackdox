<script>
	import CloseButton from "components/buttons/CloseButton.svelte";
	import { createEventDispatcher } from "svelte";
	import { onMount } from "svelte";

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
			<!-- Success Tick -->
			<div class="checkmark-circle">
				<div class="background"></div>
				<div class="checkmark draw"></div>
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
	}
	.modal_overlay {
		background: #3f51b5c9;
		height: inherit;
		width: inherit;
	}
	.content_modal_window {
		position: absolute;
		top: 50%;
		left: 50%;
		width: 500px;
		height: 90%;
		transform: translate(-50%, -50%);
		display: flex;
		flex-direction: column;
		box-shadow: 0px 10px 15px -3px rgba(0, 0, 0, 0.1);
		background-color: #fff;
	}
	.checkmark-circle {
		width: 150px;
		height: 150px;
		position: relative;
		display: inline-block;
		vertical-align: top;
	}
	.checkmark-circle .background {
		width: 150px;
		height: 150px;
		border-radius: 50%;
		background: #1ab394;
		position: absolute;
	}
	.checkmark-circle .checkmark {
		border-radius: 5px;
	}
	.checkmark-circle .checkmark.draw:after {
		-webkit-animation-delay: 300ms;
		-moz-animation-delay: 300ms;
		animation-delay: 300ms;
		-webkit-animation-duration: 1s;
		-moz-animation-duration: 1s;
		animation-duration: 1s;
		-webkit-animation-timing-function: ease;
		-moz-animation-timing-function: ease;
		animation-timing-function: ease;
		-webkit-animation-name: checkmark;
		-moz-animation-name: checkmark;
		animation-name: checkmark;
		-webkit-transform: scaleX(-1) rotate(135deg);
		-moz-transform: scaleX(-1) rotate(135deg);
		-ms-transform: scaleX(-1) rotate(135deg);
		-o-transform: scaleX(-1) rotate(135deg);
		transform: scaleX(-1) rotate(135deg);
		-webkit-animation-fill-mode: forwards;
		-moz-animation-fill-mode: forwards;
		animation-fill-mode: forwards;
	}
	.checkmark-circle .checkmark:after {
		opacity: 1;
		height: 75px;
		width: 37.5px;
		-webkit-transform-origin: left top;
		-moz-transform-origin: left top;
		-ms-transform-origin: left top;
		-o-transform-origin: left top;
		transform-origin: left top;
		border-right: 15px solid #fff;
		border-top: 15px solid #fff;
		border-radius: 2.5px !important;
		content: "";
		left: 35px;
		top: 80px;
		position: absolute;
	}
	@-webkit-keyframes checkmark {
		0% {
			height: 0;
			width: 0;
			opacity: 1;
		}
		20% {
			height: 0;
			width: 37.5px;
			opacity: 1;
		}
		40% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
		100% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
	}
	@-moz-keyframes checkmark {
		0% {
			height: 0;
			width: 0;
			opacity: 1;
		}
		20% {
			height: 0;
			width: 37.5px;
			opacity: 1;
		}
		40% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
		100% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
	}
	@keyframes checkmark {
		0% {
			height: 0;
			width: 0;
			opacity: 1;
		}
		20% {
			height: 0;
			width: 37.5px;
			opacity: 1;
		}
		40% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
		100% {
			height: 75px;
			width: 37.5px;
			opacity: 1;
		}
	}
</style>
