<script>
	import CameraIcon from "components/svg/CameraIcon.svelte";
	import { createEventDispatcher } from "svelte";
	let { avatar = $bindable() } = $props();

	let fileinput = $state();
	const dispatch = createEventDispatcher();

	function handleProfilePictureClick() {
		fileinput.click();
	}

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
<!-- svelte-ignore a11y_no_static_element_interactions -->
<div class="profile-picture-container">
	<!-- svelte-ignore a11y_missing_attribute -->
	<img src={avatar} class="profile-picture" />
	<div
		class="change-profile-picture-icon"
		onclick={handleProfilePictureClick}
	>
		<CameraIcon size="30px"></CameraIcon>
	</div>
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
	.profile-picture-container {
		position: relative;
		width: 200px;
		height: 200px;
		border-radius: 50%;
		cursor: pointer;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
	}

	.profile-picture-container:hover {
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
	}

	.profile-picture {
		width: 100%;
		height: 100%;
		object-fit: cover;
		border-radius: 50%;
	}

	.change-profile-picture-icon {
		position: absolute;
		left: 85%;
		bottom: 30px;
		font-size: 20px;
		color: #673ab7;
		cursor: pointer;
		display: flex;
		justify-content: center;
		align-items: center;
		width: 40px;
		height: 40px;
		border-radius: 50%;
		transform: translate(-50%, 50%);
		background-color: #673ab7;
		box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
		z-index: 10;
		opacity: 0.2;
		transition: all 0.2s ease-in-out;
	}

	.change-profile-picture-icon:hover {
		opacity: 1;
	}
</style>
