<script>
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	let currentTabName = $state("daily");

	function dailyClickHandler(event) {
		currentTabName = "daily";
		dispatch("change", "daily");
	}
	function weeklyClickHandler(event) {
		currentTabName = "weekly";
		dispatch("change", "weekly");
	}
	function monthlyClickHandler(event) {
		currentTabName = "monthly";
		dispatch("change", "monthly");
	}
	function yearlyClickHandler(event) {
		currentTabName = "yearly";
		dispatch("change", "yearly");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<div class="tabs">
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" onclick={dailyClickHandler}>Daily</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" onclick={weeklyClickHandler}>Weekly</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" onclick={monthlyClickHandler}>Monthly</div>
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div class="tab" onclick={yearlyClickHandler}>Yearly</div>
	<!-- <div class="tabsShadow"></div> -->
	<div
		class="glider"
		class:daily={currentTabName == "daily"}
		class:monthly={currentTabName == "monthly"}
		class:weekly={currentTabName == "weekly"}
		class:yearly={currentTabName == "yearly"}
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
	.glider.daily {
		left: 0px;
	}
	.glider.weekly {
		left: 25%;
		background: linear-gradient(90deg, #51a14c 0%, #10c33e 100%);
		box-shadow: 0px 0px 8px 0px rgba(47, 187, 12, 0.62);
	}
	.glider.monthly {
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
	.glider.yearly {
		left: 75%;
		background: linear-gradient(90deg, #b9326f 0%, #ff5ddc 100%);
		box-shadow: 0px 0px 8px 0px rgba(231, 13, 93, 0.57);
	}
</style>
