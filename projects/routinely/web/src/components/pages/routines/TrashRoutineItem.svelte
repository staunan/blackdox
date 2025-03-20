<script>
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { TodayDate } from "lib/js/datetime.js";

	export let routine = {};

	function restoreRoutineHandler(event) {
		console.log("Hello");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="routine_item" title={routine.Title}>
	<div class="routine_item_left">
		<div class="routine_title">
			{routine.Title}
		</div>
		<div class="routine_time">
			{routine.Description}
		</div>
		<div>
			<SubmitButton
				title="Restore"
				color="red"
				on:tap={restoreRoutineHandler}
			></SubmitButton>
			<SubmitButton
				title="Restore"
				color="red"
				on:tap={restoreRoutineHandler}
			></SubmitButton>
		</div>
	</div>
</div>

<style>
	.routine_item {
		display: flex;
		border-radius: 4px;
		padding: 15px;
		margin-bottom: 15px;
		font-family: monospace;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		background-color: #ff980069;
	}
	.routine_item_left {
		flex: 1;
	}
	.routine_title {
		font-size: 26px;
		font-weight: bold;
		cursor: pointer;
	}
	.routine_time {
		font-size: 18px;
	}
</style>
