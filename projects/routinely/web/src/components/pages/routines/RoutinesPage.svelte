<script>
	import PanelRoutines from "./PanelRoutines.svelte";
	import PanelTrash from "./trash/PanelTrash.svelte";
	import SpaceBetweenThreeItems from "components/layouts/SpaceBetweenThreeItems.svelte";
	import { onMount } from "svelte";
	import { routines } from "store";

	let currentPanel = "routines";
	let show_panel_dropdown = false;
	let total_items_in_trash = 0;

	routines.subscribe((v) => {
		let deleted_routines = v.filter((item) => item.IsTrash === 1);
		total_items_in_trash = deleted_routines.length;
	});

	onMount(() => {
		document.addEventListener("click", function (event) {
			console.log("Event Listening");
			if (event.target.closest(".routine_panel_menu")) {
				show_panel_dropdown = true;
			} else {
				show_panel_dropdown = false;
			}
		});
	});

	function panelMenuClickHandler(event) {
		show_panel_dropdown = true;
	}
</script>

<div class="routines_page">
	<div class="panel_title">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="routine_panel_menu" on:click={panelMenuClickHandler}>
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
				<div class="routine_panel_dropdown">
					<div
						class="panel_dropdown_list_item"
						on:click={() => {
							currentPanel = "routines";
							show_panel_dropdown = false;
						}}
					>
						Routines
					</div>
					<div
						class="panel_dropdown_list_item"
						on:click={() => {
							currentPanel = "trash";
							show_panel_dropdown = false;
						}}
					>
						Trash ({total_items_in_trash})
					</div>
				</div>
			{/if}
		</div>
		<div class="panel_title_text">
			{#if currentPanel == "routines"}
				<span>Routines</span>
			{:else if currentPanel == "trash"}
				<span>Trash ({total_items_in_trash})</span>
			{/if}
		</div>
	</div>
	<div class="panel_body">
		{#if currentPanel == "routines"}
			<PanelRoutines></PanelRoutines>
		{:else if currentPanel == "trash"}
			<PanelTrash></PanelTrash>
		{/if}
	</div>
</div>

<style>
	.routine_panel_dropdown {
		position: absolute;
		top: 110%;
		left: 0;
		background-color: #fff;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		min-height: 50px;
		z-index: 10000001;
		width: 150px;
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
	.routine_panel_menu {
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
	.panel_dropdown_list_item {
		padding: 15px;
		display: flex;
		justify-content: flex-start;
		font-size: 20px;
		font-weight: bold;
		cursor: pointer;
	}
	.panel_dropdown_list_item:hover {
		background-color: #ddd;
	}
</style>
