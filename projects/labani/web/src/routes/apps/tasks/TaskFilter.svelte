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

let all_task_status = [];
let all_projects = [];

let selectedStatus = null;
let selectedProject = null;

onMount(() => {
    getAllStatus();
    getAllProjects();
});
function filterStatusChangeHandler(event){
    selectedStatus = event.detail;
}
function filterProjectChangeHandler(event){
    selectedProject = event.detail;
}
async function getAllStatus(){
    all_task_status = Object.keys(STATUS).map((item)=>{return {"label": STATUS[item], "value": STATUS[item]}});
}
async function getAllProjects(){
    try{
        let formData = {};
        let response = await getProjects(formData);
        all_projects = response.projects.map((project)=>{return {"label": project.project_name, "value": project.internalId}});
    }catch(error){
        console.log(error);
    }
}
</script>