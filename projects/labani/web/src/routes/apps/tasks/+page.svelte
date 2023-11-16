<header>
	<div class="page_title">Task Manager</div>
	<div class="menu_items">
        <div class="menu_item">
            <SvelteButton title="New Project" on:tap={openNewProjectModal}></SvelteButton>
        </div>
        <div class="menu_item">
            <SvelteButton title="New Task" on:tap={openNewTaskModal}></SvelteButton>
        </div>
	</div>
</header>
<div class="task_list">
    <div class="task_left_panel">
        <MenuItem icon="fa-list" selected={selectedLeftMenuItem === "Tasks"} label="Tasks" on:tap={()=>{handleLeftMenuItemSelected("Tasks")}}></MenuItem>
        <MenuItem icon="fa-box" selected={selectedLeftMenuItem === "Projects"} label="Projects" on:tap={()=>{handleLeftMenuItemSelected("Projects")}}></MenuItem>
    </div>
    <div class="task_right_panel">
        {#if selectedLeftMenuItem === "Tasks"}
            <h3><MenuItem icon="fa-list" label="Tasks" clickable={false} background="#009688" color="#fff"></MenuItem></h3>
            <div class="filter_menu">
                <div class="status_filter">
                    <Dropdown items={all_task_status} currentitem={current_filter_status} on:change={filterStatusChangeHandler} />
                </div>
            </div>
            {#each tasks as task}
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div class="task_item" on:click={openTaskDetailsModal(task)}>
                    <div class="task_item_header">
                        <span class="task_title">{ task.title }</span>
                    </div>
                    <div class="task_item_status">
                        <span class="status_text">{ task.status }</span>
                    </div>
                </div>
            {/each}
            {#if tasks.length === 0}
                <div class="task_no_item">
                    No Tasks Found
                </div>
            {/if}
        {:else if selectedLeftMenuItem === "Projects"}
            <h3><MenuItem icon="fa-list" label="Projects" clickable={false} background="#795548" color="#fff"></MenuItem></h3>
        {/if}
    </div>
</div>

<!-- Task Details Modal -->
<ContentModal 
    title={selectedTask ? selectedTask.title : '' }
    active={detailsTaskModalIsActive} 
    overlayclose={true}
    on:close={closeTaskDetailsModal}
>
    <div slot="body" class="modal_body">
        {#if selectedTask}
            <div class="text_info">
                <div class="info_label">Details</div>
                <div class="info_value">{ selectedTask.details }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <DeleteButton on:tap={removeTask} title="Delete Task"></DeleteButton>
        <SvelteButton color="blue" title="Update Task" on:tap={openUpdateTaskModal}></SvelteButton>
    </div>
</ContentModal>

<!-- Create/Update Task Modal -->
<ContentModal 
    title={editTask ? "Update Task" : "New Task"}
    active={newTaskModalIsActive} 
    overlayclose={true}
    on:close={closeNewTaskModal}
>
    <div slot="body" class="modal_body">
        <Textbox label="Task Title" placeholder="Task Title" on:change={titleChangeHandler} value={title}></Textbox>
        <TextArea label="Task Details" placeholder="Details" on:change={detailsChangeHandler} value={details}></TextArea>
        <Dropdown items={all_task_status} currentitem={current_status} on:change={statusChangeHandler} />
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeNewTaskModal}></SvelteButton>
        {#if editTask}
            <SvelteButton color="blue" title="Update Task" on:tap={onUpdateTaskSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Task" on:tap={onNewTaskSubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>

<!-- Create/Update Project Modal -->
<ContentModal 
    title={editProject ? "Update Project" : "New Project"}
    active={newProjectModalIsActive} 
    overlayclose={true}
    on:close={closeNewProjectModal}
>
    <div slot="body" class="modal_body">
        <Textbox label="Project Name" placeholder="Name" on:change={projectNameChangeHandler} value={project_name}></Textbox>
        <TextArea label="Project Details" placeholder="Details" on:change={projectDetailsChangeHandler} value={project_details}></TextArea>
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Cancel" on:tap={closeNewProjectModal}></SvelteButton>
        {#if editProject}
            <SvelteButton color="blue" title="Update Project" on:tap={onUpdateProjectSubmit}></SvelteButton>
        {:else}
            <SvelteButton color="blue" title="Create Project" on:tap={onNewProjectSubmit}></SvelteButton>
        {/if}
    </div>
</ContentModal>

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import { STATUS } from "./const.js";
import {createTask, createProject, updateTask, updateProject, getTasks, deleteTask} from "apis/tasks.js";
import SvelteButton from 'components/SvelteButton.svelte';
import Dropdown from 'components/Dropdown.svelte';
import DeleteButton from 'components/DeleteButton.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
import MenuItem from 'components/MenuItem.svelte';


let newTaskModalIsActive = false;
let newProjectModalIsActive = false;
let detailsTaskModalIsActive = false;

let title = "";
let details = "";
let current_status = null;

let projects = [];
let project_name = "";
let project_details = "";
let editProject = false;
let selectedProject = null;

let tasks = [];
let selectedTask = null;
let editTask = false;

let selectedLeftMenuItem = "Tasks";

let all_task_status = Object.keys(STATUS).map((item)=>{return {"label": STATUS[item], "value": STATUS[item]}});
let current_filter_status = all_task_status[1];

onMount(() => {
    getAllTasks();
});
function handleLeftMenuItemSelected(menu_item_name){
    console.log(menu_item_name);
    selectedLeftMenuItem = menu_item_name;
}
function openNewTaskModal(){
    newTaskModalIsActive = true;
    editTask = false;
    setTaskData(null);
}
function openNewProjectModal(){
    newProjectModalIsActive = true;
}
function closeNewTaskModal(){
    newTaskModalIsActive = false;
}
function closeNewProjectModal(){
    newProjectModalIsActive = false;
}
function openTaskDetailsModal(task){
    selectedTask = task;
    detailsTaskModalIsActive = true;
}
function closeTaskDetailsModal(){
    detailsTaskModalIsActive = false;
}
function titleChangeHandler(event){
    title = event.detail;
}
function detailsChangeHandler(event){
    details = event.detail;
}
function projectNameChangeHandler(event){
    project_name = event.detail;
}
function projectDetailsChangeHandler(event){
    project_details = event.detail;
}
function statusChangeHandler(event){
    current_status = event.detail;
}
function filterStatusChangeHandler(event){
    current_filter_status = event.detail;
    getAllTasks();
    console.log(current_filter_status);
}
function openUpdateTaskModal(){
    editTask = true;
    newTaskModalIsActive = true;
    detailsTaskModalIsActive = false;
    setTaskData(selectedTask);
}
function setTaskData(task){
    if(task){
        title = task.title;
        details = task.details;
        current_status = all_task_status.filter((item)=>item.label === task.status)[0];
    }else{
        title = "";
        details = "";
        current_status = current_status = all_task_status.filter((item)=>item.label === STATUS.Pending)[0];
    }
}
async function getAllTasks(){
    try{
        let formData = {};
        if(current_filter_status){
            formData.status = current_filter_status.label;
        }else{
            formData.status = STATUS.Pending;
        }
        let response = await getTasks(formData);
        tasks = response.tasks;
    }catch(error){
        console.log(error);
    }
}
async function onNewTaskSubmit(event){
    let formData = {
        'title': title,
        'details': details,
        'status': current_status ? current_status.label : STATUS.Pending
    };
    try{
        let response = await createTask(formData);
        closeNewTaskModal();
        tasks = [response.task, ...tasks];
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateTaskSubmit(event){
    let formData = {
        'task_id': selectedTask.internalId,
        'title': title,
        'details': details,
        'status': current_status ? current_status.label : STATUS.Pending
    };
    try{
        let response = await updateTask(formData);
        closeNewTaskModal();
        tasks = tasks.map(function(task){
            if(selectedTask.internalId === task.internalId){
                return response.task;
            }else{
                return task;
            }
        });
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
        closeNewProjectModal();
        projects = [response.project, ...projects];
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function onUpdateProjectSubmit(event){
    let formData = {
        'project_id': selectedProject.internalId,
        'project_name': project_name,
        'project_details': project_details
    };
    try{
        let response = await updateProject(formData);
        closeNewProjectModal();
        projects = projects.map(function(project){
            if(selectedProject.internalId === project.internalId){
                return response.project;
            }else{
                return project;
            }
        });
        successToast(response.message);
    }catch(error){
        console.log(error);
    }
}
async function removeTask(){
    let formData = {
        'task_id': selectedTask.internalId
    };
    try{
        let response = await deleteTask(formData);
        successToast(response.message);
        getTasks();
        closeTaskDetailsModal();
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
.menu_item{
    margin-left: 20px;
}
.task_list{
    padding: 50px;
    display: flex;
}
.task_item{
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    padding: 30px;
    margin-bottom: 20px;
    border-radius: 10px;
    cursor: pointer;
    display: flex;
}
.task_no_item{
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
.task_title{
    font-size: 24px;
    font-weight: 400;
}
.task_item_header{
    display: flex;
    flex-direction: column;
    flex: 1;
}
.task_item_status{
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 20px;
    font-weight: bold;
}
.modal_body{
    padding: 20px;
}
.filter_menu{
    display: flex;
    height: 50px;
    align-items: center;
    margin-top: 25px;
    margin-bottom: 25px;
}
.task_left_panel{
    flex-basis: 20%;
    padding: 20px;
    padding-left: 0;
}
.task_right_panel{
    flex-basis: 80%;
}
</style>