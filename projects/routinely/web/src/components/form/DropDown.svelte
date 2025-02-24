<div class="dropdown_container">
    {#if label }
        <FormLabel label={label}></FormLabel>
    {/if}
    <div class="dropdown">
        <!-- svelte-ignore a11y-click-events-have-key-events -->
        <!-- svelte-ignore a11y-no-static-element-interactions -->
        <div class="dropdown_trigger" on:click={() => active = !active}>
            {#if selected_item}
            <div class="dropdown_label">{selected_item.label}</div>
            {:else}
            <div class="dropdown_no_selected_item">{placeholder}</div>
            {/if}
            <div class="dropdown_expand_icon" class:collapse={active}>
                <ArrowDown></ArrowDown>
            </div>
        </div>
        <div class="dropdown_content" class:show={active}>
            <div class="dropdown_list">
                {#if items.length > 0}
                    {#each items as item}
                    <!-- svelte-ignore a11y-click-events-have-key-events -->
                    <!-- svelte-ignore a11y-no-static-element-interactions -->
                    <div class="dropdown_listitem" on:click={() => handleDropdownItemClick(item)}>{item.label}</div>
                    {/each}
                {:else}
                <div class="dropdown_no_item" >No Data Found !</div>
                {/if}
            </div>        
        </div>
    </div>
    {#if hasError}
        <FormErrorMessage message={errorMessage}></FormErrorMessage>
    {/if}
</div>

<script>
import FormLabel from 'components/form/FormLabel.svelte';
import FormErrorMessage from 'components/form/FormErrorMessage.svelte';
import ArrowDown from 'components/svg/ArrowDown.svelte';

import {createEventDispatcher} from 'svelte';
import { onMount } from "svelte";
const dispatch = createEventDispatcher();

export let label = "";
export let items = [];
export let currentitem = null;
export let placeholder = "--Select Item--";
export let hasError = false;
export let errorMessage = "";

let active = false;
let selected_item = null;
$: {
    if(currentitem){
        selected_item = currentitem;
    }
}
onMount(()=>{
    document.addEventListener("click", function(event){
        if(!event.target.closest(".dropdown")){
            active = false;
        }
    });
});
function handleDropdownItemClick(item){
    dispatch('change', item);
    selected_item = item;
    active = !active;
}
</script>
<style>
.dropdown{
    position: relative;
    font-family: monospace;
}
.dropdown_trigger{
    height: 36px;
    font-size: 16px;
    padding: 5px;
    border: 1px solid #ccc;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    cursor: pointer;
    padding-left: 10px;
    background-color: #fff;
    box-sizing: border-box;
}
.dropdown_no_selected_item{
    height: 36px;
    width: 100%;
    font-size: 18px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 0;
    color: #757575;
}
.dropdown_content{
    position: absolute;
    top: 40px;
    left: 0;
    width: 100%;
    height: auto;
    max-height: 250px;
    overflow-y: scroll;
    border: 1px solid #ccc;
    display: none;
    z-index: 10001;
    background-color: #fff;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
}
.show{
    display: block !important;
}
.dropdown_list{
    display: flex;
    flex-direction: column;
    width: inherit;
    height: inherit;
}
.dropdown_listitem{
    height: 30px;
    border-bottom: 1px solid #ccc;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 20px;
    cursor: pointer;
    font-size: 18px;
}
.dropdown_listitem:hover{
    background-color: #ddd;
}
.dropdown_label{
    font-size: 18px;
    flex: 1;
}
.dropdown_expand_icon{
    display: flex;
    align-items: center;
    margin-left: 30px;
    transition: 300ms all;
}
.dropdown_expand_icon.collapse{
    transform: rotate(-180deg);
}
.dropdown_no_item{
    width: 100%;
    height: 50px;
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 13px;
    font-style: italic;
}
</style>