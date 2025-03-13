<script>
	import Card from "components/Card.svelte";
	import ProfilePicture from "components/form/ProfilePicture.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import LinkButton from "components/buttons/LinkButton.svelte";
	import { createEventDispatcher } from "svelte";
	import {
		uploadDisplayPhotoInRegistrationStep,
		skipUploadDisplayPhotoInRegistrationStep,
	} from "apis/apis.js";
	const dispatch = createEventDispatcher();
	let photo_upload_button_text = "Upload Display Photo";
	let disable_upload_button = false;
	let skip_button_text = "Skip this step";
	let disable_skip_button = false;

	let selectedProfilePhoto = null;
	function imageChangeHandler(e) {
		selectedProfilePhoto = e.detail;
	}
	async function skipStepHandler(e) {
		let data = {};
		skip_button_text = "Skipping...";
		disable_skip_button = true;
		setTimeout(async () => {
			let response = await skipUploadDisplayPhotoInRegistrationStep(data);
			if (response.HasError == false) {
				dispatch("skipped", response.Data);
			} else {
				console.log("Some error has occured!");
			}
		}, 1000);
	}
	function uploadPhotoHandler() {
		if (!selectedProfilePhoto) {
			console.log("Missing Image");
			return;
		}
		let data = { file: selectedProfilePhoto };
		photo_upload_button_text = "Uploading...";
		disable_upload_button = true;
		setTimeout(async () => {
			let response = await uploadDisplayPhotoInRegistrationStep(data);
			if (response.HasError == false) {
				dispatch("uploaded", response.Data);
			} else {
				console.log("Some error has occured!");
			}
		}, 1000);
	}
</script>

<Card>
	<div class="card_body">
		<h1 class="center mb10 form_heading">Display Photo</h1>
		<ProfilePicture on:change={imageChangeHandler}></ProfilePicture>
		<div class="upload_photo_button center">
			<SubmitButton
				disabled={disable_upload_button}
				title={photo_upload_button_text}
				on:tap={uploadPhotoHandler}
			></SubmitButton>
		</div>
		<div class="skip_step_container">
			<LinkButton
				disabled={disable_skip_button}
				label={skip_button_text}
				on:tap={skipStepHandler}
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
