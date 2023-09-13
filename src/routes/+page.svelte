<svelte:head>
	<title>Home</title>
	<meta name="description" content="Labani's Home" />
</svelte:head>

<section>
	<h1 class="page_title">Hi, I am Labani ♡</h1>
	<div class="gallery">
		<div class="image_container">
			<img src={at_manasa_mondir} alt="At Manasa Mondir">
		</div>
		<div class="image_container">
			<img src={beauty_queen} alt="Beauty Queen">
		</div>
	</div>
</section>
<script>
import at_manasa_mondir from '$lib/images/labani/at_manasa_mondir.jpeg';
import beauty_queen from '$lib/images/labani/beauty_queen.jpeg';
</script>
<style>
section {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	flex: 0.6;
}
.page_title{
	margin-top: 20px;
	margin-bottom: 20px;
}
.gallery{
	display: flex;
	flex-wrap: wrap;
	justify-content: center;
}
.gallery .image_container{
	flex-basis: 300px;
}
.gallery .image_container>img{
	width: 100%;
	object-fit: cover;
}
h1 {
	width: 100%;
}
</style>
