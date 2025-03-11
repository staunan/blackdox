<script>
	import { page } from "$app/state";
	import logo from "$lib/images/svelte-logo.svg";
	import github from "$lib/images/github.svg";
	import { goto } from "$app/navigation";
	import { logoutUser } from "apis/apis.js";

	async function logoutHandler() {
		let response = await logoutUser();
		if (!response.HasError) {
			goto("/login");
		} else {
			console.log("Error while logging out");
			console.log(response);
		}
	}
</script>

<header>
	<div class="corner">
		<a href="https://svelte.dev/docs/kit">
			<img src={logo} alt="SvelteKit" />
		</a>
	</div>

	<nav>
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z" />
		</svg>
		<div class="nav_links">
			<div
				aria-current={page.url.pathname === "/inbox"
					? "page"
					: undefined}
			>
				<a href="/inbox">Inbox</a>
			</div>
			<div
				aria-current={page.url.pathname === "/routines"
					? "page"
					: undefined}
			>
				<a href="/routines">Routines</a>
			</div>
			<div
				aria-current={page.url.pathname.startsWith("/progress")
					? "page"
					: undefined}
			>
				<a href="/progress">Progress</a>
			</div>
		</div>
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
		</svg>
	</nav>

	<div class="corner">
		<!-- svelte-ignore a11y_click_events_have_key_events -->
		<!-- svelte-ignore a11y_missing_attribute -->
		<!-- svelte-ignore a11y_no_static_element_interactions -->
		<a on:click={logoutHandler}>
			<img src={github} alt="GitHub" />
		</a>
	</div>
</header>

<style>
	header {
		display: flex;
		justify-content: space-between;
	}

	.corner {
		width: 3em;
		height: 3em;
	}

	.corner a {
		display: flex;
		align-items: center;
		justify-content: center;
		width: 100%;
		height: 100%;
	}

	.corner img {
		width: 2em;
		height: 2em;
		object-fit: contain;
	}

	nav {
		display: flex;
		justify-content: center;
		background: transparent;
	}
	svg {
		width: 2em;
		height: 3em;
		display: block;
	}

	path {
		fill: #795548;
	}
	.nav_links {
		position: relative;
		padding: 0;
		margin: 0;
		height: 3em;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: #795548;
		color: #fff;
	}

	.nav_links > div {
		position: relative;
		height: 100%;
	}

	.nav_links > div[aria-current="page"]::before {
		content: "";
		width: 100%;
		height: 100%;
		border-radius: 0;
		background-color: rgba(0, 0, 0, 0.3);
		position: absolute;
		left: 0;
		bottom: 0;
	}
	nav a {
		display: flex;
		height: 100%;
		align-items: center;
		padding: 0 1.5rem;
		font-weight: 700;
		text-transform: uppercase;
		letter-spacing: 0.1em;
		text-decoration: none;
		transition: color 0.2s linear;
		color: #fff;
		font-size: 20px;
	}
	nav a:hover {
		color: #ff5722;
	}
</style>
