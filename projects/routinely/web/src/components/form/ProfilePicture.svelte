<script>
	import { getUserDefaultImage } from "apis/apis.js";
	import { createEventDispatcher } from "svelte";

	const dispatch = createEventDispatcher();
	let fileinput;
	let avatar;

	async function onFileSelectedHandler(e) {
		let image = e.target.files[0];
		let reader = new FileReader();
		reader.readAsDataURL(image);
		reader.onload = async function (e) {
			avatar = e.target.result;
			dispatch("change", image);
		};
	}
</script>

<div class="picture">
	{#if avatar}
		<img class="avatar" src={avatar} alt="d" />
	{:else}
		<img class="avatar" src={getUserDefaultImage()} alt="" />
	{/if}
	<!-- svelte-ignore a11y_click_events_have_key_events -->
	<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
	<!-- svelte-ignore a11y_no_static_element_interactions -->
	<div
		class="upload_picture_container"
		on:click={() => {
			fileinput.click();
		}}
	>
		<input
			style="display:none"
			type="file"
			accept=".jpg, .jpeg, .png"
			on:change={(e) => onFileSelectedHandler(e)}
			bind:this={fileinput}
		/>
		<div class="upload_text">Change Display Photo</div>
	</div>
</div>

<style>
	.picture {
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 50%;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		position: relative;
	}
	.picture img.avatar {
		object-fit: cover;
		width: 300px;
		height: 300px;
		border-radius: 50%;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}
	.upload_picture_container {
		position: absolute;
		top: 0;
		left: 0;
		width: 300px;
		height: 300px;
		border-radius: 50%;
		display: flex;
		justify-content: center;
		align-items: center;
		background-color: rgba(0, 0, 0, 0.7);
		cursor: pointer;
		display: none;
		z-index: 10;
	}
	.picture:hover .upload_picture_container {
		display: flex !important;
	}
	.upload_text {
		color: #fff;
		font-weight: bold;
		font-size: 16px;
		font-family: monospace;
	}
</style>
