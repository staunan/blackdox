<script>
	import { createEventDispatcher } from "svelte";

	export let title = "Submit";
	export let color = "blue";
	export let disabled = false;

	const dispatch = createEventDispatcher();
	function onClick(event) {
		dispatch("tap");
	}
</script>

<button type="button" class={"btn " + color} on:click={onClick} {disabled}
	><span>{title}</span></button
>

<style>
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
	}
	.btn > * {
		position: relative;
	}
	.btn span {
		display: block;
		padding: 12px 24px;
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
		background: #fe8c00; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#f83600,
			#fe8c00
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#f83600,
			#fe8c00
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

		color: #fff;
	}
	.btn.red:hover,
	.btn.red:focus {
		background: #fe8c00; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#f83600,
			#fe8c00
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#f83600,
			#fe8c00
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
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
</style>
