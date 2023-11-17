<!-- Create/Update Task Modal -->
<ContentModal 
    title={edit ? "Update Task" : "New Task"}
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body" class="modal_body">
        <Textbox label="Task Title" placeholder="Task Title" on:change={titleChangeHandler} value={title}></Textbox>
        <TextArea label="Task Details" placeholder="Details" on:change={detailsChangeHandler} value={details}></TextArea>
        <div class="form-container">
            <div class="form-label">Select Status</div>
            <div class="form-input">
                <Dropdown placeholder="-- Select Status --" items={all_task_status} currentitem={selectedStatus} on:change={statusChangeHandler} />
            </div>
        </div>
        <div class="form-container">
            <div class="form-label">Assign Project</div>
            <div class="form-input">
                <Dropdown placeholder="-- Choose Project --" items={all_projects} currentitem={selectedProject} on:change={projectChangeHandler} />
            </div>
        </div>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeModal}></SvelteButton>
        {#if edit}
            <SvelteButton color="blue" title="Update Task" on:tap={onUpdateTaskSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Task" on:tap={onNewTaskSubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>
<script>
import { onMount } from 'svelte';
import {createEventDispatcher} from 'svelte';
import Dropdown from 'components/Dropdown.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
import SvelteButton from 'components/SvelteButton.svelte';
import { STATUS } from "./const.js";
import {createTask, updateTask, getProjects} from "apis/tasks.js";
import {successToast} from "lib/js/toast.js";
const dispatch = createEventDispatcher();

let all_task_status = [];
let all_projects = [];

export let edit = false;
export let active = false;
export let task = null;


let title = "";
let details = "";
let selectedStatus = null;
let selectedProject = null;

$: {
    if(task){
        setTaskData(task);
    }else{
        setTaskData(null);
    }
}
onMount(() => {
    getAllStatus();
    getAllProjects();
    setTaskData(null);
});
function closeModal(){
    dispatch("close");
}
function titleChangeHandler(event){
    title = event.detail;
}
function detailsChangeHandler(event){
    details = event.detail;
}
function projectChangeHandler(event){
    selectedProject = event.detail;
}
function setTaskData(task_data){
    if(task_data){
        title = task_data.title;
        details = task_data.details;
        
        selectedStatus = all_task_status.filter((item)=>item.value === task.status)[0];
        selectedProject = all_projects.filter((item)=>item.value === task.project_id)[0];
    }else{
        title = "";
        details = "";
        selectedStatus = selectedStatus = all_task_status.filter((item)=>item.label === STATUS.Pending)[0];
        selectedProject = all_projects[0];
    }
}
function statusChangeHandler(event){
    selectedStatus = event.detail;
}
async function onNewTaskSubmit(event){
    let formData = {
        'title': title,
        'details': details,
        'project_id': selectedProject.value,
        'status': selectedStatus ? selectedStatus.label : STATUS.Pending
    };
    try{
        let response = await createTask(formData);
        dispatch('created', response.task);
        successToast(response.message);
        setTaskData(null);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateTaskSubmit(event){
    let formData = {
        'task_id': task.internalId,
        'project_id': selectedProject.value,
        'title': title,
        'details': details,
        'status': selectedStatus ? selectedStatus.label : STATUS.Pending
    };
    try{
        let response = await updateTask(formData);
        dispatch('updated', response.task);
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function getAllProjects(){
    try{
        let formData = {};
        let response = await getProjects(formData);
        all_projects = response.projects.map((project)=>{return {"label": project.project_name, "value": project.internalId}});
        if(task){
            setTaskData(task);
        }else{
            setTaskData(null);
        }
    }catch(error){
        console.log(error);
    }
}
async function getAllStatus(){
    all_task_status = Object.keys(STATUS).map((item)=>{return {"label": STATUS[item], "value": STATUS[item]}});
}
</script>
<style>
.modal_body{
    padding: 20px;
}
.form-container{
    margin-bottom: 30px;
}
.form-label{
    margin-bottom: 5px;
    font-weight: 500;
    font-size: 14px;
}
</style>