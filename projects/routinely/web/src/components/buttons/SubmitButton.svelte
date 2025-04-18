<script>
	import { createEventDispatcher } from "svelte";
	const hasSlot = $$slots.default;

	export let title = "Submit";
	export let color = "blue";
	export let disabled = false;

	const dispatch = createEventDispatcher();
	function onClick(event) {
		dispatch("tap");
	}
</script>

<button type="button" class={"btn " + color} on:click={onClick} {disabled}>
	<div class="button_content">
		{#if hasSlot}
			<div class="button_icon">
				<slot />
			</div>
		{/if}
		<div class="button_text">{title}</div>
	</div>
</button>

<style>
	.button_icon {
		padding-right: 15px;
	}
	.button_content {
		display: flex;
		justify-content: center;
		align-items: center;
		padding-left: 20px;
		padding-right: 20px;
		padding-top: 5px;
		padding-bottom: 5px;
	}
	button:disabled {
		background-color: #ccc !important;
		color: #333 !important;
	}
	.btn {
		position: relative;
		display: block;
		padding: 0;
		overflow: hidden;
		border-width: 0;
		outline: none;
		border-radius: 2px;
		transition: background-color 0.3s;
		cursor: pointer;
		font-size: 14px;
		font-family: monospace;
		height: 50px;
	}
	.btn > * {
		position: relative;
	}
	.btn:before {
		content: "";
		position: absolute;
		top: 50%;
		left: 50%;
		display: block;
		width: 0;
		padding-top: 0;
		border-radius: 100%;
		background-color: rgba(236, 240, 241, 0.3);
		-webkit-transform: translate(-50%, -50%);
		-moz-transform: translate(-50%, -50%);
		-ms-transform: translate(-50%, -50%);
		-o-transform: translate(-50%, -50%);
		transform: translate(-50%, -50%);
	}
	.btn:active:before {
		width: 120%;
		padding-top: 120%;
		transition:
			width 0.2s ease-out,
			padding-top 0.2s ease-out;
	}
	/* Colors */
	/* Blue */
	.btn.blue {
		box-shadow: 0 1px 4px rgba(0, 0, 0, 0.6);
		background-color: #673ab7;
		color: #fff;
	}
	.btn.blue:hover,
	.btn.blue:focus {
		background-color: #3f51b5;
	}
	/* Red */
	.btn.red {
		box-shadow: 0 1px 4px rgba(0, 0, 0, 0.6);
		background-color: #d32f2f;
		color: #fff;
	}
	.btn.red:hover,
	.btn.red:focus {
		background-color: #d32f2f;
		color: #fff;
	}
	/* Green */
	.btn.green {
		box-shadow: 0 1px 4px rgba(0, 0, 0, 0.6);
		background: #52c234; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#061700,
			#52c234
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#061700,
			#52c234
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

		color: #fff;
	}
	.btn.green:hover,
	.btn.green:focus {
		background: #52c234; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#061700,
			#52c234
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#061700,
			#52c234
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
	}
	/* GreyBlue */
	.btn.greyblue {
		box-shadow: 0 1px 4px rgba(0, 0, 0, 0.6);
		background-color: #0000001a;
		color: #000;
	}
	.btn.greyblue:hover,
	.btn.greyblue:focus {
		background-color: #5e35b1 !important;
		color: #fff !important;
	}
</style>
