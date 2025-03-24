<script>
	import { onMount } from "svelte";
	import NoItemInHistory from "components/pages/routine_details/NoItemInHistory.svelte";
	import HistoryItem from "components/pages/routine_details/HistoryItem.svelte";
	import { getRoutineHistory } from "apis/apis.js";

	export let routine;
	let histories = [];

	onMount(async () => {
		let res = await getRoutineHistory({
			id: routine.ID,
		});
		if (res.HasError) {
			console.log(res);
		} else {
			histories = res.Data;
			console.log(histories);
		}
	});
</script>

{#if routine}
	{#if histories.length === 0}
		<div class="no_item_container">
			<NoItemInHistory></NoItemInHistory>
		</div>
	{:else}
		<div class="routine_histories">
			{#each histories as history}
				<HistoryItem {history}></HistoryItem>
			{/each}
		</div>
	{/if}
{/if}

<style>
	.routine_histories {
		padding-top: 20px;
	}
	.no_item_container {
		padding-top: 20px;
	}
</style>
