<script>
	import InboxItem from "components/pages/inbox/InboxItem.svelte";
	import NoItemInInbox from "components/pages/inbox/NoItemInInbox.svelte";
	import { store, inboxes } from "store";

	import { onMount } from "svelte";

	onMount(async () => {
		try {
			if ($inboxes == null || $inboxes.length == 0) {
				await store.init();
			}
		} catch (error) {
			console.log(error);
		}
	});

	function entryAddedHandler(event) {
		let entry = event.detail;
		store.addEntry(entry);
	}
	function entryRemovedHandler(event) {
		let entry = event.detail;
		store.removeEntry(entry);
	}
</script>

<div class="inbox_container">
	<div class="inbox_items">
		{#if $inboxes && $inboxes.length}
			{#each $inboxes as inboxItem}
				<InboxItem
					item={inboxItem}
					on:entryadded={entryAddedHandler}
					on:entryremoved={entryRemovedHandler}
				></InboxItem>
			{/each}
		{:else}
			<NoItemInInbox></NoItemInInbox>
		{/if}
	</div>
</div>
