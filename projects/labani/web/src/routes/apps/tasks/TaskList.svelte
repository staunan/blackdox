<h3><MenuItem icon="fa-list" label="Tasks" clickable={false} background="#009688" color="#fff"></MenuItem></h3>
<div class="filter_menu">
    <TaskFilter on:change={onFilterChange}></TaskFilter>
</div>
{#each tasks as task}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="task_item" on:click={handleOnItemClick(task)}>
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
<script>
import {createEventDispatcher} from 'svelte';
import TaskFilter from "./TaskFilter.svelte";
import MenuItem from 'components/MenuItem.svelte';
import {getTasks} from "apis/tasks.js";
import { STATUS } from "./const.js";
const dispatch = createEventDispatcher();
let tasks = [];
let currentFilter = {};
let selectedTask = null;

function onFilterChange(event){
    currentFilter = event.detail;
    getAllTasks();
}
async function getAllTasks(){
    try{
        let formData = {};
        if(currentFilter && currentFilter.status){
            formData.status = currentFilter.status;
        }else{
            formData.status = STATUS.Pending;
        }
        if(currentFilter && currentFilter.project){
            formData.project_id = currentFilter.project;
        }
        
        let response = await getTasks(formData);
        tasks = response.tasks;
    }catch(error){
        console.log(error);
    }
}
function handleOnItemClick(task){
    selectedTask = task;
    dispatch('selected', selectedTask);
}
</script>
<style>
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