<script>
	import SpaceBetweenThreeItems from "components/layouts/SpaceBetweenThreeItems.svelte";
</script>

<SpaceBetweenThreeItems>
	<div slot="left" class="red_petal"></div>
	<div slot="center" class="green_petal"></div>
	<div slot="right" class="blue_petal"></div>
</SpaceBetweenThreeItems>
<SpaceBetweenThreeItems>
	<div slot="left" class="green_petal"></div>
	<div slot="center" class="red_petal"></div>
	<div slot="right" class="blue_petal"></div>
</SpaceBetweenThreeItems>
<SpaceBetweenThreeItems>
	<div slot="left" class="red_petal"></div>
	<div slot="center" class="blue_petal"></div>
	<div slot="right" class="green_petal"></div>
</SpaceBetweenThreeItems>

<style>
	.red_petal {
		width: 5px;
		height: 5px;
		background-color: red;
	}
	.green_petal {
		width: 5px;
		height: 5px;
		background-color: green;
	}
	.blue_petal {
		width: 5px;
		height: 5px;
		background-color: blue;
	}
</style>
