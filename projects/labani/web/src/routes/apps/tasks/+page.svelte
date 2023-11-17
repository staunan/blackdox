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
            <TaskList on:selected={onTaskItemSelected} />
        {:else if selectedLeftMenuItem === "Projects"}
            <ProjectList on:selected={onProjectItemSelected} />
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
import SvelteButton from 'components/SvelteButton.svelte';
import MenuItem from 'components/MenuItem.svelte';
import EditTaskModal from "./EditTaskModal.svelte";
import EditProjectModal from "./EditProjectModal.svelte";
import TaskDetailsModal from "./TaskDetailsModal.svelte";
import ProjectList from "./ProjectList.svelte";
import TaskList from "./TaskList.svelte";

let selectedLeftMenuItem = "Tasks";
let newTaskModalIsActive = false;
let newProjectModalIsActive = false;
let detailsTaskModalIsActive = false;

let selectedTask = null;
let selectedProject = null;

let editTask = false;
let editProject = false;
function openNewProjectModal(){
    newProjectModalIsActive = true;
}
function openNewTaskModal(){
    selectedTask = null;
    newTaskModalIsActive = true;
    editTask = false;
}
function handleLeftMenuItemSelected(menu_item_name){
    selectedLeftMenuItem = menu_item_name;
}
function onTaskItemSelected(event){
    selectedTask = event.detail;
    detailsTaskModalIsActive = true;
}
function onProjectItemSelected(event){
    selectedProject = event.detail;
}
function onNewTaskCreated(event){
    tasks = [event.detail, ...tasks];
    newTaskModalIsActive = false;
}
function onNewProjectCreated(event){
    projects = [event.detail, ...projects];
    newProjectModalIsActive = false;
}
function openTaskUpdateModal(){
    closeTaskDetailsModal();
    newTaskModalIsActive = true;
    editTask = true;
}
function closeNewTaskModal(){
    newTaskModalIsActive = false;
}
function closeNewProjectModal(){
    newProjectModalIsActive = false;
}
function closeTaskDetailsModal(){
    detailsTaskModalIsActive = false;
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
function onTaskRemoved(){
    tasks = tasks.filter(function(each){
       return selectedTask.internalId !== each.internalId
    });
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
.task_left_panel{
    flex-basis: 20%;
    padding: 20px;
    padding-left: 0;
}
.task_right_panel{
    flex-basis: 80%;
}
</style>