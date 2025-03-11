<script>
	import Card from "components/Card.svelte";
	import EmailBox from "components/form/EmailBox.svelte";
	import PasswordBox from "components/form/PasswordBox.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import LinkButton from "components/buttons/LinkButton.svelte";
	import { goto } from "$app/navigation";
	import { loginUser } from "apis/apis.js";
	import { store } from "store";

	let userEmail = "";
	let userEmailHasError = false;
	let userEmailErrorMessage = "";
	let userPassword = "";
	let userPasswordHasError = "";
	let userPasswordErrorMessage = "";
	let login_button_disabled = false;
	let login_button_text = "Login to Routinely";
	function emailChangedHandler(event) {
		userEmail = event.detail;
	}
	function passwordChangedHandler(event) {
		userPassword = event.detail;
	}
	function signupClickHandler() {
		goto("/signup");
	}
	function forgotPasswordClickHandler(event) {
		goto("/forgot-password");
	}
	async function loginButtonHandler(event) {
		// Login User --
		login_button_disabled = true;
		login_button_text = "Authenticating...";
		let loginDataObj = validateLoginForm();
		if (loginDataObj === false) {
			login_button_disabled = false;
			login_button_text = "Login to Routinely";
			console.log("Invalid Data");
			return;
		}
		try {
			let response = await loginUser(loginDataObj);
			setUserDetails.setUserDetails(response.Data);
			console.log(response.Data);
		} catch (error) {
			console.log(error);
		}
	}
	function validateLoginForm() {
		let userObj = {};
		let hasError = false;
		// User Email --
		if (userEmail == "") {
			hasError = true;
			userEmailHasError = true;
			userEmailErrorMessage = "Email is required";
		} else if (!validateEmail(userEmail)) {
			// Check if email is valid format --
			hasError = true;
			userEmailHasError = true;
			userEmailErrorMessage = "Email format is invalid";
		} else {
			userEmailHasError = false;
			userEmailErrorMessage = "";
			userObj.email = userEmail;
		}
		// User Password --
		if (userPassword == "") {
			hasError = true;
			userPasswordHasError = true;
			userPasswordErrorMessage = "Password is required";
		} else if (!validatePassword(userPassword)) {
			hasError = true;
			userPasswordHasError = true;
			userPasswordErrorMessage =
				"Invalid Password: Your password should be at least 8 character long, it must have at least 1 symbol (Like, $, &, %, *, @ etc), at least 1 upper case letter and at least 1 lower case letter and at least a number (Like, 1, 3, 6 etc)";
		} else {
			userPasswordHasError = false;
			userPasswordErrorMessage = "";
			userObj.password = userPassword;
		}
		if (hasError) {
			return false;
		} else {
			return userObj;
		}
	}
	const validateEmail = (email) => {
		return String(email)
			.toLowerCase()
			.match(
				/^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|.(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
			);
	};
	function validatePassword(password) {
		var re = /^(?=.*\d)(?=.*[!@#$%^&*])(?=.*[a-z])(?=.*[A-Z]).{8,}$/;
		return re.test(password);
	}
</script>

<Card>
	<div class="card_body">
		<h1 class="center mb10 form_heading">Login to Routinely</h1>
		<form>
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
			<div class="center mt10 submit_button_container">
				<SubmitButton
					title={login_button_text}
					on:tap={() => {
						loginButtonHandler();
					}}
					disabled={login_button_disabled}
				></SubmitButton>
			</div>

			<div class="row">
				<div class="signup_section">
					<LinkButton
						label="Create an account"
						on:tap={signupClickHandler}
					></LinkButton>
				</div>
				<div class="forgot_password_section">
					<LinkButton
						label="Forgot Password?"
						on:tap={forgotPasswordClickHandler}
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
	.forgot_password_section {
		flex: 1;
		display: flex;
		justify-content: flex-end;
		align-items: center;
	}
</style>
