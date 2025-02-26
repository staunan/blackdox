<script>
	export let message = "Message...";
	let node;
	$: {
		if (message) {
			if (node) {
				node.innerHTML = message;
			} else {
				setTimeout(function () {
					if (node) {
						node.innerHTML = message;
					}
				}, 200);
			}
		}
	}
</script>

<div class="routine_mode_display_string" bind:this={node}></div>

<style>
	.routine_mode_display_string {
		padding-top: 10px;
		padding-bottom: 10px;
		font-size: 14px;
		font-family: monospace;
		color: green;
	}
</style>
