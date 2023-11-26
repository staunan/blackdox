<ContentModal 
    title={edit ? "Update Story" : "New Story"}
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body">
        <QuillEditor content={default_story} on:change={quillContentChangeHandler}></QuillEditor>


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
import {createEventDispatcher} from 'svelte';
import {createStory, updateStory} from "apis/stories.js";
import {successToast} from "lib/js/toast.js";
const dispatch = createEventDispatcher();

export let active = false;
export let edit = false;
export let story = "";
let default_story = "";
let quill_story_content = "";
let tags = [];

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