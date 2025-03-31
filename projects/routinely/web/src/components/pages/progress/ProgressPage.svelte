<script>
	import { onMount } from "svelte";
	import { getDayProgress } from "apis/apis.js";
	import { TodayDate, YesterdayDate } from "lib/js/datetime.js";
	import NotCompletedIcon from "components/svg/NotCompletedIcon.svelte";
	import CompletedCheckmarkIcon from "components/svg/CompletedCheckmarkIcon.svelte";
	import Calendar from "components/form/Calendar.svelte";
	let progressData = [];

	async function getDayProgressHandler(date) {
		let formData = {
			date: date,
		};
		let res = await getDayProgress(formData);
		if (res.HasError) {
			console.log(res);
		} else {
			progressData = res.Data;
			console.log(progressData);
		}
	}

	async function dateChangeHandler(event) {
		await getDayProgressHandler(event.detail);
	}
</script>

<div class="date_selector">
	<div class="date_input">
		<Calendar value="2025-03-29" on:change={dateChangeHandler}></Calendar>
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
		width: 200px;
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
</style>
