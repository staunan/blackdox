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
                <span class="bookmark_title">{ bookmark.title }</span>
                <span class="bookmark_username">{ bookmark.username }</span>
            </div>
            <div class="copy_icon">

            </div>
        </div>
    {/each}
    {#if bookmarks.length === 0}
        <div class="bookmark_no_item">
            No Data Found
        </div>
    {/if}
    <div style="display: none;">
        <input type="text" id="copyToClipboard">
    </div>
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
                <div class="info_value"><a href={ selectedBookmark.domain } target="_blank">{ selectedBookmark.domain }</a></div>
            </div>
            <div class="text_info">
                <div class="info_label">Username</div>
                <div class="info_value">{ selectedBookmark.username }</div>
            </div>
            <div class="text_info">
                <div class="info_label">Password</div>
                <div class="info_value password_value">
                    <div class="password_protected">*****</div>
                    <!-- svelte-ignore a11y-click-events-have-key-events -->
                    <!-- svelte-ignore a11y-no-static-element-interactions -->
                    <div class="copy_password_button" on:click={copyTextToClipboard(selectedBookmark.password)}><i class="fa-solid fa-copy"></i></div>
                </div>
            </div>
            <div class="text_info">
                <div class="info_label">Note</div>
                <div class="info_value">{ selectedBookmark.note }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Delete Bookmark" on:tap={removeBookmark}></SvelteButton>
        <SvelteButton color="blue" title="Update Bookmark" on:tap={openUpdateBookmarkModal}></SvelteButton>
    </div>
</ContentModal>

<!-- Create/Update Bookmark Modal -->
<ContentModal 
    title={editBookmark ? "Update Bookmark" : "Create Bookmark"}
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
        {#if editBookmark}
            <SvelteButton color="blue" title="Update Bookmark" on:tap={onUpdateBookmarkSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Submit Bookmark" on:tap={onCreateBookmarkSubmit}></SvelteButton>
        {/if}
        
    </div>
</ContentModal>

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import {createBookmark, updateBookmark, getBookmarks, deleteBookmark} from "apis/bookmark.js";
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
let editBookmark = false;

onMount(() => {
    getAllBookmarks();
});

function openCreateBookmarkModal(){
    createBookmarkModalIsActive = true;
    editBookmark = false;
    setBookmarkData(null);
}
function openBookmarkDetailsModal(bookmark){
    selectedBookmark = bookmark;
    detailsBookmarkModalIsActive = true;
}
function closeCreateBookmarkModal(){
    createBookmarkModalIsActive = false;
}
function closeDetailsBookmarkModal(){
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
function openUpdateBookmarkModal(){
    editBookmark = true;
    createBookmarkModalIsActive = true;
    detailsBookmarkModalIsActive = false;
    setBookmarkData(selectedBookmark);
}
function setBookmarkData(bookmark){
    if(bookmark){
        title = bookmark.title;
        domain = bookmark.domain;
        username = bookmark.username;
        password = bookmark.password;
        note = bookmark.note;
    }else{
        title = "";
        domain = "";
        username = "";
        password = "";
        note = "";
    }
}
async function getAllBookmarks(){
    try{
        let response = await getBookmarks();
        bookmarks = response.bookmarks;
    }catch(error){
        console.log(error);
    }
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
        let response = await createBookmark(formData);
        closeCreateBookmarkModal();
        bookmarks = [response.bookmark, ...bookmarks];
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateBookmarkSubmit(event){
    let formData = {
        'bookmark_id': selectedBookmark.internalId,
        'title': title,
        'username': username,
        'domain': domain,
        'password': password,
        'note': note
    };
    try{
        let response = await updateBookmark(formData);
        closeCreateBookmarkModal();
        bookmarks = bookmarks.map(function(bookmark){
            if(selectedBookmark.internalId === bookmark.internalId){
                return response.bookmark;
            }else{
                return bookmark;
            }
        });
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function removeBookmark(){
    let formData = {
        'bookmark_id': selectedBookmark.internalId
    };
    try{
        let response = await deleteBookmark(formData);
        successToast(response.message);
        getAllBookmarks();
        closeDetailsBookmarkModal();
    }catch(error){
        console.log(error);
    }
}
function copyTextToClipboard(textToCopy){
    // Get the text field
    var copyText = document.getElementById("copyToClipboard");
    copyText.value = textToCopy;

    // Select the text field
    copyText.select();
    copyText.setSelectionRange(0, 99999); // For mobile devices

    // Copy the text inside the text field
    window.navigator['clipboard'].writeText(copyText.value);
    successToast("Password Copied");    
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
    display: flex;
}
.bookmark_no_item{
    display: flex;
    justify-content: center;
    align-items: center;
    margin: auto;
    width: 75%;
    height: 200px;
    border-radius: 10px;
    background-color: #ddd;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
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
.bookmark_title{
    font-size: 24px;
    font-weight: 400;
}
.bookmark_username{
    font-size: 14px;
    font-weight: bold;
}
.bookmark_item_header{
    display: flex;
    flex-direction: column;
    flex: 1;
}
.password_value{
    display: flex;
}
.password_protected{
    flex: 1;
    font-size: 40px;
}
.copy_password_button{
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
}
</style>