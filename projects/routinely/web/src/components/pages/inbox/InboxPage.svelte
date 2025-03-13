<script>
	import InboxItem from "components/pages/inbox/InboxItem.svelte";
	import { store, inboxes, routines } from "store";

	let all_inbox_items = [];
	routines.subscribe((r) => {
		console.log("Inbox Items", r);
		all_inbox_items = r;
	});
	if ($routines && $routines.length == 0) {
		store.getRoutines();
	}
</script>

<div>
	{#if all_inbox_items && all_inbox_items.length}
		{#each all_inbox_items as inboxItem}
			<InboxItem item={inboxItem}></InboxItem>
		{/each}
	{:else}
		<div class="no_item_in_inbox">
			<h1>No Items in Inbox</h1>
		</div>
	{/if}
</div>
