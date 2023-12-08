<div class="tags">
    {#each allTags as tag}
        <div class="tag_item"></div>
    {/each}
    <div class="select_tag_dropdown">
        <div class="dropdown">
            <!-- svelte-ignore a11y-click-events-have-key-events -->
            <!-- svelte-ignore a11y-no-static-element-interactions -->
            <div class="dropdown_trigger" on:click={() => active = !active}>
                <div class="dropdown_label">Select Tag</div>
                <div class="dropdown_expand_icon"><i class="fa fa-arrow-down" aria-hidden="true"></i></div>
            </div>
            <div class="dropdown_content" class:show={active}>
                <div class="dropdown_list">
                    {#if allTags.length > 0}
                        {#each allTags as tag}
                        <!-- svelte-ignore a11y-click-events-have-key-events -->
                        <!-- svelte-ignore a11y-no-static-element-interactions -->
                        <div class="dropdown_listitem" on:click={() => handleDropdownItemClick(tag)}>{tag}</div>
                        {/each}
                    {:else}
                    <div class="dropdown_no_item" >No Data Found !</div>
                    {/if}
                </div>        
            </div>
        </div>
    </div>
</div>
<script>
import {createEventDispatcher} from 'svelte';
import { onMount } from "svelte";
const dispatch = createEventDispatcher();

export let selected = [];
let allTags = ["Tag 1", "Tag 2", "Tag 3"];
let selectedTags = [];
let active = false;

$: {
    if(selected){
        selectedTags = selected.map((item)=> item);
    }
}
onMount(()=>{
    document.addEventListener("click", function(event){
        if(!event.target.closest(".dropdown")){
            active = false;
        }
    });
});
function handleDropdownItemClick(tag){
    let match = false;
    for(let i=0; i<selectedTags.length; i++){
        if(selectedTags[i] === tag){
            match = true;
            break;
        }
    }
    if(match){
        selectedTags = selectedTags.filter((item)=>item === tag);
    }else{
        selectedTags = [...selectedTags, tag];
    }
    dispatch('change', selectedTags);
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
    margin-left: 30px;
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