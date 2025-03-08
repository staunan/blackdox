<script>
	import CloseButton from "components/buttons/CloseButton.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
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
	function goToListHandler() {
		dispatch("gotolist");
	}
	function createAnotherHandler() {
		dispatch("createanother");
	}
</script>

{#if active}
	<div class="content_modal">
		<div class="modal_overlay"></div>
		<div class="content_modal_window">
			<div class="content_area animate__animated animate__tada">
				<!-- Success Tick -->
				<div class="success_tick">
					<div class="checkmark-circle">
						<div class="background"></div>
						<div class="checkmark draw"></div>
					</div>
				</div>
				<div class="success_title">Awesome!</div>
				<div class="success_message">
					Your routine has been created successfully.
				</div>
				<div class="row">
					<div class="go_to_list">
						<SubmitButton
							title="Go to list"
							on:tap={goToListHandler}
							color="red"
						></SubmitButton>
					</div>
					<div class="create_another">
						<SubmitButton
							title="Create Another"
							on:tap={createAnotherHandler}
							color="blue"
						></SubmitButton>
					</div>
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
		padding-bottom: 20px;
		display: flex;
		flex-direction: column;
	}
	.success_tick {
		width: 100%;
		display: flex;
		justify-content: center;
		align-items: center;
		padding-top: 20px;
	}
	.success_title {
		justify-content: center;
		align-items: center;
		padding-top: 30px;
		display: flex;
		width: 100%;
		font-size: 24px;
		font-weight: 700;
		letter-spacing: 10px;
	}
	.row {
		display: flex;
		width: 100%;
		align-items: center;
		justify-content: space-between;
		padding: 20px;
	}
	.go_to_list {
		flex: 1;
		display: flex;
		justify-content: start;
		align-items: center;
	}
	.create_another {
		flex: 1;
		display: flex;
		justify-content: end;
		align-items: center;
	}
	.success_message {
		display: flex;
		justify-content: center;
		align-items: center;
		padding-top: 20px;
	}
	/* Start - Checkmark Circle */
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
		-webkit-animation-delay: 500ms;
		-moz-animation-delay: 500ms;
		animation-delay: 500ms;
		-webkit-animation-duration: 2s;
		-moz-animation-duration: 2s;
		animation-duration: 2s;
		-webkit-animation-timing-function: ease-out;
		-moz-animation-timing-function: ease-out;
		animation-timing-function: ease-out;
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
	/* End - Checkmark Circle */
</style>
