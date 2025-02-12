<section>
    <div class="menu_bar_container">
        <MenuBar></MenuBar>
    </div>
	<div class="signup_form">
		<div class="form_title">Create Account</div>

        <Textbox label="First Name" placeholder="First Name" on:change={firstNameChangeHandler} value={first_name}></Textbox>
        <Textbox label="Last Name" placeholder="Last Name" on:change={lastNameChangeHandler} value={last_name}></Textbox>
        <Textbox label="User Name" placeholder="User Name" on:change={userNameChangeHandler} value={user_name}></Textbox>
        <Textbox label="User Email" placeholder="User Email" on:change={userEmailChangeHandler} value={user_email}></Textbox>
        

	</div>
</section>
<script>
import MenuBar from 'components/pages/MenuBar.svelte';
import Textbox from 'components/Textbox.svelte';

let first_name = "";
let last_name = "";
let user_name = "";
let user_email = "";

function firstNameChangeHandler(event){
    first_name = event.detail;
}
function lastNameChangeHandler(event){
    last_name = event.detail;
}
function userNameChangeHandler(event){
    user_name = event.detail;
}
function userEmailChangeHandler(event){
    user_email = event.detail;
}
</script>
<style>
.menu_bar_container{
    padding-top: 40px;
}
.signup_form{
    width: 50%;
    margin: auto;
    padding-top: 10px;
    border-radius: 5px;
    border: 1px solid #ddd;
    padding: 10px;
}
.form_title{
    display: flex;
    justify-content: center;
    padding-bottom: 20px;
    font-weight: bold;
    text-transform: uppercase;
}
</style>
