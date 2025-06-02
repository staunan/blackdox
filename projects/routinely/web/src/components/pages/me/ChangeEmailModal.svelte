<script>
	import Modal from "components/modals/Modal.svelte";
	import TextBox from "components/form/TextBox.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import { createEventDispatcher } from "svelte";
	import { changeUserEmail, checkIfEmailAvailable } from "apis/apis.js";
	import { store, user_details } from "store";
	import { validateEmail } from "lib/js/validation.js";

	export let active = false;
	export let overlayclose = true;

	let user = null;
	user_details.subscribe((v) => {
		if (v) {
			user = v;
		}
	});

	let newEmail = "";
	let hasError = false;
	let errorMessage = "";

	let typingTimer = null;
	const typingDelay = 1000;

	let isEmailAvailable = false;
	let available_message = "";

	const dispatch = createEventDispatcher();

	function closeModal() {
		dispatch("close");
	}

	async function submitHandler() {
		if (!newEmail) {
			hasError = true;
			errorMessage = "Email is required.";
			return;
		}
		let formData = {
			email: newEmail,
		};
		let response = await changeUserEmail(formData);

		if (response.HasError == false) {
			await store.getUser();
			dispatch("emailchanged");
			newEmail = "";
		}
	}
	function emailChangedHandler(event) {
		isEmailAvailable = false;
		available_message = "";
		clearTimeout(typingTimer); // Clear the previous timer
		typingTimer = setTimeout(() => {
			checkIfEmailIsAvailable(event.detail);
		}, typingDelay);
	}
	async function checkIfEmailIsAvailable(email_text) {
		if (email_text) {
			if (!validateEmail(email_text)) {
				return;
			}
			if (email_text === user.Email) {
				return;
			}
			try {
				let formData = {
					email: email_text,
				};
				let response = await checkIfEmailAvailable(formData);
				if (response.HasError === true) {
					console.log(response.Message);
				} else {
					if (response.Data === true) {
						isEmailAvailable = true;
						available_message = "Email is available";
						newEmail = email_text;
					} else {
						isEmailAvailable = false;
						available_message = "Email is taken";
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
		<h2>Change Email</h2>
		<TextBox
			label="New Email"
			placeholder="Enter your new email"
			value={newEmail}
			on:change={emailChangedHandler}
			{hasError}
			{errorMessage}
		/>
		{#if available_message.length > 0}
			{#if isEmailAvailable === true}
				<div class="available">{available_message}</div>
			{:else}
				<div class="not_available">{available_message}</div>
			{/if}
		{/if}
		<div></div>
		<div class="button_group">
			<SubmitButton
				title="Update"
				disabled={!isEmailAvailable}
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
