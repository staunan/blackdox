<script>
	import Modal from "components/modals/Modal.svelte";
	import TextBox from "components/form/TextBox.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { createEventDispatcher } from "svelte";
	import { checkIfUsernameAvailable } from "apis/apis.js";
	import { store, user_details } from "store";
	import { validateUsername } from "lib/js/validation.js";

	export let active = false;
	export let overlayclose = true;

	let user = null;
	user_details.subscribe((v) => {
		if (v) {
			user = v;
		}
	});

	let newUsername = "";
	let hasError = false;
	let errorMessage = "";

	let typingTimer = null;
	const typingDelay = 1000;

	let isUsernameAvailable = false;
	let available_message = "";

	const dispatch = createEventDispatcher();

	function closeModal() {
		dispatch("close");
	}

	async function submitHandler() {
		if (!newUsername) {
			hasError = true;
			errorMessage = "Email is required.";
			return;
		}
		let formData = {
			username: newUsername,
		};
		// let response = await changeUsername(formData);

		// if (response.HasError == false) {
		// 	await store.getUser();
		// 	dispatch("usernameChanged");
		// 	newUsername = "";
		// }
	}
	function newUsernameChangedHandler(event) {
		isUsernameAvailable = false;
		available_message = "";
		clearTimeout(typingTimer); // Clear the previous timer
		typingTimer = setTimeout(() => {
			checkIfUsernameIsAvailable(event.detail);
		}, typingDelay);
	}
	async function checkIfUsernameIsAvailable(username_text) {
		if (username_text) {
			if (!validateUsername(username_text)) {
				return;
			}
			if (username_text === user.Username) {
				return;
			}
			try {
				let formData = {
					username: username_text,
				};
				let response = await checkIfUsernameAvailable(formData);
				if (response.HasError === true) {
					console.log(response.Message);
				} else {
					if (response.Data === true) {
						isUsernameAvailable = true;
						available_message = "Username is available";
						newUsername = username_text;
					} else {
						isUsernameAvailable = false;
						available_message = "Username is taken";
					}
				}
			} catch (error) {
				console.log(error);
			}
		}
	}
</script>

<Modal {active} {overlayclose} on:close={closeModal}>
	<div class="modal_body">
		<h2>Change Username</h2>
		<TextBox
			label="New Username"
			placeholder="Enter your new username"
			value={newUsername}
			on:change={newUsernameChangedHandler}
			{hasError}
			{errorMessage}
		/>
		{#if available_message.length > 0}
			{#if isUsernameAvailable === true}
				<div class="available">{available_message}</div>
			{:else}
				<div class="not_available">{available_message}</div>
			{/if}
		{/if}
		<div></div>
		<div class="button_group">
			<SubmitButton
				title="Update"
				disabled={!isUsernameAvailable}
				on:tap={submitHandler}
			/>
			<SubmitButton title="Cancel" on:tap={closeModal} color="grey" />
		</div>
	</div>
</Modal>

<style>
	h2 {
		text-align: center;
	}
	.modal_body {
		padding: 20px;
	}
	.button_group {
		display: flex;
		justify-content: space-between;
		margin-top: 20px;
	}
	.available {
		font-size: 14px;
		color: green;
		padding-top: 10px;
	}
	.not_available {
		padding-top: 10px;
		font-size: 14px;
		color: red;
	}
</style>
