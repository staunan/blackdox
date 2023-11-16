<!-- Task Details Modal -->
<ContentModal 
    title={task ? task.title : '' }
    active={active} 
    overlayclose={true}
    on:close={closeModal}
>
    <div slot="body" class="modal_body">
        {#if task}
            <div class="text_info">
                <div class="info_label">Details</div>
                <div class="info_value">{ task.details }</div>
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
</style>