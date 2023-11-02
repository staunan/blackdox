<header>
	<div class="page_title">Spending Tracker</div>
	<div class="menu_items">
		<SvelteButton title="New Spending" on:tap={openNewSpendingModal}></SvelteButton>
	</div>
</header>
<div class="bookmark_list">
    {#each spendings as spending}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="bookmark_item" on:click={openSpendingDetailsModal(spending)}>
            <div class="bookmark_item_header">
                <span class="bookmark_title">{ spending.item_name }</span>
                <span class="bookmark_username">Rs. { spending.price } /-</span>
            </div>
        </div>
    {/each}
    {#if spendings.length === 0}
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
    title={selectedSpending ? selectedSpending.item_name : '' }
    active={detailsSpendingkModalIsActive} 
    overlayclose={true}
    on:close={closeDetailsSpendingModal}
>
    <div slot="body">
        {#if selectedSpending}
            <div class="text_info">
                <div class="info_label">Price</div>
                <div class="info_value">Rs. { selectedSpending.price } /-</div>
            </div>
            <div class="text_info">
                <div class="info_label">Quantity</div>
                <div class="info_value">{ selectedSpending.quantity }</div>
            </div>
            <div class="text_info">
                <div class="info_label">Note</div>
                <div class="info_value">{ selectedSpending.note }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Delete Spending" on:tap={removeSpending}></SvelteButton>
        <SvelteButton color="blue" title="Update Spending" on:tap={openUpdateBookmarkModal}></SvelteButton>
    </div>
</ContentModal>

<!-- Create/Update Bookmark Modal -->
<ContentModal 
    title={editSpending ? "Update Bookmark" : "Create Bookmark"}
    active={newSpendingModalIsActive} 
    overlayclose={true}
    on:close={closeNewSpendingModal}
>
    <div slot="body">
        <Textbox label="Item Name" placeholder="Item Name" on:change={itemNameChangeHandler} value={item_name}></Textbox>
        <Textbox label="Price in Rupees" placeholder="Price in Rupees" on:change={priceChangeHandler} value={price}></Textbox>
        <Textbox label="Quantity" placeholder="Quantity" on:change={quantityChangeHandler} value={quantity}></Textbox>
        <TextArea label="Note" placeholder="Note" on:change={noteChangeHandler} value={note}></TextArea>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeNewSpendingModal}></SvelteButton>
        {#if editSpending}
            <SvelteButton color="blue" title="Update Spending" on:tap={onUpdateSpendingSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Submit Spending" on:tap={onNewSpendingSubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import {newSpending, getSpendings, deleteSpending} from "apis/spending.js";
import SvelteButton from 'components/SvelteButton.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
let newSpendingModalIsActive = false;
let detailsSpendingkModalIsActive = false;

let item_name = "";
let price = "";
let quantity = "";
let note = "";

let spendings = [];
let selectedSpending = null;
let editSpending = false;

onMount(() => {
    getAllSpendings();
});

function openNewSpendingModal(){
    newSpendingModalIsActive = true;
    editSpending = false;
    setBookmarkData(null);
}
function closeNewSpendingModal(){
    newSpendingModalIsActive = false;
}
function openSpendingDetailsModal(spending){
    selectedSpending = spending;
    detailsSpendingkModalIsActive = true;
}
function closeDetailsSpendingModal(){
    detailsSpendingkModalIsActive = false;
}
function itemNameChangeHandler(event){
    item_name = event.detail;
}
function priceChangeHandler(event){
    price = event.detail;
}
function quantityChangeHandler(event){
    quantity = event.detail;
}
function noteChangeHandler(event){
    note = event.detail;
}
function openUpdateBookmarkModal(){
    editSpending = true;
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
async function getAllSpendings(){
    try{
        let response = await getSpendings();
        spendings = response.spendings;
    }catch(error){
        console.log(error);
    }
}
async function onNewSpendingSubmit(event){
    let formData = {
        'item_name': item_name,
        'price': price,
        'quantity': quantity,
        'note': note
    };
    try{
        let response = await newSpending(formData);
        closeNewSpendingModal();
        spendings = [response.spending, ...spendings];
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateSpendingSubmit(event){
    let formData = {
        'spending_id': selectedSpending.internalId,
        'item_name': item_name,
        'price': price,
        'quantity': quantity,
        'note': note
    };
    try{
        let response = await updateSpending(formData);
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
async function removeSpending(){
    let formData = {
        'spending_id': selectedSpending.internalId
    };
    try{
        let response = await deleteSpending(formData);
        successToast(response.message);
        getAllSpendings();
        closeDetailsSpendingModal();
    }catch(error){
        console.log(error);
    }
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