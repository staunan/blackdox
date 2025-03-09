<script>
	import Card from "components/Card.svelte";
	import ProfilePicture from "components/form/ProfilePicture.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import LinkButton from "components/buttons/LinkButton.svelte";
	import { uploadProfilePhoto } from "apis/apis.js";

	let selectedProfilePhoto = null;
	function imageChangeHandler(e) {
		selectedProfilePhoto = e.detail;
	}
	function skipStepHandler(e) {
		alert("will implement");
	}
	async function uploadPhotoHandler() {
		let data = { file: selectedProfilePhoto };
		let response = await uploadProfilePhoto(data);
	}
</script>

<Card>
	<div class="card_body">
		<h1 class="center mb10 form_heading">Display Photo</h1>
		<ProfilePicture on:change={imageChangeHandler}></ProfilePicture>
		<div class="upload_photo_button center">
			<SubmitButton
				title="Upload Display Photo"
				on:tap={uploadPhotoHandler}
			></SubmitButton>
		</div>
		<div class="skip_step_container">
			<LinkButton label="Skip this step" on:tap={skipStepHandler}
			></LinkButton>
		</div>
	</div>
</Card>

<style>
	.skip_step_container {
		width: 100%;
		display: flex;
		justify-content: flex-end;
		align-items: flex-end;
	}
	.card_body {
		min-height: 500px;
		width: 100%;
		padding-left: 20px;
		padding-right: 20px;
		padding-top: 50px;
		padding-bottom: 20px;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}
	.upload_photo_button {
		padding: 30px;
		padding-top: 60px;
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
