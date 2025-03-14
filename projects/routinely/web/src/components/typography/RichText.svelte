<script>
	export let content = "";
	let node;
	$: {
		if (content) {
			if (node) {
				node.innerHTML = content;
			} else {
				setTimeout(function () {
					if (node) {
						node.innerHTML = content;
					}
				}, 200);
			}
		}
	}
</script>

<div class="rich_text" bind:this={node}></div>
