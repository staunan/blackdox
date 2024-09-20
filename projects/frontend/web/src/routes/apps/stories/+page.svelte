<header>
	<div class="page_title">Stories</div>
	<div class="menu_items">
		<SvelteButton title="New Story" on:tap={openNewStoryModal}></SvelteButton>
	</div>
</header>
<div class="story_list">
    {#each stories as story}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="story_item" on:click={openStoryDetailsModal(story)}>
            <div class="story_item_header">
                <div class="story_content"><QuillView content={story.story}></QuillView></div>
                <div title={story.createdAt}>{ getRelateiveTime(story.createdAt) }</div>
            </div>
        </div>
    {/each}
    {#if stories.length === 0}
        <div class="story_no_item">
            No Story Found
        </div>
    {/if}
</div>

<!-- Story Details Modal -->
<ContentModal 
    title="Story Details"
    active={detailsStoryModalIsActive} 
    overlayclose={true}
    on:close={closeStoryDetailsModal}
>
    <div slot="body" class="details_body">
        {#if selectedStory}
            <div class="text_info">
                <div class="info_value">
                    <QuillView content={selectedStory.story}></QuillView>
                </div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <DeleteButton on:tap={removeStory} title="Delete Story"></DeleteButton>
        <SvelteButton color="blue" title="Update Story" on:tap={openUpdateStoryModal}></SvelteButton>
    </div>
</ContentModal>

<!-- Create/Update Story Modal -->
<EditStory 
    active={editStoryModalIsActive}
    edit={editStory}
    story={selectedStory}
    on:close={closeNewStoryModal}
    on:created={onNewStoryCreated}
    on:updated={onStoryUpdated}
></EditStory>

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import {getStories, deleteStory} from "apis/stories.js";
import SvelteButton from 'components/SvelteButton.svelte';
import DeleteButton from 'components/DeleteButton.svelte';
import QuillView from 'components/QuillView.svelte';
import ContentModal from 'components/ContentModal.svelte';
import EditStory from './EditStory.svelte';

let editStoryModalIsActive = false;
let detailsStoryModalIsActive = false;

let stories = [];
let selectedStory = null;
let editStory = false;

onMount(async () => {
    getAllStories();
});
function openNewStoryModal(){
    selectedStory = null;
    editStory = false;
    editStoryModalIsActive = true;
}
function openUpdateStoryModal(){
    editStory = true;
    editStoryModalIsActive = true;
    detailsStoryModalIsActive = false;
}
function closeNewStoryModal(){
    editStoryModalIsActive = false;
}
function onNewStoryCreated(event){
    let newStory = event.detail;
    stories = [newStory, ...stories];
}
function onStoryUpdated(event){
    let updatedStory = event.detail;
    stories = stories.map(function(storyItem){
        if(updatedStory.internalId === storyItem.internalId){
            return updatedStory;
        }else{
            return storyItem;
        }
    });
}
function openStoryDetailsModal(story){
    selectedStory = story;
    detailsStoryModalIsActive = true;
}
function closeStoryDetailsModal(){
    detailsStoryModalIsActive = false;
}
async function getAllStories(){
    try{
        let response = await getStories();
        stories = response.stories;
    }catch(error){
        console.log(error);
    }
}
async function removeStory(){
    let formData = {
        'story_id': selectedStory.internalId
    };
    try{
        let response = await deleteStory(formData);
        successToast(response.message);
        getAllStories();
        closeStoryDetailsModal();
    }catch(error){
        console.log(error);
    }
}
function getRelateiveTime(datetime) {
    var msPerMinute = 60 * 1000;
    var msPerHour = msPerMinute * 60;
    var msPerDay = msPerHour * 24;
    var msPerMonth = msPerDay * 30;
    var msPerYear = msPerDay * 365;

    var currentDate = Date.parse(new Date());
    var previousDate = Date.parse(new Date(datetime));

    var elapsed = currentDate - previousDate;

    if (elapsed < msPerMinute) {
        return Math.round(elapsed/1000) + ' seconds ago';   
    }else if (elapsed < msPerHour) {
        return Math.round(elapsed/msPerMinute) + ' minutes ago';   
    }else if (elapsed < msPerDay ) {
        return Math.round(elapsed/msPerHour ) + ' hours ago';   
    }else if (elapsed < msPerMonth) {
        return Math.round(elapsed/msPerDay) + ' days ago';   
    }else if (elapsed < msPerYear) {
        return Math.round(elapsed/msPerMonth) + ' months ago';   
    }else {
        return Math.round(elapsed/msPerYear ) + ' years ago';   
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
.story_list{
    padding: 50px;
    padding-top: 110px;
}
.story_item{
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    padding: 30px;
    margin-bottom: 20px;
    border-radius: 10px;
    cursor: pointer;
    display: flex;
}
.story_no_item{
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
.info_value{
    font-size: 24px;
    font-weight: 400;
}
.story_content{
    font-size: 24px;
    font-weight: 400;
}
.story_item_header{
    display: flex;
    flex-direction: column;
    flex: 1;
}
.details_body{
    padding: 20px;
}
</style>