<script>
	let tabs = [
		{
			id: 1,
			label: "Daily",
			id: 2,
			label: "Weekly",
			id: 3,
			label: "Monthly",
			id: 4,
			label: "Yearly",
		},
	];
	let selectedTab = tabs[0];
	function tabClickedHandler(event) {
		selectedTab = tab;
	}
</script>

<div class="tabs">
	{#each tabs as tab}
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y-no-static-element-interactions -->
		<div
			class="tab_item"
			class:selected={selectedTab.id == tab.id}
			on:click={() => tabClickedHandler(tab)}
			title={tab.label}
		>
			{tab.label}
		</div>
	{/each}
</div>
<div class="tab_content_one"></div>
<div class="tab_content_two"></div>
<div class="tab_content_three"></div>
<div class="tab_content_four"></div>

<style>
	.tabs {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		font-family: monospace;
	}
	.tab_item {
		padding-right: 30px;
		font-size: 16px;
		padding-top: 3;
	}
</style>
