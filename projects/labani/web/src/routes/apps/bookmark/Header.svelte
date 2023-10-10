<header>
	<div class="website_title">Bookmark Manager</div>
	<div class="menu_items">
		<Button title="Create Bookmark" on:tap={buttonTap}></Button>
		<FormModal 
			active={createBookmarkModalIsActive} 
			overlayclose={true}
			on:close={closeFormModal}
		>
			<div slot="header"></div>
		</FormModal>
	</div>
</header>
<script>
import Button from 'components/SvelteButton.svelte';
import FormModal from 'components/FormModal.svelte';
let createBookmarkModalIsActive = false;
function buttonTap(event){
	createBookmarkModalIsActive = true;
}
function closeFormModal(event){
	createBookmarkModalIsActive = false;
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