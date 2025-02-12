<div class="form_field">
    {#if !value}
    <div class="search_icon">
        <i class="fa fa-search" aria-hidden="true"></i>
    </div>
    {:else}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="clear_search" on:click={clear_search}>
        <i class="fa fa-times" aria-hidden="true"></i>
    </div>
    {/if}
    
    <input type="text" placeholder={placeholder} value={value} on:keyup={onTextChange}>
</div>
<script>
export let placeholder = "Input Placeholder...";
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
export let value = "";
function onTextChange(event){
    value = event.target.value;
    dispatch('change', event.target.value);
}
function clear_search(){
    value = "";
    dispatch('change', "");
}
</script>
<style>
.clear_search{
    color: red;
    cursor: pointer;
}
.search_icon{
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
}
.clear_search{
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    right: 10px;
    top: 50%;
    transform: translateY(-50%);
}
.search_icon i{
    color: #3e8ed0;
}
.form_field{
    position: relative;
}
.form_field input{
    height: 36px;
    outline: none;
    border: 1px solid #b5b5b5;
    padding-left: 10px;
    box-shadow: inset 0 0.0625em 0.125em rgba(10,10,10,.05);
    border-radius: .375em;
    width: 100%;
    font-size: 16px;
    font-weight: 300px;
}
.form_field input:hover,.form_field input:focus{
    border-color: #3e8ed0;
}
</style>