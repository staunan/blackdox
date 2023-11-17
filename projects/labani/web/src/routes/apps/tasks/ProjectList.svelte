<div class="page_title">
    <div class="page_name">
        <i class="fa fa-box"></i>
        <span>Projects</span>
    </div>
    <div class="create_task">
        <SvelteButton title="New Project" on:tap={openNewProjectModal}></SvelteButton>
    </div>
</div>
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
<div class="modals">
    <EditProjectModal 
        task={selectedProject}
        active={newProjectModalIsActive} 
        edit={editProject} 
        on:close={closeNewProjectModal}
        on:created={onNewProjectCreated}
        on:updated={onProjectUpdated}
    ></EditProjectModal>
</div>
<script>
import { onMount } from 'svelte';
import {createEventDispatcher} from 'svelte';
import {getProjects} from "apis/tasks.js"
import SvelteButton from 'components/SvelteButton.svelte';
import EditProjectModal from "./EditProjectModal.svelte";
;
const dispatch = createEventDispatcher();
let projects = [];
let selectedProject = null;
let newProjectModalIsActive = false;
let editProject = false;
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
function openNewProjectModal(){
    newProjectModalIsActive = true;
}
function openProjectDetailsModal(project){
    selectedProject = project;
    dispatch('selected', selectedProject);
}
function closeNewProjectModal(){
    newProjectModalIsActive = false;
}
function onNewProjectCreated(event){
    projects = [event.detail, ...projects];
    newProjectModalIsActive = false;
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
.page_title{
    display: flex;
    padding-top: 20px;
    padding-bottom: 20px;
}
.page_name{
    display: flex;
    align-items: center;
    justify-content: flex-start;
    font-size: 24px;
    font-weight: bold;
}
.page_name i{
    margin-right: 15px;
}
.create_task{
    flex: 1;
    display: flex;
    justify-content: flex-end;
    align-items: center;
}
</style>