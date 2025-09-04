<script>
	import SuccessTick from "components/animicons/SuccessTick.svelte";
	import { markRoutineAsDone, markRoutineAsNotDone } from "apis/apis.js";
	import { createEventDispatcher } from "svelte";
	import { TodayDate } from "lib/js/datetime.js";
	import { goto } from "$app/navigation";

	let { routine = $bindable({}) } = $props();

	const dispatch = createEventDispatcher();

	function routineClickedHandler(routine) {
		goto("/routine/" + routine.Slug);
	}
	function getRoutineTimeString(time) {
		if (!time) {
			return "";
		}
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
	async function routineCheckHandler(event) {
		console.log(event.detail);
		if (event.detail === true) {
			let data = {
				routine_id: routine.ID,
				checked_on_date: TodayDate(),
			};
			try {
				let response = await markRoutineAsDone(data);
				routine.IsCompleted = true;
				routine.CheckedOnDate = response.Data.CheckedOnDate;
				routine.CompletedOn = response.Data.CreatedAt;
			} catch (error) {
				console.log(error);
			}
		} else if (event.detail === false) {
			if (routine.IsCompleted == true) {
				let data = {
					routine_id: routine.ID,
					checked_on_date: routine.CheckedOnDate,
				};
				try {
					let response = await markRoutineAsNotDone(data);
					delete routine.IsCompleted;
					delete routine.CheckedOnDate;
					delete routine.CompletedOn;
				} catch (error) {
					console.log(error);
				}
			}
		}
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="routine_item" title={routine.Title}>
	<div class="routine_item_left">
		<div
			class="routine_title"
			onclick={() => routineClickedHandler(routine)}
		>
			{routine.Title}
		</div>
		<div class="routine_time">
			{getRoutineTimeString(routine.Time)}
		</div>
	</div>
	<div class="routine_item_right">
		<SuccessTick
			checked={routine.IsCompleted === true ? true : false}
			on:change={(event) => routineCheckHandler(event)}
		></SuccessTick>
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
