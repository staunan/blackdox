<!-- Create/Update Project Modal -->
<ContentModal 
    title={edit ? "Update Project" : "New Project"}
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body" class="modal_body">
        <Textbox label="Project Name" placeholder="Name" on:change={projectNameChangeHandler} value={project_name}></Textbox>
        <TextArea label="Project Details" placeholder="Details" on:change={projectDetailsChangeHandler} value={project_details}></TextArea>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeModal}></SvelteButton>
        {#if edit}
            <SvelteButton color="blue" title="Update Project" on:tap={onUpdateProjectSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Project" on:tap={onNewProjectSubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>
<script>
import ContentModal from 'components/ContentModal.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import SvelteButton from 'components/SvelteButton.svelte';
import {createEventDispatcher} from 'svelte';
import {createProject, updateProject} from "apis/tasks.js";
import {successToast} from "lib/js/toast.js";

const dispatch = createEventDispatcher();
export let edit = false;
export let active = false;
export let project = null;

let project_name = "";
let project_details = "";

$: {
    if(project){
        console.log("Project Changed", project);
        setProjectData(project);
    }
}
function setProjectData(project_data){
    if(project_data){
        project_name = project_data.project_name;
        project_details = project_data.project_details;
    }else{
        project_details = "";
        project_details = "";
    }
}
function closeModal(){
    dispatch("close");
}
function projectNameChangeHandler(event){
    project_name = event.detail;
}
function projectDetailsChangeHandler(event){
    project_details = event.detail;
}
async function onUpdateProjectSubmit(event){
    let formData = {
        'project_id': selectedProject.internalId,
        'project_name': project_name,
        'project_details': project_details
    };
    try{
        let response = await updateProject(formData);
        dispatch('updated', response.project);
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onNewProjectSubmit(event){
    let formData = {
        'project_name': project_name,
        'project_details': project_details
    };
    try{
        let response = await createProject(formData);
        dispatch('created', response.project);
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
</script>
<style>
.modal_body{
    padding: 20px;
}
</style>