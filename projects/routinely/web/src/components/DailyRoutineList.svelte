<script>
	import SuccessTick from "components/checkmark/SuccessTick.svelte";
	import { TodayDate } from "lib/js/datetime.js";
	import { markRoutineAsDone } from "apis/apis.js";

	export let active = false;
	export let routines = [];

	function routineClickedHandler(routine) {
		console.log(routine);
	}
	async function routineCheckHandler(event, routine) {
		if (event.detail === true) {
			let data = {
				routine_id: routine.ID,
				checked_on_date: TodayDate(),
			};

			try {
				let response = await markRoutineAsDone(data);
				console.log(response);
			} catch (error) {
				console.log(error);
			}
		}
	}
</script>

{#if active}
	<h1>Daily Routine List</h1>
	<div class="routine_container">
		{#each routines as routine}
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
					<div class="routine_time">{routine.Time}</div>
				</div>
				<div class="routine_item_right">
					<SuccessTick
						checked={false}
						on:change={(event) =>
							routineCheckHandler(event, routine)}
					></SuccessTick>
				</div>
			</div>
		{/each}
	</div>
{/if}

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
