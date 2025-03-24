<script>
	import InboxItem from "components/pages/inbox/InboxItem.svelte";
	import NoItemInInbox from "components/pages/inbox/NoItemInInbox.svelte";
	import { store, inboxes, routines } from "store";
	import { TodayDayName } from "lib/js/datetime.js";

	let all_inbox_items = [];
	routines.subscribe((r) => {
		all_inbox_items = generateInboxItems(r);
	});
	if ($routines && $routines.length == 0) {
		store.getRoutines();
	}
	function generateInboxItems(routines) {
		let inbox_items = [];
		for (let i = 0; i < routines.length; i++) {
			if (routines[i].Mode == "Daily") {
				if (validateDailyRoutine(routines[i])) {
					inbox_items.push(routines[i]);
				}
			}
		}
		return inbox_items;
	}
	function validateDailyRoutine(routine) {
		if (routine.IsTrash == 1) {
			return false;
		}
		if (routine.Status != "active") {
			return false;
		}
		let routine_days = routine.DailyBasisDays.split(",");
		if (!routine_days.includes(TodayDayName())) {
			return false;
		}
		return true;
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
