<script>
	import { onMount } from "svelte";
	import { getDayProgress } from "apis/apis.js";
	import { TodayDate, YesterdayDate } from "lib/js/datetime.js";
	import NotCompletedIcon from "components/svg/NotCompletedIcon.svelte";
	import CompletedCheckmarkIcon from "components/svg/CompletedCheckmarkIcon.svelte";
	import Calendar from "components/form/Calendar.svelte";

	let selectedDate = TodayDate();
	let progressData = [];

	async function getDayProgressHandler(date) {
		let formData = {
			date: date,
		};
		let res = await getDayProgress(formData);
		if (res.HasError) {
			console.log(res);
		} else {
			if (res.Data == null) {
				progressData = [];
			} else {
				progressData = res.Data;
			}
		}
	}

	async function dateChangeHandler(event) {
		selectedDate = event.detail;
		await getDayProgressHandler(selectedDate);
	}

	onMount(async () => {
		selectedDate = TodayDate();
		await getDayProgressHandler(selectedDate);
	});
</script>

<div class="date_selector">
	<div class="date_input">
		<Calendar value={selectedDate} on:change={dateChangeHandler}></Calendar>
	</div>
</div>

<div class="progress_items">
	{#if progressData && progressData.length}
		{#each progressData as progressItem, index}
			<div class="progress_item">
				<div class="item_title">
					{index + 1}. {progressItem.RoutineTitle}
				</div>
				{#if progressItem.IsCompleted == true}
					<div class="item_checked_status">
						<CompletedCheckmarkIcon></CompletedCheckmarkIcon>
					</div>
				{:else}
					<div class="item_checked_status">
						<NotCompletedIcon></NotCompletedIcon>
					</div>
				{/if}
			</div>
		{/each}
	{:else}
		<div class="day_missed">
			<div class="day_missed_icon">
				<NotCompletedIcon></NotCompletedIcon>
			</div>
			<div class="day_missed_text">Day Missed</div>
		</div>
	{/if}
</div>

<style>
	.date_selector {
		display: flex;
		justify-content: center;
		align-items: center;
		margin-top: 20px;
	}
	.date_input {
		display: flex;
		justify-content: center;
	}
	.progress_items {
		border: 1px solid #ccc;
		border-bottom: none;
		margin-top: 20px;
	}
	.progress_item {
		display: flex;
		padding: 10px;
		font-size: 18px;
		font-weight: bold;
		border-bottom: 1px solid #ccc;
		display: flex;
	}
	.item_title {
		flex: 1;
	}
	.day_missed {
		display: flex;
		justify-content: center;
		align-items: center;
		flex-direction: column;
		padding: 50px;
		background-color: #ddd;
	}
	.day_missed {
		border-bottom: 1px solid #ccc;
	}
	.day_missed_text {
		font-size: 24px;
		font-weight: bold;
	}
</style>
