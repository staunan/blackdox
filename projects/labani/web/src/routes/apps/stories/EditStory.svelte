<ContentModal 
    title={edit ? "Update Story" : "New Story"}
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body">
        <div class="story_editor_form">
            <div class="rich_text_editor">
                <div class="form_label">Write your story</div>
                <QuillEditor content={default_story} on:change={quillContentChangeHandler}></QuillEditor>
            </div>
            <div class="tag_selector">
                <div class="form_label">Select Tags</div>
                <TagSelector selected={selectedTags}></TagSelector>
            </div>
        </div>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeModal}></SvelteButton>
        {#if edit}
            <SvelteButton color="blue" title="Update Story" on:tap={handleUpdateStoryButtonClick}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Story" on:tap={handleCreateStoryButtonClick}></SvelteButton>
        {/if}
    </div>
</ContentModal>
<script>
import QuillEditor from 'components/QuillEditor.svelte';
import ContentModal from 'components/ContentModal.svelte';
import SvelteButton from 'components/SvelteButton.svelte';
import TagSelector from './TagSelector.svelte';
import {createEventDispatcher} from 'svelte';
import {createStory, updateStory} from "apis/stories.js";
import {successToast} from "lib/js/toast.js";
const dispatch = createEventDispatcher();

export let active = false;
export let edit = false;
export let story = "";
let default_story = "";
let quill_story_content = "";
let selectedTags = ["Tag 2"];

$: {
    if(story){
        default_story = story.story;
    }else{
        default_story = "";
    }
}

function closeModal(){
    dispatch("close");
}
function quillContentChangeHandler(event){
    quill_story_content = event.detail;
}
async function handleCreateStoryButtonClick(event){
    let formData = {
        'story': quill_story_content
    };
    try{
        let response = await createStory(formData);
        closeModal();
        dispatch("created", response.story)
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function handleUpdateStoryButtonClick(event){
    let formData = {
        'story_id': story.internalId,
        'story': quill_story_content
    };
    try{
        let response = await updateStory(formData);
        closeModal();
        dispatch("updated", response.story);
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
</script>
<style>
.rich_text_editor{
    padding: 20px;
}
.tag_selector{
    padding: 20px;
}
</style>