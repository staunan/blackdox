{#if active}
    <div class="form_modal">
        <div class="modal_overlay"></div>
        <div class="form_modal_window" >
            <div class="form_modal_header">
                <slot name="header" />
            </div>
            <div class="form_modal_body">
                <slot name="body" />
            </div>
            <div class="form_modal_footer">
                <slot name="footer" />
            </div>
        </div>
    </div> 
{/if}
<script>
import {createEventDispatcher} from 'svelte';
import { onMount } from "svelte";
export let active = false;
export let overlayclose = false;
const dispatch = createEventDispatcher();
onMount(()=>{
    document.addEventListener("click", function(event){
        if(event.target.closest(".modal_overlay")){
            if(overlayclose){
                dispatch('close');
            }
        }
    });
});
</script>
<style>
.form_modal{
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    width: 100%;
    height: 100vh;
}
.modal_overlay{
    background: #3f51b5c9;
    height: inherit;
    width: inherit;
}
.form_modal_window{
    position: absolute;
    top: 50%;
    left: 50%;
    width: 500px;
    height: 90%;
    transform: translate(-50%, -50%);
    display: flex;
    flex-direction: column;
    box-shadow: 0px 10px 15px -3px rgba(0,0,0,0.1);
    background-color: #fff;
}
.form_modal_header{
    height: 36px;
    width: 100%;
    background-color: red;
}
.form_modal_body{
    flex: 1;
    overflow-y: auto;
}
.form_modal_footer{
    height: 36px;
    width: 100%;
    background-color: red;
}
</style>