<script>
	import SpaceBetweenThreeItems from "components/layouts/SpaceBetweenThreeItems.svelte";
</script>

<SpaceBetweenThreeItems>
	{#snippet left()}
		<div  class="red_petal"></div>
	{/snippet}
	{#snippet center()}
		<div  class="green_petal"></div>
	{/snippet}
	{#snippet right()}
		<div  class="blue_petal"></div>
	{/snippet}
</SpaceBetweenThreeItems>
<SpaceBetweenThreeItems>
	{#snippet left()}
		<div  class="green_petal"></div>
	{/snippet}
	{#snippet center()}
		<div  class="red_petal"></div>
	{/snippet}
	{#snippet right()}
		<div  class="blue_petal"></div>
	{/snippet}
</SpaceBetweenThreeItems>
<SpaceBetweenThreeItems>
	{#snippet left()}
		<div  class="red_petal"></div>
	{/snippet}
	{#snippet center()}
		<div  class="blue_petal"></div>
	{/snippet}
	{#snippet right()}
		<div  class="green_petal"></div>
	{/snippet}
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
