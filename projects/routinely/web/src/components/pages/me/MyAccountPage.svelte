<script>
	import { user_details } from "store";
	import { config } from "config/api_url.js";
	import { getUserDefaultImage } from "apis/apis.js";

	let user = null;
	let avatar;
	user_details.subscribe((v) => {
		if (v) {
			user = v;
		}
	});
</script>

<div class="display_profile_picture_container">
	<div class="picture">
		{#if user && user.DisplayPictureName}
			<img
				class="avatar"
				src={config.user_display_picture_url +
					"user_profile_pictures/" +
					user.DisplayPictureName}
				alt="d"
			/>
		{:else}
			<img class="avatar" src={getUserDefaultImage()} alt="" />
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
</div>

<style>
	.display_profile_picture_container {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding-top: 30px;
	}
	.picture img.avatar {
		object-fit: cover;
		width: 200px;
		height: 200px;
		border-radius: 50%;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
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
