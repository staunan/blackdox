<script>
	import Card from "components/Card.svelte";
	import ProfilePicture from "components/form/ProfilePicture.svelte";
</script>

<Card>
	<div class="card_body">
		<h1 class="center mb10 form_heading">Upload Display Photo</h1>
		<ProfilePicture></ProfilePicture>
	</div>
</Card>

<style>
	.card_body {
		min-height: 500px;
		width: 100%;
		padding-left: 200px;
		padding-right: 200px;
		padding-top: 50px;
		padding-bottom: 50px;
	}
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.mb10 {
		margin-bottom: 10px;
	}
	.form_heading {
		font-family: monospace;
		font-size: 24px;
		font-weight: bold;
		margin-top: 10px;
		padding-bottom: 40px;
	}
</style>
