<!-- Task Details Modal -->
<ContentModal 
    title="Task Details"
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body" class="modal_body">
        {#if task}
            <div class="task_details">
                <div class="task_title">{ task.title }</div>
                <div class="task_description">{ task.details }</div>
            </div>
            <div class="project_and_status">
                
                <div class="status_info">
                    <div class="status_label">
                        <i class="fa fa-list"></i>
                        <span>Project</span>
                    </div>
                    <div class="status_value">{ task.project_details[0].project_name }</div>
                </div>
                <div class="status_info">
                    <div class="status_label">Status</div>
                    <div class="status_value">{ task.status }</div>
                </div>
            </div>
        {/if}
    </div>
    <div slot="footer" class="footer_button_wrapper d-flex justify-content-space-between">
        <DeleteButton on:tap={removeTask} title="Delete Task"></DeleteButton>
        <SvelteButton color="blue" title="Update Task" on:tap={openEdit}></SvelteButton>
    </div>
</ContentModal>
<script>
import {createEventDispatcher} from 'svelte';
import ContentModal from 'components/ContentModal.svelte';
import DeleteButton from 'components/DeleteButton.svelte';
import SvelteButton from 'components/SvelteButton.svelte';
import {successToast} from "lib/js/toast.js";
import {deleteTask} from "apis/tasks.js";

export let task;
export let active = false;
const dispatch = createEventDispatcher();

function closeModal(){
    dispatch("close");
}
async function removeTask(){
    let formData = {
        'task_id': task.internalId
    };
    try{
        let response = await deleteTask(formData);
        successToast(response.message);
        dispatch("removed");
    }catch(error){
        console.log(error);
    }
}
function openEdit(){
    dispatch("edit");
}
</script>
<style>
.modal_body{
    padding: 20px;
}
.task_title{
    font-size: 18px;
    font-weight: bold;
}
.task_description{
    font-size: 14px;
    font-weight: 400;
}
.project_and_status{
    padding-top: 20px;
    padding-bottom: 20px;
    display: flex;
}
.project_and_status>div{
    flex: 1;
    display: flex;
    justify-content: flex-start;
}
.status_info{
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: flex-start;
}
.status_label{
    display: flex;
    align-items: center;
    font-size: 14px;
    font-weight: bold;
    justify-content: flex-start;
}
.status_label i{
    margin-right: 15px;
}
.status_value{
    font-size: 22px;
}
</style>