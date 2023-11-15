<div class="dropdown">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <!-- svelte-ignore a11y-no-static-element-interactions -->
    <div class="dropdown_trigger" on:click={() => active = !active}>
        {#if selected_item}
        <div class="dropdown_label">{selected_item.label}</div>
        {:else}
        <div class="dropdown_no_selected_item">--Select Item--</div>
        {/if}
        <div class="dropdown_expand_icon"><i class="fa fa-arrow-down" aria-hidden="true"></i></div>
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
<script>
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
export let items = [];
export let currentitem = null;
let active = false;
let selected_item = null;
$: {
    if(currentitem){
        selected_item = currentitem;
    }
}
function handleDropdownItemClick(item){
    dispatch('change', item);
    selected_item = item;
    active = !active;
}
</script>
<style>
.dropdown{
    position: relative;
}
.dropdown_trigger{
    height: 40px;
    font-size: 16px;
    padding: 10px;
    border: 1px solid #ccc;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    cursor: pointer;
}
.dropdown_no_selected_item{
    height: 40px;
    width: 100%;
    font-size: 14px;
    font-style: italic;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    padding-left: 20px;
}
.dropdown_content{
    position: absolute;
    top: 40px;
    left: 0;
    width: 100%;
    min-height: 50px;
    height: 150px;
    overflow-y: scroll;
    border: 1px solid #ccc;
    display: none;
    z-index: 10001;
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
}
.dropdown_listitem:hover{
    background-color: #ddd;
}
.dropdown_label{
    flex: 1;
}
.dropdown_expand_icon{
    display: flex;
    align-items: center;
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