<script>
	import PanelInbox from "components/pages/inbox/PanelInbox.svelte";
	import SpaceBetweenThreeItems from "components/layouts/SpaceBetweenThreeItems.svelte";
	import { onMount } from "svelte";
	import { onDestroy } from "svelte";

	let currentPanel = "inbox";
	let show_panel_dropdown = $state(false);
	let event_listener = null;

	onMount(() => {
		const funcRef = (event) => {
			if (event.target.closest(".panel_menu")) {
				show_panel_dropdown = true;
			} else if (!event.target.closest(".panel_dropdown")) {
				show_panel_dropdown = false;
			}
		};
		window.addEventListener("click", funcRef);

		// Called when component is destroyed --
		return () => {
			window.removeEventListener("click", funcRef);
		};
	});

	function panelMenuClickHandler(event) {
		show_panel_dropdown = true;
	}
</script>

<div class="inbox_page">
	<div class="panel_title">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div
			style="display: none;"
			class="panel_menu"
			onclick={panelMenuClickHandler}
		>
			<SpaceBetweenThreeItems>
				{#snippet left()}
								<div  class="red_petal"></div>
							{/snippet}
				{#snippet center()}
								<div  class="green_petal"></div>
							{/snippet}
				{#snippet right()}
								<div  class="blue_petal"></div>
							{/snippet}
			</SpaceBetweenThreeItems>
			<SpaceBetweenThreeItems>
				{#snippet left()}
								<div  class="green_petal"></div>
							{/snippet}
				{#snippet center()}
								<div  class="red_petal"></div>
							{/snippet}
				{#snippet right()}
								<div  class="blue_petal"></div>
							{/snippet}
			</SpaceBetweenThreeItems>
			<SpaceBetweenThreeItems>
				{#snippet left()}
								<div  class="red_petal"></div>
							{/snippet}
				{#snippet center()}
								<div  class="blue_petal"></div>
							{/snippet}
				{#snippet right()}
								<div  class="green_petal"></div>
							{/snippet}
			</SpaceBetweenThreeItems>
			{#if show_panel_dropdown}
				<div class="panel_dropdown"></div>
			{/if}
		</div>
		<div class="panel_title_text">
			{#if currentPanel == "inbox"}
				<span class="">Inbox (0)</span>
			{/if}
		</div>
	</div>
	<div class="panel_body">
		{#if currentPanel == "inbox"}
			<PanelInbox></PanelInbox>
		{/if}
	</div>
</div>

<style>
	.inbox_page {
		padding-top: 30px;
	}
	.panel_dropdown {
		position: absolute;
		top: 110%;
		left: 0;
		background-color: #fff;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		width: 100px;
		min-height: 50px;
	}
	.panel_title {
		padding-bottom: 20px;
		display: flex;
	}
	.panel_title_text {
		font-size: 37px;
		font-weight: bold;
	}
	.panel_menu {
		display: flex;
		flex-direction: column;
		height: 30px;
		width: 30px;
		cursor: pointer;
		position: relative;
		top: 10px;
	}
	.red_petal {
		width: 5px;
		height: 5px;
		background-color: red;
	}
	.green_petal {
		width: 5px;
		height: 5px;
		background-color: green;
	}
	.blue_petal {
		width: 5px;
		height: 5px;
		background-color: blue;
	}
</style>
