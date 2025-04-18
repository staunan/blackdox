<script>
	import RoutinelyPageContainer from "components/RoutinelyPageContainer.svelte";
	import MyAccountPage from "components/pages/me/MyAccountPage.svelte";
	import SubmitButton from "components/buttons/SubmitButton.svelte";
	import Card from "components/Card.svelte";
	import Right from "components/layouts/Right.svelte";
	import Left from "components/layouts/Left.svelte";
	import LogoutButtonIcon from "components/svg/LogoutButtonIcon.svelte";
	import ResetPasswordIcon from "components/svg/ResetPasswordIcon.svelte";
	import UsernameIcon from "components/svg/UsernameIcon.svelte";
	import EmailIcon from "components/svg/EmailIcon.svelte";
	import EditPenIcon from "components/svg/EditPenIcon.svelte";
	import { store, user_details } from "store";
	import { goto } from "$app/navigation";
	import { logoutUser } from "apis/apis.js";

	let user = null;
	user_details.subscribe((v) => {
		if (v) {
			user = v;
		}
	});

	async function logoutHandler() {
		let response = await logoutUser();
		if (!response.HasError) {
			store.logout();
			goto("/");
		} else {
			console.log("Error while logging out");
			console.log(response);
		}
	}
</script>

<RoutinelyPageContainer>
	{#if user}
		<MyAccountPage></MyAccountPage>
		<div class="update_username_section">
			<Card padding={4} background="#0000001a">
				<div class="section_title">
					<EmailIcon size="25px"></EmailIcon>
					<div class="section_title_text">Email</div>
				</div>
				<div class="section_title_desc">
					<div class="email">{user.Email}</div>
				</div>
				<Left>
					<div class="change_email">
						Change your primary email which is used for log in
					</div>
				</Left>
				<Left>
					<SubmitButton
						title="Change Email"
						on:tap={logoutHandler}
						color="greyblue"
					>
						<EditPenIcon size="20px"></EditPenIcon>
					</SubmitButton>
				</Left>
			</Card>
		</div>
		<div class="update_username_section">
			<Card padding={4} background="#0000001a">
				<div class="section_title">
					<UsernameIcon size="25px"></UsernameIcon>
					<div class="section_title_text">Username</div>
				</div>
				<div class="section_title_desc">
					<div class="email">{user.Username}</div>
				</div>
				<Left>
					<div class="change_email">
						Update your routinely username
					</div>
				</Left>
				<Left>
					<SubmitButton
						title="Update Username"
						on:tap={logoutHandler}
						color="greyblue"
					>
						<EditPenIcon size="20px"></EditPenIcon>
					</SubmitButton>
				</Left>
			</Card>
		</div>
		<div class="update_password_section">
			<Card padding={4} background="#0000001a">
				<div class="section_title">
					<ResetPasswordIcon size="25px"></ResetPasswordIcon>
					<div class="section_title_text">Reset Password</div>
				</div>
				<Left>
					<div class="change_email" style="padding-top: 50px;">
						Update your routinely password to secure your account.
					</div>
				</Left>
				<Left>
					<SubmitButton
						title="Reset Password"
						on:tap={logoutHandler}
						color="greyblue"
					>
						<ResetPasswordIcon size="20px"></ResetPasswordIcon>
					</SubmitButton>
				</Left>
			</Card>
		</div>
		<div class="edit_routine_section">
			<Card padding={4} background="#0000001a">
				<div class="section_title">
					<div style="position: relative; top: 4px;">
						<LogoutButtonIcon size="25px"></LogoutButtonIcon>
					</div>
					<div class="section_title_text">Logout</div>
				</div>
				<Left>
					<div class="change_email" style="padding-top: 50px;">
						Logout from Routinely
					</div>
				</Left>
				<Left>
					<SubmitButton
						title="Logout"
						on:tap={logoutHandler}
						color="red"
					></SubmitButton>
				</Left>
			</Card>
		</div>
	{/if}
</RoutinelyPageContainer>

<style>
	.edit_routine_section,
	.update_password_section,
	.update_username_section,
	.update_username_section {
		padding-top: 20px;
		padding-bottom: 20px;
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
	.email {
		padding-top: 10px;
		padding-bottom: 40px;
		font-size: 18px;
		font-weight: bold;
	}
	.change_email {
		font-size: 16px;
		padding-bottom: 10px;
	}
</style>
