<script>
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { goto } from "$app/navigation";

	function goToCreateRoutinesHandler(event) {
		goto("create-routine");
	}
</script>

<div class="no_item_in_inbox">
	<h3>Inbox Empty</h3>
	<p>
		You do not have any routine. Click on the following button to create a
		new routine:
	</p>
	<SubmitButton
		title="Create Routine"
		on:tap={goToCreateRoutinesHandler}
		color="blue"
	></SubmitButton>
</div>

<style>
	.no_item_in_inbox {
		display: flex;
		flex-direction: column;
		width: 100%;
		height: 300px;
		border-radius: 4px;
		background-color: #ddd;
		font-size: 20px;
		font-weight: bold;
		color: #333;
		justify-content: center;
		align-items: center;
	}
</style>
