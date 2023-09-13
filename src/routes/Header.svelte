<script>
	import { page } from '$app/stores';
	import logo from '$lib/images/svelte-logo.svg';
	import github from '$lib/images/github.svg';
</script>

<header>
	<div class="website_title">Blackdox</div>
</header>

<style>
.website_title{
	font-size: 25px;
	font-weight: bold;
	color: #fff;
}
header {
	display: flex;
	justify-content: center;
	background: #C33764;  /* fallback for old browsers */
	background: -webkit-linear-gradient(to right, #1D2671, #C33764);  /* Chrome 10-25, Safari 5.1-6 */
	background: linear-gradient(to right, #1D2671, #C33764); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
	align-items: center;
	height: 60px;
}
</style>
