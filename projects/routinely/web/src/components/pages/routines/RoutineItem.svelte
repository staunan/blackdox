<script>
	import SuccessTick from "components/animicons/SuccessTick.svelte";
	import { markRoutineAsDone, markRoutineAsNotDone } from "apis/apis.js";
	import { createEventDispatcher } from "svelte";
	import { TodayDate } from "lib/js/datetime.js";

	export let routine = {};

	const dispatch = createEventDispatcher();
	function routineClickedHandler(routine) {
		dispatch("click", routine);
	}
	function getRoutineTimeString(time) {
		let time_arr = time.split(":");
		let zone = "";
		let hour = 0;
		if (Number(time_arr[0]) < 12) {
			zone = "AM";
			hour = Number(time_arr[0]);
		} else {
			zone = "PM";
			hour = Number(time_arr[0]) - 12;
		}
		return hour + ":" + time_arr[1] + " " + zone;
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="routine_item" title={routine.Title}>
	<div class="routine_item_left">
		<div
			class="routine_title"
			on:click={() => routineClickedHandler(routine)}
		>
			{routine.Title}
		</div>
		<div class="routine_time">
			{getRoutineTimeString(routine.Time)}
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
	.routine_item_right {
		display: flex;
		justify-content: center;
		align-items: center;
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
