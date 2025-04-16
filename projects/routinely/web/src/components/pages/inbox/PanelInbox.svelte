<script>
	import InboxItem from "components/pages/inbox/InboxItem.svelte";
	import NoItemInInbox from "components/pages/inbox/NoItemInInbox.svelte";
	import { getInboxes } from "apis/apis.js";

	import { onMount } from "svelte";

	let inboxItems = [];
	onMount(async () => {
		await fetchInboxes();
	});

	const fetchInboxes = async () => {
		try {
			const response = await getInboxes({});
			inboxItems = response.Data;
		} catch (error) {
			console.error("Error fetching data:", error);
		}
	};

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
		{#if inboxItems && inboxItems.length}
			{#each inboxItems as inboxItem}
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
