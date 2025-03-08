<script>
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	export let dots = [];
	export let selected = dots[0];
	export let disabled = false;

	let current_dot = null;

	$: {
		if (selected) {
			if (dots.length > 0) {
				for (let i = 0; i < dots.length; i++) {
					if (selected.id == dots[i].id) {
						current_dot = dots[i];
						break;
					}
				}
			}
		}
	}

	function dotClickHandler(event, dot) {
		current_dot = dot;
		dispatch("change", current_dot);
	}
</script>

<div class="table" class:disabled={disabled == true}>
	<div class="cell">
		<ul class="dots">
			{#each dots as dot}
				<!-- svelte-ignore a11y_click_events_have_key_events -->
				<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
				<li
					class:active={current_dot.id == dot.id}
					on:click={(event) => {
						dotClickHandler(event, dot);
					}}
				>
					<!-- svelte-ignore a11y_invalid_attribute -->
					<div class="dot_item">{dot.title}</div>
				</li>
			{/each}
		</ul>
	</div>
</div>

<style>
	.table {
		width: 100%;
		height: 100%;
		display: table;
	}
	.table.disabled * {
		cursor: not-allowed;
		pointer-events: none;
	}
	.cell {
		display: table-cell;
		vertical-align: middle;
		margin: auto;
		text-align: center;
	}
	.dots {
		position: relative;
		display: inline-block;
		margin: 0;
		padding: 0;
		list-style: none;
		cursor: default;
	}
	.dots .active .dot_item {
		-webkit-transform: scale3d(1.3, 1.3, 1.3);
		transform: scale3d(1.3, 1.3, 1.3);
	}
	.dots .active .dot_item:after {
		height: 100%;
	}
	.dots li {
		position: relative;
		display: block;
		float: left;
		margin: 0 16px;
		width: 20px;
		height: 20px;
		cursor: pointer;
	}
	.dots li .dot_item {
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		outline: none;
		border-radius: 50%;
		text-indent: -999em;
		cursor: pointer;
		position: absolute;
		overflow: hidden;
		background-color: transparent;
		box-shadow: inset 0 0 0 2px #6a9113;
		-webkit-transition: all 0.3s ease;
		transition: all 0.3s ease;
		-webkit-transform: scale3d(1, 1, 1);
		transform: scale3d(1, 1, 1);
	}
	.dots li .dot_item:after {
		content: "";
		position: absolute;
		bottom: 0;
		height: 0;
		left: 0;
		width: 100%;
		background: #6a9113; /* fallback for old browsers */
		background: -webkit-linear-gradient(
			to right,
			#141517,
			#6a9113
		); /* Chrome 10-25, Safari 5.1-6 */
		background: linear-gradient(
			to right,
			#141517,
			#6a9113
		); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */

		box-shadow: 0 0 1px #fff;
		-webkit-transition: height 0.3s ease;
		transition: height 0.3s ease;
	}

	/* Font Face: Lato */
	@font-face {
		font-family: "Lato";
		font-style: normal;
		font-weight: 100;
		src:
			local("Lato Hairline"),
			local("Lato-Hairline"),
			url(http://themes.googleusercontent.com/static/fonts/lato/v6/boeCNmOCCh-EWFLSfVffDg.woff)
				format("woff");
	}
	@font-face {
		font-family: "Lato";
		font-style: normal;
		font-weight: 300;
		src:
			local("Lato Light"),
			local("Lato-Light"),
			url(http://themes.googleusercontent.com/static/fonts/lato/v6/KT3KS9Aol4WfR6Vas8kNcg.woff)
				format("woff");
	}
	@font-face {
		font-family: "Lato";
		font-style: normal;
		font-weight: 400;
		src:
			local("Lato Regular"),
			local("Lato-Regular"),
			url(http://themes.googleusercontent.com/static/fonts/lato/v6/9k-RPmcnxYEPm8CNFsH2gg.woff)
				format("woff");
	}
	@font-face {
		font-family: "Lato";
		font-style: normal;
		font-weight: 700;
		src:
			local("Lato Bold"),
			local("Lato-Bold"),
			url(http://themes.googleusercontent.com/static/fonts/lato/v6/wkfQbvfT_02e2IWO3yYueQ.woff)
				format("woff");
	}
</style>
