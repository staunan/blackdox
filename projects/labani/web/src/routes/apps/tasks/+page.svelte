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
                <TaskFilter on:change={onFilterChange}></TaskFilter>
            </div>
            {#each tasks as task}
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div class="task_item" on:click={openTaskDetailsModal(task)}>
                    <div class="task_item_header">
                        <div class="task_title">{ task.title }</div>
                        {#if task.project_details && task.project_details.length === 1}
                        <div class="task_project_name">
                            <i class="fa fa-box"></i>
                            <span>{ task.project_details[0].project_name }</span>
                        </div>
                        {/if}
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
            {#each projects as project}
                <!-- svelte-ignore a11y-click-events-have-key-events -->
                <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div class="task_item" on:click={openProjectDetailsModal(project)}>
                    <div class="task_item_header">
                        <div class="project_name">{ project.project_name }</div>
                        <div class="project_details">{ project.project_details }</div>
                    </div>
                </div>
            {/each}
            {#if tasks.length === 0}
                <div class="task_no_item">
                    No Tasks Found
                </div>
            {/if}
        {/if}
    </div>
</div>

<div class="modals">
    <EditTaskModal 
        task={selectedTask} 
        active={newTaskModalIsActive} 
        edit={editTask} 
        on:close={closeNewTaskModal}
        on:created={onNewTaskCreated}
        on:updated={onTaskUpdated}
    ></EditTaskModal>

    <EditProjectModal 
        task={selectedProject}
        active={newProjectModalIsActive} 
        edit={editProject} 
        on:close={closeNewProjectModal}
        on:created={onNewProjectCreated}
        on:updated={onProjectUpdated}
    ></EditProjectModal>

    <TaskDetailsModal 
        task={selectedTask}
        active={detailsTaskModalIsActive} 
        on:close={closeTaskDetailsModal}
        on:edit={openTaskUpdateModal}
        on:removed={onTaskRemoved}
    ></TaskDetailsModal>
</div>

<script>
import { onMount } from 'svelte';
import {getTasks, getProjects} from "apis/tasks.js";
import SvelteButton from 'components/SvelteButton.svelte';
import MenuItem from 'components/MenuItem.svelte';
import EditTaskModal from "./EditTaskModal.svelte";
import EditProjectModal from "./EditProjectModal.svelte";
import TaskDetailsModal from "./TaskDetailsModal.svelte";
import TaskFilter from "./TaskFilter.svelte";

let selectedLeftMenuItem = "Tasks";
let newTaskModalIsActive = false;
let newProjectModalIsActive = false;
let detailsTaskModalIsActive = false;

let selectedTask = null;
let selectedProject = null;

let editTask = false;
let editProject = false;

let tasks = [];
let projects = [];
let currentFilter = {};

onMount(() => {
    getAllTasks();
    getAllProjects();
});
function handleLeftMenuItemSelected(menu_item_name){
    console.log(menu_item_name);
    selectedLeftMenuItem = menu_item_name;
}
function openNewTaskModal(){
    selectedTask = null;
    newTaskModalIsActive = true;
    editTask = false;
}
function openTaskUpdateModal(){
    closeTaskDetailsModal();
    newTaskModalIsActive = true;
    editTask = true;
}
function closeNewTaskModal(){
    newTaskModalIsActive = false;
}
function openNewProjectModal(){
    newProjectModalIsActive = true;
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
function openProjectDetailsModal(project){
    selectedProject = project;
}
function onFilterChange(event){
    currentFilter = event.detail;
}
function onNewTaskCreated(event){
    tasks = [event.detail, ...tasks];
    newTaskModalIsActive = false;
}
function onNewProjectCreated(event){
    projects = [event.detail, ...projects];
    newProjectModalIsActive = false;
}
function onTaskUpdated(event){
    let task = event.detail;
    tasks = tasks.map(function(each){
        if(task.internalId === each.internalId){
            return task;
        }else{
            return each;
        }
    });
}
function onTaskRemoved(){
    tasks = tasks.filter(function(each){
       return selectedTask.internalId !== each.internalId
    });
}
function onProjectUpdated(event){
    let project = event.detail;
    projects = projects.map(function(each){
        if(project.internalId === each.internalId){
            return project;
        }else{
            return each;
        }
    });
}
async function getAllTasks(){
    try{
        let formData = {};
        if(current_filter_status){
            formData.status = current_filter_status.label;
        }else{
            formData.status = STATUS.Pending;
        }
        if(selectedProject){
            formData.project_id = selectedProject.value;
        }
        
        let response = await getTasks(formData);
        tasks = response.tasks;
    }catch(error){
        console.log(error);
    }
}
async function getAllProjects(){
    try{
        let formData = {};
        let response = await getProjects(formData);
        projects = response.projects;
        console.log(projects);
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
.project_name{
    font-size: 24px;
    font-weight: 600;
}
.project_details{
    font-size: 14px;
    font-weight: 400;
}
.task_project_name{
    margin-top: 10px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
}
.task_project_name i{
    margin-right: 15px;
}
</style>