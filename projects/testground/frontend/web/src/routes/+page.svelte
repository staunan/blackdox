<svelte:head>
	<title>Home</title>
	<meta name="description" content="Labani's Home" />
</svelte:head>

<section>
	<MenuBar></MenuBar>
	<div class="gallery">

	</div>
</section>
<script>
import MenuBar from 'components/pages/MenuBar.svelte';
</script>
<style>
section {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	flex: 0.6;
}
.gallery{
	display: flex;
	flex-wrap: wrap;
	justify-content: center;
}
</style>
