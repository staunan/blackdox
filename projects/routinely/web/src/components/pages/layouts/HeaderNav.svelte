<script>
	import { page } from "$app/state";
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
		<div aria-current={page.url.pathname === "/inbox" ? "page" : undefined}>
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
	<svg
		viewBox="0 0 2 3"
		aria-hidden="true"
		style="position: relative; right: 3px;"
	>
		<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
	</svg>
</nav>

<style>
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
