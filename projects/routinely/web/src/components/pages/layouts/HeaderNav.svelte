<script>
	import InboxIcon from "components/svg/InboxIcon.svelte";
	import RoutinesIcon from "components/svg/RoutinesIcon.svelte";
	import { goto } from "$app/navigation";
	import { page } from "$app/state";
	let currentTabName = "inbox";

	let navlinks = ["/inbox", "/routines"];
	function goToInboxClickHandler(event) {
		currentTabName = "inbox";
		goto("/inbox");
	}
	function goToRoutinesClickHandler(event) {
		currentTabName = "routines";
		goto("/routines");
	}
</script>

<nav>
	<svg
		viewBox="0 0 2 3"
		aria-hidden="true"
		style="position: relative; left: 3px;"
	>
		<path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z" />
	</svg>
	<div class="nav_links">
		<div
			class:selected={page.url.pathname === "/inbox"}
			class="nav_link_item"
		>
			<InboxIcon size="30px"></InboxIcon>
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div class="nav_link_text" on:click={goToInboxClickHandler}>
				Inbox
			</div>
		</div>
		<div
			class:selected={page.url.pathname === "/routines"}
			class="nav_link_item"
		>
			<RoutinesIcon size="30px"></RoutinesIcon>
			<!-- svelte-ignore a11y_click_events_have_key_events -->
			<!-- svelte-ignore a11y_no_static_element_interactions -->
			<div class="nav_link_text" on:click={goToRoutinesClickHandler}>
				Routines
			</div>
		</div>
		{#if navlinks.includes(page.url.pathname)}
			<div
				class="glider"
				class:inbox={currentTabName == "inbox"}
				class:routines={currentTabName == "routines"}
			></div>
		{/if}
	</div>
	<svg
		viewBox="0 0 2 3"
		aria-hidden="true"
		style="position: relative; right: 3px;"
	>
		<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
	</svg>
</nav>

<style>
	.nav_link_text {
		padding-left: 10px;
	}
	nav {
		display: flex;
		justify-content: center;
		background: transparent;
	}
	svg {
		width: 35px;
		height: inherit;
		display: block;
	}
	path {
		fill: #795548;
	}
	.nav_links {
		position: relative;
		padding: 0;
		margin: 0;
		height: inherit;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: #795548;
		color: #fff;
	}

	.nav_link_item {
		padding-left: 30px;
		padding-right: 30px;
		padding-top: 10px;
		padding-bottom: 10px;
		cursor: pointer;
		font-size: 20px;
		font-weight: bold;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.glider {
		width: 50%;
		height: 2px;
		border-radius: 0;
		position: absolute;
		box-shadow: 0px 0px 8px 0px hsl(262deg 100% 70% / 70%);
		background: linear-gradient(
			113deg,
			hsl(260deg 100% 64%) 0%,
			hsl(190deg 100% 55%) 100%
		);
		transition: all 0.3s;
		top: 48px;
		z-index: 2;
	}
	.glider.inbox {
		left: 0px;
	}
	.glider.routines {
		left: 50%;
		background: linear-gradient(90deg, #51a14c 0%, #10c33e 100%);
		box-shadow: 0px 0px 8px 0px rgba(47, 187, 12, 0.62);
	}
</style>
