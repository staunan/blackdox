<script>
	import { user_details } from "store";
	import { config } from "config/api_url.js";
	import ProfilePictureUpdatedSuccess from "components/pages/me/ProfilePictureUpdatedSuccess.svelte";
	import { store } from "store";
	import {
		getUserDefaultImage,
		uploadDisplayPhotoInMyAccount,
	} from "apis/apis.js";
	import ProfilePicture from "components/pages/me/ProfilePicture.svelte";

	let user = null;
	let avatar;
	let isDisplayPhotoUploadedSuccessModalActive = false;
	user_details.subscribe((v) => {
		if (v) {
			user = v;
		}
	});

	async function profilePictureChangedHandler(event) {
		let image = event.detail;
		if (!image) {
			console.log("Missing Image");
			return;
		}
		let data = { file: image };
		let response = await uploadDisplayPhotoInMyAccount(data);
		if (response.HasError == false) {
			isDisplayPhotoUploadedSuccessModalActive = true;
			await store.getUser();
		} else {
			console.log("Some error has occured!");
		}
	}
	function closeDisplayPhotoUploadedSuccessModalHandler() {
		isDisplayPhotoUploadedSuccessModalActive = false;
	}
</script>

<div class="display_profile_picture_container">
	<div class="picture">
		{#if user && user.DisplayPictureName}
			<ProfilePicture
				avatar={config.user_display_picture_url +
					"user_profile_pictures/" +
					user.DisplayPictureName}
				on:change={profilePictureChangedHandler}
			></ProfilePicture>
		{:else}
			<ProfilePicture avatar={getUserDefaultImage()}></ProfilePicture>
		{/if}
	</div>
	{#if user && user.FullName}
		<div class="name">
			{user.FullName}
		</div>
	{/if}
	{#if user && user.Email}
		<div class="email">
			{user.Email}
		</div>
	{/if}

	<ProfilePictureUpdatedSuccess
		active={isDisplayPhotoUploadedSuccessModalActive}
		on:close={closeDisplayPhotoUploadedSuccessModalHandler}
	></ProfilePictureUpdatedSuccess>
</div>

<style>
	.display_profile_picture_container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding-top: 30px;
	}
	.name {
		font-size: 30px;
		font-weight: bold;
		padding-top: 20px;
	}
	.email {
		font-size: 15px;
		color: #3f51b5;
		font-weight: bold;
	}
</style>
