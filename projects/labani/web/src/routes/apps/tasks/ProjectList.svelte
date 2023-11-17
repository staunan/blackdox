<h3><MenuItem icon="fa-box" label="Projects" clickable={false} background="#795548" color="#fff"></MenuItem></h3>
{#each projects as project}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="project_item" on:click={openProjectDetailsModal(project)}>
        <div class="project_item_header">
            <div class="project_name">{ project.project_name }</div>
            <div class="project_details">{ project.project_details }</div>
        </div>
    </div>
{/each}
{#if projects.length === 0}
    <div class="empty_no_item">
        No Project Found
    </div>
{/if}
<script>
import { onMount } from 'svelte';
import {createEventDispatcher} from 'svelte';
import MenuItem from 'components/MenuItem.svelte';
import {getProjects} from "apis/tasks.js";
const dispatch = createEventDispatcher();
let projects = [];
let selectedProject = null;
onMount(() => {
    getAllProjects();
});
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
function openProjectDetailsModal(project){
    selectedProject = project;
    dispatch('selected', selectedProject);
}
</script>
<style>
.project_item{
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    padding: 30px;
    margin-bottom: 20px;
    border-radius: 10px;
    cursor: pointer;
    display: flex;
}
.empty_no_item{
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
.project_item_header{
    display: flex;
    flex-direction: column;
    flex: 1;
}
.project_name{
    font-size: 24px;
    font-weight: 600;
}
.project_details{
    font-size: 14px;
    font-weight: 400;
}
</style>