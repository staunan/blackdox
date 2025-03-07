<script>
	import Card from "components/Card.svelte";
	import EmailBox from "components/form/EmailBox.svelte";
	import PasswordBox from "components/form/PasswordBox.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import LinkButton from "components/buttons/LinkButton.svelte";
	import { goto } from "$app/navigation";

	let userEmail = "";
	let userEmailHasError = false;
	let userEmailErrorMessage = "";
	let userPassword = "";
	let userPasswordHasError = "";
	let userPasswordErrorMessage = "";
	function emailChangedHandler(event) {
		console.log("Implementation Pending");
	}

	function loginButtonHandler(event) {
		console.log("Implementation Pending");
	}
	function passwordChangedHandler(event) {
		console.log("Implementation Pending");
	}
	function signupClickHandler() {
		goto("/signup");
	}
	function forgotPasswordClickHandler(event) {
		console.log("Pending");
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
					title="Login to Routinely"
					on:tap={loginButtonHandler}
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
