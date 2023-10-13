{#if active}
    <div class="form_modal">
        <div class="modal_overlay"></div>
        <div class="form_modal_window" >
            <div class="form_modal_header">
                <div class="modal_title">
                    <span>{ title }</span>
                </div>
                <div class="closeButton">
                    <CloseButton on:tap={closeModal}></CloseButton>
                </div>
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
import CloseButton from 'components/CloseButton.svelte';
import {createEventDispatcher} from 'svelte';
import { onMount } from "svelte";
export let active = false;
export let overlayclose = false;
export let title = "Modal Title";
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
function closeModal(){
    dispatch('close');
}
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
    height: 40px;
    width: 100%;
    background-color: #fff;
    box-shadow: rgb(38, 57, 77) 0px 20px 30px -10px;
    display: flex;
}
.modal_title{
    font-size: 16px;
    font-weight: bold;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 20px;
    flex: 1;
}
.closeButton{
    display: flex;
    justify-content: center;
    align-items: center;
    width: 60px;
    height: 40px;
}
.form_modal_body{
    flex: 1;
    overflow-y: auto;
}
.form_modal_body{
    padding: 20px;
}
.form_modal_footer{
    height: 60px;
    border-top: 1px solid #ccc;
    padding-left: 20px;
    padding-right: 20px;
    padding-top: 10px;
    padding-bottom: 10px;
}
</style>