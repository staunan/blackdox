<header>
	<div class="page_title">Bookmark Manager</div>
	<div class="menu_items">
		<SvelteButton title="Create Bookmark" on:tap={openCreateBookmarkModal}></SvelteButton>
	</div>
</header>
<div class="bookmark_list">
    {#each bookmarks as bookmark}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="bookmark_item" on:click={openBookmarkDetailsModal(bookmark)}>
            <div class="bookmark_item_header">
                <span>{ bookmark.title }</span>
            </div>
        </div>
    {/each}
</div>

<!-- Bookmark Details View Modal -->
<ContentModal 
    title={selectedBookmark ? selectedBookmark.title : '' }
    active={detailsBookmarkModalIsActive} 
    overlayclose={true}
    on:close={closeDetailsBookmarkModal}
>
    <div slot="body">
        {#if selectedBookmark}
            <div class="text_info">
                <div class="info_label">Domain</div>
                <div class="info_value">{ selectedBookmark.domain }</div>
            </div>
            <div class="text_info">
                <div class="info_label">Username</div>
                <div class="info_value">{ selectedBookmark.username }</div>
            </div>
            <div class="text_info">
                <div class="info_label">Password</div>
                <div class="info_value">{ selectedBookmark.password }</div>
            </div>
            <div class="text_info">
                <div class="info_label">Note</div>
                <div class="info_value">{ selectedBookmark.note }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Delete Bookmark" on:tap={deleteBookmark}></SvelteButton>
    </div>
</ContentModal>

<!-- Create Bookmark Modal -->
<ContentModal 
    title="Create Bookmark"
    active={createBookmarkModalIsActive} 
    overlayclose={true}
    on:close={closeCreateBookmarkModal}
>
    <div slot="body">
        <Textbox label="Bookmark Title" placeholder="Title" on:change={titleChangeHandler} value={title}></Textbox>
        <Textbox label="URL Address" placeholder="Website Link" on:change={domainChangeHandler} value={domain}></Textbox>
        <Textbox label="Username" placeholder="Username" on:change={usernameChangeHandler} value={username}></Textbox>
        <Textbox label="Password" placeholder="Password" on:change={passwordChangeHandler} value={password}></Textbox>
        <TextArea label="Note" placeholder="Note" on:change={noteChangeHandler} value={note}></TextArea>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeCreateBookmarkModal}></SvelteButton>
        <SvelteButton color="blue" title="Submit Bookmark" on:tap={onCreateBookmarkSubmit}></SvelteButton>
    </div>
</ContentModal>

<script>
import { onMount } from 'svelte';
import {createBookmark, getBookmarks} from "apis/bookmark.js";
import SvelteButton from 'components/SvelteButton.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
let createBookmarkModalIsActive = false;
let detailsBookmarkModalIsActive = false;
let title = "";
let domain = "";
let username = "";
let password = "";
let note = "";
let bookmarks = [];
let selectedBookmark = null;

onMount(async () => {
    try{
        let response = await getBookmarks();
        bookmarks = response.bookmarks;
        console.log("Bookmarks:", bookmarks);
    }catch(error){
        console.log(error.response.data.message);
    }
});

function openCreateBookmarkModal(event){
    createBookmarkModalIsActive = true;
}
function openBookmarkDetailsModal(bookmark){
    selectedBookmark = bookmark;
    detailsBookmarkModalIsActive = true;
}
function closeCreateBookmarkModal(event){
    createBookmarkModalIsActive = false;
}
function closeDetailsBookmarkModal(event){
    detailsBookmarkModalIsActive = false;
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
async function onCreateBookmarkSubmit(event){
    let formData = {
        'title': title,
        'username': username,
        'domain': domain,
        'password': password,
        'note': note
    };
    try{
        let result = await createBookmark(formData);
        console.log("Result:", result);
    }catch(error){
        console.log(error.response.data.message);
    }
}
function deleteBookmark(){
    
}
</script>
<style>
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
    z-index: 10;
    padding-left: 20px;
}
.page_title{
    font-size: 25px;
    font-weight: bold;
    color: #fff;
}
.menu_items{
    display: flex;
    flex: 1;
    justify-content: flex-end;
    align-items: center;
    padding-right: 20px;
}
.bookmark_list{
    padding: 50px;
    padding-top: 110px;
}
.bookmark_item{
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    padding: 30px;
    margin-bottom: 20px;
    border-radius: 10px;
    cursor: pointer;
}
.text_info{
    margin-bottom: 30px;
}
.info_label{
    font-size: 13px;
    font-weight: bold;
}
.info_value{
    font-size: 24px;
    font-weight: 400;
}
</style>