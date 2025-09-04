<script>
	import PanelIcon from "components/pages/routines/PanelIcon.svelte";
	import PanelRoutines from "./routines/PanelRoutines.svelte";
	import PanelTrash from "./trash/PanelTrash.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";

	import TrashIcon from "components/svg/TrashIcon.svelte";
	import PlusIcon from "components/svg/PlusIcon.svelte";
	import RoutinesIcon from "components/svg/RoutinesIcon.svelte";

	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import SearchBar from "components/form/SearchBar.svelte";

	let currentPanel = $state("routines");
	let show_panel_dropdown = $state(false);
	let total_items_in_trash = 0;
	let search_text = $state("");

	onMount(() => {
		search_text = localStorage.getItem("search");
		const funcRef = (event) => {
			if (event.target.closest(".routine_panel_menu")) {
				show_panel_dropdown = true;
			} else {
				show_panel_dropdown = false;
			}
		};
		window.addEventListener("click", funcRef);

		return () => {
			window.removeEventListener("click", funcRef);
		};
	});

	function panelMenuClickHandler(event) {
		show_panel_dropdown = true;
	}
	function goToCreateRoutinesHandler(event) {
		goto("create-routine");
	}
	function searchTextChangeHandler(event) {
		search_text = event.detail;
		localStorage.setItem("search", search_text);
	}
</script>

<div class="routines_page">
	<div class="panel_title">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<div class="routine_panel_menu" onclick={panelMenuClickHandler}>
			<PanelIcon></PanelIcon>
			{#if show_panel_dropdown}
				<div class="routine_panel_dropdown">
					<div
						class="panel_dropdown_list_item routine"
						onclick={() => {
							currentPanel = "routines";
							show_panel_dropdown = false;
						}}
					>
						<div
							style="width: 40px; height: 40px; display: flex; justify-content: center; align-items: center;"
						>
							<RoutinesIcon size="30px"></RoutinesIcon>
						</div>
						<div class="panel_dropdown_list_item_text">
							Routines
						</div>
					</div>
					<div
						class="panel_dropdown_list_item trash"
						onclick={() => {
							currentPanel = "trash";
							show_panel_dropdown = false;
						}}
					>
						<TrashIcon></TrashIcon>
						<div class="panel_dropdown_list_item_text">Trash</div>
					</div>
				</div>
			{/if}
		</div>

		{#if currentPanel == "routines"}
			<div class="panel_title_text">
				<div class="page_title">Routines</div>
				<div class="search_routine_container">
					<SearchBar
						value={search_text}
						placeholder="Type to Search Routine..."
						on:change={searchTextChangeHandler}
					></SearchBar>
				</div>
				<div class="right_panel">
					<SubmitButton
						title="Create Routine"
						on:tap={goToCreateRoutinesHandler}
						color="blue"
					>
						<PlusIcon size="20px"></PlusIcon>
					</SubmitButton>
				</div>
			</div>
		{:else if currentPanel == "trash"}
			<div class="panel_title_text">
				<span>Trash</span>
			</div>
		{/if}
	</div>
	<div class="panel_body">
		{#if currentPanel == "routines"}
			<PanelRoutines search={search_text}></PanelRoutines>
		{:else if currentPanel == "trash"}
			<PanelTrash></PanelTrash>
		{/if}
	</div>
</div>

<style>
	.routines_page {
		padding-top: 30px;
	}
	.routine_panel_dropdown {
		position: absolute;
		top: 110%;
		left: 0;
		background-color: #fff;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
		min-height: 50px;
		z-index: 10000001;
		width: 200px;
	}
	.panel_title {
		padding-bottom: 20px;
		display: flex;
	}
	.panel_title_text {
		font-size: 37px;
		font-weight: bold;
		padding-left: 20px;
		display: flex;
		width: 100%;
		height: 50px;
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

	.panel_dropdown_list_item {
		padding: 15px;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		font-size: 20px;
		font-weight: bold;
		cursor: pointer;
		border-bottom: 1px solid #ccc;
		background-color: #ddd;
	}
	.panel_dropdown_list_item:last-child {
		border-bottom: none;
	}
	.panel_dropdown_list_item_text {
		padding-left: 10px;
	}
	.panel_dropdown_list_item:hover {
		background-color: #ff8686;
	}
	.right_panel {
		flex: 1;
		display: flex;
		justify-content: flex-end;
	}
	.search_routine_container {
		padding-left: 30px;
		display: flex;
		align-items: center;
	}
</style>
