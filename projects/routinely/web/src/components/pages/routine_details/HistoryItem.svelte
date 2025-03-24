<script>
	export let history;
</script>

{#if history}
	<div class="history_item">
		<div class="history_date">{history.CreatedAt}</div>
		<div class="history_message">{history.HistoryContent}</div>
	</div>
{/if}

<style>
	.history_item {
		padding: 20px;
		border-radius: 5px;
		margin-bottom: 10px;
		background-color: #ddd;
	}
	.history_date {
		font-size: 15px;
	}
	.history_message {
		font-size: 18px;
		font-weight: bold;
	}
</style>
