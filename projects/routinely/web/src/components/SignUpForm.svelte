<script>
	import Card from "components/Card.svelte";
	import TextBox from "components/form/TextBox.svelte";
	import EmailBox from "components/form/EmailBox.svelte";
	import PasswordBox from "components/form/PasswordBox.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import LinkButton from "components/buttons/LinkButton.svelte";
	import { goto } from "$app/navigation";
	import { createAccount } from "apis/apis.js";
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	let username = "";
	let usernameHasError = "";
	let usernameErrorMessage = "";
	let userFullName = "";
	let userFullNameHasError = false;
	let userFullNameErrorMessage = "";
	let userEmail = "";
	let userEmailHasError = false;
	let userEmailErrorMessage = "";
	let userPassword = "";
	let userPasswordHasError = "";
	let userPasswordErrorMessage = "";
	let confirmPassword = "";
	let confirmPasswordHasError = "";
	let confirmPasswordErrorMessage = "";

	function userFullNameChangedHandler(event) {
		userFullName = event.detail;
		if (userFullName && userFullNameHasError) {
			userFullNameHasError = false;
			userFullNameErrorMessage = "";
		}
	}
	function usernameChangedHandler(event) {
		username = event.detail;
		if (username && usernameHasError) {
			usernameHasError = false;
			usernameErrorMessage = "";
		}
	}
	function emailChangedHandler(event) {
		userEmail = event.detail;
		if (userEmail && userEmailHasError) {
			userEmailHasError = false;
			userEmailErrorMessage = "";
		}
	}
	function passwordChangedHandler(event) {
		userPassword = event.detail;
		if (userPassword && userPasswordHasError) {
			userPasswordHasError = false;
			userPasswordErrorMessage = "";
		}
	}
	function confirmPasswordChangedHandler(event) {
		confirmPassword = event.detail;
		if (confirmPassword && confirmPasswordHasError) {
			confirmPasswordHasError = false;
			confirmPasswordErrorMessage = "";
		}
	}

	async function createAccountButtonHandler() {
		let userObj = validateUserData();
		if (userObj === false) {
			console.log("Invalid Form");
			return;
		}
		try {
			let response = await createAccount(userObj);
			let user = response.Data;
			dispatch("created", user);
		} catch (error) {
			console.log(error);
		}
	}
	function loginButtonClickHandler(event) {
		goto("/login");
	}
	function validateUserData() {
		let userObj = {};
		let hasError = false;
		// User Full Name --
		if (userFullName == "") {
			hasError = true;
			userFullNameHasError = true;
			userFullNameErrorMessage = "Please provide your full name";
		} else {
			userFullNameHasError = false;
			userFullNameErrorMessage = "";
			userObj.fullname = userFullName;
		}
		// Username --
		if (username == "") {
			hasError = true;
			usernameHasError = true;
			usernameErrorMessage = "Please provide a unique username";
		} else {
			usernameHasError = false;
			usernameErrorMessage = "";
			userObj.username = username;
		}
		// User Email --
		if (userEmail == "") {
			hasError = true;
			userEmailHasError = true;
			userEmailErrorMessage = "Please provide email";
		} else {
			userFullNameHasError = false;
			userFullNameErrorMessage = "";
			userObj.email = userEmail;
		}
		// User Password --
		if (userPassword == "") {
			hasError = true;
			userPasswordHasError = true;
			userPasswordErrorMessage = "Please provide password";
		} else {
			userPasswordHasError = false;
			userPasswordErrorMessage = "";
			userObj.password = userPassword;
		}
		// Confirm Password --
		if (confirmPassword == "") {
			hasError = true;
			confirmPasswordHasError = true;
			confirmPasswordErrorMessage = "Please confirm your password";
		} else if (userPassword != confirmPassword) {
			hasError = true;
			confirmPasswordHasError = true;
			confirmPasswordErrorMessage = "Passwords do not match";
		} else {
			confirmPasswordHasError = false;
			confirmPasswordErrorMessage = "";
		}
		if (hasError) {
			return false;
		} else {
			return userObj;
		}
	}
</script>

<Card>
	<div class="card_body">
		<h1 class="center mb10 form_heading">Create new account</h1>
		<form>
			<TextBox
				label="Name"
				placeholder="Full Name"
				on:change={userFullNameChangedHandler}
				value={userFullName}
				hasError={userFullNameHasError}
				errorMessage={userFullNameErrorMessage}
			></TextBox>

			<div class="form_gap"></div>
			<TextBox
				label="Username"
				placeholder="Username"
				on:change={usernameChangedHandler}
				value={username}
				hasError={usernameHasError}
				errorMessage={usernameErrorMessage}
			></TextBox>

			<div class="form_gap"></div>
			<EmailBox
				label="Email"
				placeholder="Email"
				on:change={emailChangedHandler}
				value={userEmail}
				hasError={userEmailHasError}
				errorMessage={userEmailErrorMessage}
			></EmailBox>

			<div class="form_gap"></div>
			<PasswordBox
				label="Password"
				placeholder="Password"
				on:change={passwordChangedHandler}
				value={userPassword}
				hasError={userPasswordHasError}
				errorMessage={userPasswordErrorMessage}
			></PasswordBox>

			<div class="form_gap"></div>
			<PasswordBox
				label="Confirm Password"
				placeholder="Retype your Password"
				on:change={confirmPasswordChangedHandler}
				value={confirmPassword}
				hasError={confirmPasswordHasError}
				errorMessage={confirmPasswordErrorMessage}
			></PasswordBox>

			<div class="form_gap"></div>
			<div class="center mt10 submit_button_container">
				<SubmitButton
					title="Create account"
					on:tap={createAccountButtonHandler}
				></SubmitButton>
			</div>

			<div class="row">
				<div class="signup_section">
					<LinkButton
						label="Already have an account? Login"
						on:tap={loginButtonClickHandler}
					></LinkButton>
				</div>
			</div>
		</form>
	</div>
</Card>

<style>
	.card_body {
		padding: 20px;
	}
	.center {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.mb10 {
		margin-bottom: 10px;
	}
	.mt10 {
		margin-top: 10px;
	}
	.form_heading {
		font-family: monospace;
		font-size: 24px;
		font-weight: bold;
		margin-top: 10px;
		padding-bottom: 40px;
	}
	.submit_button_container {
		margin-bottom: 10px;
		display: flex;
		justify-content: center;
		align-items: center;
		padding-bottom: 10px;
	}
	.form_gap {
		width: 100%;
		height: 20px;
	}
	.row {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.signup_section {
		flex: 1;
		display: flex;
		justify-content: flex-start;
		align-items: center;
	}
</style>
