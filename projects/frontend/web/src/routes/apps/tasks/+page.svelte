<header>
	<div class="page_title">Task Manager</div>
	<div class="menu_items">
        
	</div>
</header>
<div class="task_list">
    <div class="task_left_panel">
        <MenuItem icon="fa-list" selected={selectedLeftMenuItem === "Tasks"} label="Tasks" on:tap={()=>{handleLeftMenuItemSelected("Tasks")}}></MenuItem>
        <MenuItem icon="fa-box" selected={selectedLeftMenuItem === "Projects"} label="Projects" on:tap={()=>{handleLeftMenuItemSelected("Projects")}}></MenuItem>
    </div>
    <div class="task_right_panel">
        {#if selectedLeftMenuItem === "Tasks"}
            <TaskList filter={currentFilter} />
        {:else if selectedLeftMenuItem === "Projects"}
            <ProjectList on:selected={handleOnProjectSelect} />
        {/if}
    </div>
</div>
<script>
import MenuItem from 'components/MenuItem.svelte';
import ProjectList from "./ProjectList.svelte";
import TaskList from "./TaskList.svelte";

let selectedLeftMenuItem = "Tasks";
let selectedProject = null;
let currentFilter = {};
function handleLeftMenuItemSelected(menu_item_name){
    selectedLeftMenuItem = menu_item_name;
}
function handleOnProjectSelect(event){
    selectedProject = event.detail;
    currentFilter = {"project": selectedProject.internalId};
    selectedLeftMenuItem = "Tasks";
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