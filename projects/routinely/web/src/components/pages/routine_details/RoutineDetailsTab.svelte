<script>
	import AboutIcon from "components/svg/AboutIcon.svelte";
	import HistoryIcon from "components/svg/HistoryIcon.svelte";
	import ProgressIcon from "components/svg/ProgressIcon.svelte";
	import SettingsIcon from "components/svg/SettingsIcon.svelte";

	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	let currentTabName = "about";

	function aboutClickHandler(event) {
		currentTabName = "about";
		dispatch("change", "about");
	}
	function progressClickHandler(event) {
		currentTabName = "progress";
		dispatch("change", "progress");
	}
	function historyClickHandler(event) {
		currentTabName = "history";
		dispatch("change", "history");
	}
	function settingsClickHandler(event) {
		currentTabName = "settings";
		dispatch("change", "settings");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<div class="tabs">
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" on:click={aboutClickHandler}>
		<AboutIcon></AboutIcon>
		<div class="tab_title">About</div>
	</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" on:click={progressClickHandler}>
		<ProgressIcon></ProgressIcon>
		<div class="tab_title">Progress</div>
	</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" on:click={historyClickHandler}>
		<HistoryIcon size="35px"></HistoryIcon>
		<div class="tab_title">History</div>
	</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" on:click={settingsClickHandler}>
		<SettingsIcon></SettingsIcon>
		<div class="tab_title">Settings</div>
	</div>
	<!-- <div class="tabsShadow"></div> -->
	<div
		class="glider"
		class:about={currentTabName == "about"}
		class:progress={currentTabName == "progress"}
		class:history={currentTabName == "history"}
		class:settings={currentTabName == "settings"}
	></div>
</div>

<style>
	.tabs {
		display: flex;
		justify-content: flex-start;
		position: relative;
	}
	.tab {
		flex-basis: 25%;
		padding-left: 20px;
		display: flex;
		justify-content: flex-start;
		align-items: center;
		background-color: hsl(232deg 6% 17%);
		color: #fff;
		height: 60px;
		cursor: pointer;
		font-size: 18px;
		font-weight: bold;
	}
	.glider {
		width: 25%;
		padding: 0px 15px;
		height: 5px;
		border-radius: 0;
		position: absolute;
		box-shadow: 0px 0px 8px 0px hsl(262deg 100% 70% / 70%);
		background: linear-gradient(
			113deg,
			hsl(260deg 100% 64%) 0%,
			hsl(190deg 100% 55%) 100%
		);
		transition: all 0.3s;
		top: 55px;
		z-index: 2;
	}
	.glider.about {
		left: 0px;
	}
	.glider.progress {
		left: 25%;
		background: linear-gradient(90deg, #51a14c 0%, #10c33e 100%);
		box-shadow: 0px 0px 8px 0px rgba(47, 187, 12, 0.62);
	}
	.glider.history {
		left: 50%;
		background: linear-gradient(
			90deg,
			#faffcc 0%,
			#f5eea3 10%,
			#ffe48a 40%,
			#ffb54d 65%,
			#ff974d 85%,
			#ff8052 100%
		);
		box-shadow: 0px 0px 8px 0px hsl(17.72deg 100% 70% / 70%);
	}
	.glider.settings {
		left: 75%;
		background: linear-gradient(90deg, #b9326f 0%, #ff5ddc 100%);
		box-shadow: 0px 0px 8px 0px rgba(231, 13, 93, 0.57);
	}
	.tab_title {
		padding-left: 20px;
	}
</style>
