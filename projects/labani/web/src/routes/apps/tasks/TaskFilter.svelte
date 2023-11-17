<div class="filter_menu">
    <div class="status_filter">
        <Dropdown items={all_task_status} currentitem={selectedStatus} on:change={filterStatusChangeHandler} />
    </div>
    <div class="project_filter">
        <Dropdown items={all_projects} currentitem={selectedProject} on:change={filterProjectChangeHandler} />
    </div>
</div>
<script>
import { onMount } from 'svelte';
import { STATUS } from "./const.js";
import {getProjects} from "apis/tasks.js";
import Dropdown from 'components/Dropdown.svelte';
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();

let all_task_status = [];
let all_projects = [];

let selectedStatus = null;
let selectedProject = null;
let filterData = {};

onMount(() => {
    getAllStatus();
    getAllProjects();
});
function filterStatusChangeHandler(event){
    selectedStatus = event.detail;
    filterChange();
}
function filterProjectChangeHandler(event){
    selectedProject = event.detail;
    filterChange();
}
async function getAllStatus(){
    all_task_status = Object.keys(STATUS).map((item)=>{return {"label": STATUS[item], "value": STATUS[item]}});
    selectedStatus = all_task_status[1]; // Default Status
    filterChange();
}
async function getAllProjects(){
    try{
        let formData = {};
        let response = await getProjects(formData);
        all_projects = response.projects.map((project)=>{return {"label": project.project_name, "value": project.internalId}});
        selectedProject = all_projects[0]; // Default Project
        filterChange();
    }catch(error){
        console.log(error);
    }
}
function filterChange(){
    if(selectedProject && selectedStatus){
        console.log("Filter Changed");
        filterData = {
            'status': selectedStatus.value,
            'project': selectedProject.value
        };
        dispatch('change', filterData);
    }
}
</script>
<style>
.filter_menu{
    display: flex;
    height: 50px;
    align-items: center;
    margin-top: 25px;
    margin-bottom: 25px;
}
.status_filter{
    margin-right: 20px;
}
</style>