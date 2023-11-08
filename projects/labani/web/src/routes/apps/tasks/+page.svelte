<header>
	<div class="page_title">Task Manager</div>
	<div class="menu_items">
		<SvelteButton title="New Task" on:tap={openNewTaskModal}></SvelteButton>
	</div>
</header>
<div class="task_list">
    {#each tasks as task}
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="task_item" on:click={openTaskDetailsModal(task)}>
            <div class="task_item_header">
                <span class="task_title">{ task.title }</span>
            </div>
        </div>
    {/each}
    {#if tasks.length === 0}
        <div class="task_no_item">
            No Tasks Found
        </div>
    {/if}
</div>

<!-- Task Details Modal -->
<ContentModal 
    title={selectedTask ? selectedTask.title : '' }
    active={detailsTaskModalIsActive} 
    overlayclose={true}
    on:close={closeTaskDetailsModal}
>
    <div slot="body">
        {#if selectedTask}
            <div class="text_info">
                <div class="info_label">Details</div>
                <div class="info_value">{ selectedTask.details }</div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <SvelteButton color="red" title="Delete Task" on:tap={removeTask}></SvelteButton>
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
    <div slot="body">
        <Textbox label="Task Title" placeholder="Task Title" on:change={titleChangeHandler} value={title}></Textbox>
        <TextArea label="Task Details" placeholder="Details" on:change={detailsChangeHandler} value={details}></TextArea>
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

<script>
import { onMount } from 'svelte';
import {successToast} from "lib/js/toast.js";
import {createTask, updateTask, getTasks, deleteTask} from "apis/tasks.js";
import SvelteButton from 'components/SvelteButton.svelte';
import Textbox from 'components/Textbox.svelte';
import TextArea from 'components/TextArea.svelte';
import ContentModal from 'components/ContentModal.svelte';
let newTaskModalIsActive = false;
let detailsTaskModalIsActive = false;

let title = "";
let details = "";

let tasks = [];
let selectedTask = null;
let editTask = false;

onMount(() => {
    getAllTasks();
});

function openNewTaskModal(){
    newTaskModalIsActive = true;
    editTask = false;
    setTaskData(null);
}
function closeNewTaskModal(){
    newTaskModalIsActive = false;
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
    }else{
        title = "";
        details = "";
    }
}
async function getAllTasks(){
    try{
        let response = await getTasks();
        tasks = response.tasks;
    }catch(error){
        console.log(error);
    }
}
async function onNewTaskSubmit(event){
    let formData = {
        'title': title,
        'details': details
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
        'details': details
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
async function removeTask(){
    let formData = {
        'task_id': selectedTask.internalId
    };
    try{
        let response = await deleteTask(formData);
        successToast(response.message);
        getAllTasks();
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
.task_list{
    padding: 50px;
    padding-top: 110px;
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
</style>