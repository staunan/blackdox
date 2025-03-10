<script>
	import Card from "components/Card.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { goto } from "$app/navigation";
	import "animate.css";

	function goToInboxHnalder() {
		goto("/inbox");
	}
</script>

<Card>
	<!-- Success Tick -->
	<div class="success_tick">
		<div class="checkmark-circle">
			<div class="background"></div>
			<div class="checkmark draw"></div>
		</div>
	</div>
	<div class="success_title">Awesome!</div>
	<div class="success_message">
		Your Profile has been created successfully.
	</div>
	<div class="go_to_inbox">
		<SubmitButton title="Go to Inbox" on:tap={goToInboxHnalder}
		></SubmitButton>
	</div>
</Card>

<style>
	.go_to_inbox {
		display: flex;
		justify-content: center;
		align-items: center;
		padding-bottom: 50px;
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
	.success_message {
		display: flex;
		justify-content: center;
		align-items: center;
		padding-top: 50px;
		padding-bottom: 50px;
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
	.success_message {
		display: flex;
		justify-content: center;
		align-items: center;
		padding-top: 20px;
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
