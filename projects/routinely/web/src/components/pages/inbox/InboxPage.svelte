<script>
	import PanelInbox from "./PanelInbox.svelte";
	import SpaceBetweenThreeItems from "../../layouts/SpaceBetweenThreeItems.svelte";
	import { onMount } from "svelte";

	let currentPanel = "inbox";
	let show_panel_dropdown = false;

	onMount(() => {
		document.addEventListener("click", function (event) {
			console.log("Event Listening");
			if (event.target.closest(".panel_menu")) {
				show_panel_dropdown = true;
			} else if (!event.target.closest(".panel_dropdown")) {
				show_panel_dropdown = false;
			}
		});
	});

	function panelMenuClickHandler(event) {
		show_panel_dropdown = true;
	}
</script>

<div class="inbox_page">
	<div class="panel_title">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="panel_menu" on:click={panelMenuClickHandler}>
			<SpaceBetweenThreeItems>
				<div slot="left" class="red_petal"></div>
				<div slot="center" class="green_petal"></div>
				<div slot="right" class="blue_petal"></div>
			</SpaceBetweenThreeItems>
			<SpaceBetweenThreeItems>
				<div slot="left" class="green_petal"></div>
				<div slot="center" class="red_petal"></div>
				<div slot="right" class="blue_petal"></div>
			</SpaceBetweenThreeItems>
			<SpaceBetweenThreeItems>
				<div slot="left" class="red_petal"></div>
				<div slot="center" class="blue_petal"></div>
				<div slot="right" class="green_petal"></div>
			</SpaceBetweenThreeItems>
			{#if show_panel_dropdown}
				<div class="panel_dropdown"></div>
			{/if}
		</div>
		<div class="panel_title_text">
			{#if currentPanel == "inbox"}
				<span class="">Inbox</span>
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
		padding-left: 20px;
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
