<script>
	import { getUserDefaultImage } from "apis/apis.js";
	import { createEventDispatcher } from "svelte";
	import "animate.css";

	const dispatch = createEventDispatcher();
	let fileinput = $state();
	let avatar = $state();

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

<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_noninteractive_element_interactions -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div
	class="picture"
	onclick={() => {
		fileinput.click();
	}}
>
	{#if avatar}
		<img class="avatar" src={avatar} alt="d" />
	{:else}
		<img class="avatar" src={getUserDefaultImage()} alt="" />
	{/if}

	<div class="hidden_input">
		<input
			style="display:none"
			type="file"
			accept=".jpg, .jpeg, .png"
			onchange={(e) => onFileSelectedHandler(e)}
			bind:this={fileinput}
		/>
	</div>
</div>

<style>
	.picture {
		display: flex;
		justify-content: center;
		align-items: center;
		border-radius: 50%;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
		cursor: pointer;
	}
	.picture img.avatar {
		object-fit: cover;
		width: 300px;
		height: 300px;
		border-radius: 50%;
		box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
	}
</style>
