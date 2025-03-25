<script>
	import { ParseDateToRoutineHistory } from "lib/js/datetime.js";
	import CreateRoutineIcon from "components/svg/CreateRoutineIcon.svelte";
	import CompletedCheckmarkIcon from "components/svg/CompletedCheckmarkIcon.svelte";
	import NotCompletedIcon from "components/svg/NotCompletedIcon.svelte";
	import RichText from "components/typography/RichText.svelte";
	export let history;
</script>

{#if history}
	<div class="history_item">
		<div class="history_date">
			<RichText content={ParseDateToRoutineHistory(history.CreatedAt)}
			></RichText>
		</div>
		{#if history.HistoryType == "Routine Created"}
			<div class="icon_container" style="position: relative; top: 3px;">
				<CreateRoutineIcon></CreateRoutineIcon>
			</div>
		{:else if history.HistoryType == "Routine Checked"}
			<div class="icon_container" style="position: relative; top: 3px;">
				<CompletedCheckmarkIcon></CompletedCheckmarkIcon>
			</div>
		{:else if history.HistoryType == "Routine Unchecked"}
			<div class="icon_container" style="position: relative; top: 3px;">
				<NotCompletedIcon></NotCompletedIcon>
			</div>
		{/if}
		<div class="history_message">
			<RichText content={history.HistoryContent}></RichText>
		</div>
	</div>
{/if}

<style>
	.history_item {
		padding: 20px;
		background-color: #ddd;
		display: flex;
		height: 60px;
	}
	.icon_container {
		padding-left: 20px;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.history_date {
		font-size: 15px;
		width: 200px;
		border-right: 2px solid #ff5722;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.history_message {
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 18px;
		padding-left: 20px;
	}
</style>
