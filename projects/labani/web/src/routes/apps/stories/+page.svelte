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
                <span class="story_content">{ story.story }</span>
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
    <div slot="body">
        {#if selectedStory}
            <div class="text_info">
                <div class="info_value">{ selectedStory.story }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Delete Story" on:tap={removeStory}></SvelteButton>
        <SvelteButton color="blue" title="Update Story" on:tap={openUpdateStoryModal}></SvelteButton>
    </div>
</ContentModal>

<!-- Create/Update Story Modal -->
<ContentModal 
    title={editStory ? "Update Story" : "New Story"}
    active={newStoryModalIsActive} 
    overlayclose={true}
    on:close={closeNewStoryModal}
>
    <div slot="body">
        <TextArea label="Story" placeholder="Write your story..." on:change={storyContentChangeHandler} value={story}></TextArea>
        <QuillEditor></QuillEditor>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeNewStoryModal}></SvelteButton>
        {#if editStory}
            <SvelteButton color="blue" title="Update Story" on:tap={onUpdateStorySubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Story" on:tap={onNewStorySubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import {createStory, updateStory, getStories, deleteStory} from "apis/stories.js";
import SvelteButton from 'components/SvelteButton.svelte';
import QuillEditor from 'components/QuillEditor.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
let newStoryModalIsActive = false;
let detailsStoryModalIsActive = false;

let story = "";

let stories = [];
let selectedStory = null;
let editStory = false;

onMount(() => {
    getAllStories();
});

function openNewStoryModal(){
    newStoryModalIsActive = true;
    editStory = false;
    setStoryData(null);
}
function closeNewStoryModal(){
    newStoryModalIsActive = false;
}
function openStoryDetailsModal(story){
    selectedStory = story;
    detailsStoryModalIsActive = true;
}
function closeStoryDetailsModal(){
    detailsStoryModalIsActive = false;
}
function storyContentChangeHandler(event){
    story = event.detail;
}
function openUpdateStoryModal(){
    editStory = true;
    newStoryModalIsActive = true;
    detailsStoryModalIsActive = false;
    setStoryData(selectedStory);
}
function setStoryData(story_data){
    if(story_data){
        story = story_data.story;
    }else{
        story = "";
    }
}
async function getAllStories(){
    try{
        let response = await getStories();
        stories = response.stories;
    }catch(error){
        console.log(error);
    }
}
async function onNewStorySubmit(event){
    let formData = {
        'story': story
    };
    try{
        let response = await createStory(formData);
        closeNewStoryModal();
        stories = [response.story, ...stories];
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateStorySubmit(event){
    let formData = {
        'story_id': selectedStory.internalId,
        'story': story
    };
    try{
        let response = await updateStory(formData);
        closeNewStoryModal();
        stories = stories.map(function(story){
            if(selectedStory.internalId === story.internalId){
                return response.story;
            }else{
                return story;
            }
        });
        successToast(response.message);
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
.info_label{
    font-size: 13px;
    font-weight: bold;
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
</style>