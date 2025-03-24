<script>
	import { config } from "config/api_url.js";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	export let user;

	function tabHandler(event) {
		dispatch("click");
	}
</script>

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="header_profile_image" on:click={tabHandler}>
	{#if user}
		<!-- svelte-ignore a11y_missing_attribute -->
		<img
			src={config.user_display_picture_url +
				"user_profile_pictures/" +
				user.DisplayPictureName}
		/>
	{:else}
		<!-- svelte-ignore a11y_img_redundant_alt -->
		<!-- svelte-ignore a11y_missing_attribute -->
		<img src={config.user_display_picture_url + "user_default_image.jpg"} />
	{/if}
</div>

<style>
	.header_profile_image {
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 50%;
		border: 1px solid #ccc;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
	}
	img {
		border-radius: 50%;
		cursor: pointer;
		height: 45px;
		width: 45px;
		box-shadow: rgba(0, 0, 0, 0.1) 0px 4px 12px;
	}
</style>
