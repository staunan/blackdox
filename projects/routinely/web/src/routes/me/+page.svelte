<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import MyAccountPage from "components/pages/me/MyAccountPage.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Card from "components/Card.svelte";
	import Right from "components/layouts/Right.svelte";
	import LogoutButtonIcon from "components/svg/LogoutButtonIcon.svelte";
	import ResetPasswordIcon from "components/svg/ResetPasswordIcon.svelte";
	import UsernameIcon from "components/svg/UsernameIcon.svelte";

	import { logoutUser } from "apis/apis.js";
	async function logoutHandler() {
		let response = await logoutUser();
		if (!response.HasError) {
			goto("/login");
		} else {
			console.log("Error while logging out");
			console.log(response);
		}
	}
</script>

<RoutinelyPageContainer>
	<MyAccountPage></MyAccountPage>
	<div class="update_username_section">
		<Card padding={4} background="#0000001a">
			<div class="section_title">
				<UsernameIcon></UsernameIcon>
				<div class="section_title_text">Change Email</div>
			</div>
			<p class="section_title_desc">
				Change your primary email which is used for log in
			</p>
			<Right>
				<SubmitButton
					title="Change Email"
					on:tap={logoutHandler}
					color="green"
				>
					<UsernameIcon size="30px"></UsernameIcon>
				</SubmitButton>
			</Right>
		</Card>
	</div>
	<div class="update_username_section">
		<Card padding={4} background="#0000001a">
			<div class="section_title">
				<UsernameIcon></UsernameIcon>
				<div class="section_title_text">Update Username</div>
			</div>
			<p class="section_title_desc">Update your routinely username</p>
			<Right>
				<SubmitButton
					title="Update Username"
					on:tap={logoutHandler}
					color="green"
				>
					<UsernameIcon size="30px"></UsernameIcon>
				</SubmitButton>
			</Right>
		</Card>
	</div>
	<div class="update_password_section">
		<Card padding={4} background="#0000001a">
			<div class="section_title">
				<ResetPasswordIcon></ResetPasswordIcon>
				<div class="section_title_text">Reset Password</div>
			</div>
			<p class="section_title_desc">
				Update your routinely password to secure your account.
			</p>
			<Right>
				<SubmitButton title="Reset Password" on:tap={logoutHandler}
					><ResetPasswordIcon size="30px"
					></ResetPasswordIcon></SubmitButton
				>
			</Right>
		</Card>
	</div>
	<div class="edit_routine_section">
		<Card padding={4} background="#0000001a">
			<div class="section_title">
				<LogoutButtonIcon></LogoutButtonIcon>
				<div class="section_title_text">Logout</div>
			</div>
			<p class="section_title_desc">Logout from Routinely</p>
			<Right>
				<SubmitButton title="Log out" color="red" on:tap={logoutHandler}
					><LogoutButtonIcon size="30px"
					></LogoutButtonIcon></SubmitButton
				>
			</Right>
		</Card>
	</div>
</RoutinelyPageContainer>

<style>
	.edit_routine_section,
	.update_password_section,
	.update_username_section,
	.update_username_section {
		padding-top: 30px;
	}
	.section_title {
		display: flex;
		justify-content: flex-start;
		align-items: center;
		width: auto;
	}
	.section_title_text {
		padding-left: 20px;
		font-size: 24px;
		font-weight: bold;
	}
	.section_title_desc {
		font-size: 20px;
	}
</style>
