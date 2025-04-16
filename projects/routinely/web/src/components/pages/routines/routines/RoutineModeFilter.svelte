<script>
	import { createEventDispatcher } from "svelte";

	export let selected = "all"; // Prop to pass the selected button value
	const dispatch = createEventDispatcher();

	const buttons = ["All", "Daily", "Weekly", "Monthly", "Yearly"];

	function selectButton(button) {
		selected = button;
		dispatch("buttonClick", { selected });
	}
</script>

<div class="button-group">
	{#each buttons as button}
		<button
			class:selected={selected === button}
			on:click={() => selectButton(button)}
		>
			{button}
		</button>
	{/each}
</div>

<style>
	.button-group {
		display: flex;
		gap: 10px;
	}

	button {
		padding: 10px 20px;
		border: none;
		border-radius: 5px;
		cursor: pointer;
		font-size: 16px;
		transition: background 0.3s;
		background: #f0f0f0;
	}

	button.selected {
		background: #6200ea;
		color: white;
	}
</style>
