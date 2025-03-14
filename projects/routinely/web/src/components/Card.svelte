<script>
	export let padding = 0;
</script>

<div class="shadow card" style={"padding:" + padding * 5 + "px;"}>
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
