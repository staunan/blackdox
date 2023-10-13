<header>
	<div class="website_title">Bookmark Manager</div>
	<div class="menu_items">
		<SvelteButton title="Create Bookmark" on:tap={buttonTap}></SvelteButton>
		<FormModal 
			title="Create Bookmark"
			active={createBookmarkModalIsActive} 
			overlayclose={true}
			on:close={closeFormModal}
		>
			<div slot="body">
				<Textbox label="Bookmark Title" placeholder="Title" on:change={titleChangeHandler} value={title}></Textbox>
				<Textbox label="URL Address" placeholder="Website Link" on:change={domainChangeHandler} value={domain}></Textbox>
				<Textbox label="Username" placeholder="Username" on:change={usernameChangeHandler} value={username}></Textbox>
				<Textbox label="Password" placeholder="Password" on:change={passwordChangeHandler} value={password}></Textbox>
				<TextArea label="Note" placeholder="Note" on:change={noteChangeHandler} value={note}></TextArea>
			</div>
			<div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
				<SvelteButton color="red" title="Cancel" on:tap={closeFormModal}></SvelteButton>
				<SvelteButton color="blue" title="Submit Bookmark" on:tap={onCreateBookmarkSubmit}></SvelteButton>
			</div>
		</FormModal>
	</div>
</header>
<script>
import SvelteButton from 'components/SvelteButton.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import FormModal from 'components/FormModal.svelte';
let createBookmarkModalIsActive = false;
let title = "";
let domain = "";
let username = "";
let password = "";
let note = "";
function buttonTap(event){
	createBookmarkModalIsActive = true;
}
function closeFormModal(event){
	createBookmarkModalIsActive = false;
}
function titleChangeHandler(event){
	title = event.detail;
}
function domainChangeHandler(event){
	domain = event.detail;
}
function usernameChangeHandler(event){
	username = event.detail;
}
function passwordChangeHandler(event){
	password = event.detail;
}
function noteChangeHandler(event){
	note = event.detail;
}
function onCreateBookmarkSubmit(event){
	console.log("Working");
}
</script>
<style>
.website_title{
	font-size: 25px;
	font-weight: bold;
	color: #fff;
}
header{
	display: flex;
	justify-content: flex-start;
	background: #C33764;  /* fallback for old browsers */
	background: -webkit-linear-gradient(to right, #1D2671, #C33764);  /* Chrome 10-25, Safari 5.1-6 */
	background: linear-gradient(to right, #1D2671, #C33764); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
	align-items: center;
	height: 60px;
	position: fixed;
	left: 0;
	right: 0;
	top: 0;
	z-index: 10001;
    padding-left: 20px;
}
.menu_items{
    display: flex;
    flex: 1;
    justify-content: flex-end;
    align-items: center;
	padding-right: 20px;
}
</style>