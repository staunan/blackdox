<script>
	import InboxItem from "components/pages/inbox/InboxItem.svelte";
	import NoItemInInbox from "components/pages/inbox/NoItemInInbox.svelte";
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

<div class="inbox_container">
	<div class="inbox_items">
		{#if all_inbox_items && all_inbox_items.length}
			{#each all_inbox_items as inboxItem}
				<InboxItem item={inboxItem}></InboxItem>
			{/each}
		{:else}
			<NoItemInInbox></NoItemInInbox>
		{/if}
	</div>
</div>
