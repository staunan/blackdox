<script>
	export let padding = 0;
	export let background = "";
</script>

<div
	class="shadow card"
	style={"padding:" +
		(padding ? padding * 5 : 0) +
		"px;" +
		"background:" +
		(background ? background : "transparent") +
		";"}
>
	<slot></slot>
</div>

<style>
	.card {
		width: 100%;
		height: auto;
		border-radius: 4px;
		font-family: monospace;
	}
	.shadow {
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
	}
</style>
