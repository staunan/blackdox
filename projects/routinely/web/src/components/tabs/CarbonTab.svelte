<script>
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	function dailyClickHandler(event) {
		dispatch("change", "daily");
	}
	function weeklyClickHandler(event) {
		dispatch("change", "weekly");
	}
	function monthlyClickHandler(event) {
		dispatch("change", "monthly");
	}
	function yearlyClickHandler(event) {
		dispatch("change", "yearly");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<div class="tabs">
	<input
		id="tab-1"
		type="radio"
		class="tab tab-selector"
		checked="checked"
		name="tab"
	/>
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<label for="tab-1" class="tab tab-primary" on:click={dailyClickHandler}
		>Daily</label
	>
	<input id="tab-2" type="radio" class="tab tab-selector" name="tab" />
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<label for="tab-2" class="tab tab-success" on:click={weeklyClickHandler}
		>Weekly</label
	>
	<input id="tab-3" type="radio" class="tab tab-selector" name="tab" />
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<label for="tab-3" class="tab tab-default" on:click={monthlyClickHandler}
		>Monthly</label
	>
	<input id="tab-4" type="radio" class="tab tab-selector" name="tab" />
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<label for="tab-4" class="tab tab-warning" on:click={yearlyClickHandler}
		>Yearly</label
	>
	<div class="tabsShadow"></div>
	<div class="glider"></div>
</div>

<style>
	/* Basic */
	@import url("https://use.typekit.net/ovt6ynt.css");
	@import url("https://fonts.googleapis.com/css2?family=EB+Garamond:ital,wght@0,400;0,500;0,600;0,700;0,800;1,400;1,500;1,600;1,700;1,800&family=Oswald:wght@200;300;400;500;600;700&family=Roboto+Condensed:wght@300;400;700&family=Roboto:wght@100;300;400;500;700;900&family=Source+Sans+Pro:ital,wght@0,200;0,300;0,400;0,600;0,700;0,900;1,200;1,300;1,400;1,600;1,700;1,900&family=Spectral:wght@200;300;400;500;600;700;800&display=swap");
	.tabs {
		text-align: left;
		position: relative;
	}
	.tabs input {
		position: absolute;
		opacity: 0;
		margin: 0;
		padding: 0;
		visibility: hidden;
	}

	.tabs input + label:hover {
		color: #777;
	}
	.tabs .tabsShadow {
		width: 100%;
		padding: 10px;
		box-shadow: 0 0 25px 0 rgba(0, 0, 0, 0.04);
		position: absolute;
		height: 35px;
		z-index: 0;
		top: 0;
		pointer-events: none;
	}

	#tab-1:checked ~ .glider {
		left: 0px;
	}
	#tab-2:checked + label {
		color: #51a14c;
	}
	#tab-2:checked ~ .glider {
		background: linear-gradient(90deg, #51a14c 0%, #10c33e 100%);
		box-shadow: 0px 0px 8px 0px rgba(47, 187, 12, 0.62);
	}

	#tab-3:checked + label {
		color: #525252;
	}
	.tabs input + label {
		margin: 12px -2px;
		width: 25%;
		transition: all 0.3s;
		text-align: center;
		border: none;
		border-bottom: 0.5pt solid #1d1e22;
		color: #e7e9f5;
		display: inline-flex;
		-webkit-box-pack: center;
		padding: 12pt 15pt;
		vertical-align: top;
		justify-content: flex-start;
		background: #333;
		box-shadow: 1px 1px 2px 1px #231f20;
		text-transform: uppercase;
		font-weight: 300;
		font-size: 8pt;
		background: #26292c;
		margin-bottom: 0.5px;
		border-radius: 0;
		border-top: 1px solid #3b3d3f;
		letter-spacing: 0.1666666667em !important;
		font-family: "antique-olive" !important;
		background: hsl(232deg 6% 17%);
		cursor: pointer;
	}
	label.tab.tab-primary {
		border-radius: 0;
	}
	label.tab.tab-warning {
		border-radius: 0;
	}
	#tab-1:checked + label {
		color: deepskyblue;
	}
	#tab-3:checked + label {
		color: #ffadfa;
		color: #ffaa75;
	}
	#tab-4:checked + label {
		color: #ffadfa;
	}

	#tab-2:checked ~ .glider {
		left: 25%;
	}
	#tab-3:checked ~ .glider {
		background: linear-gradient(
			90deg,
			#faffcc 0%,
			#f5eea3 10%,
			#ffe48a 40%,
			#ffb54d 65%,
			#ff974d 85%,
			#ff8052 100%
		);
		left: 50%;
		box-shadow: 0px 0px 8px 0px hsl(17.72deg 100% 70% / 70%);
	}

	#tab-4:checked ~ .glider {
		background: linear-gradient(90deg, #b9326f 0%, #ff5ddc 100%);
		box-shadow: 0px 0px 8px 0px rgba(231, 13, 93, 0.57);
		left: 75% !important;
	}

	.glider {
		width: 25%;
		padding: 0px 15px;
		height: 5px;
		border-radius: 0 0 1px 1px;
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
</style>
