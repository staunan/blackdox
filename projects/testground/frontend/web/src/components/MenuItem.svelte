<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div 
    class="menu_item" 
    class:selected={selected} 
    class:clickable={clickable} 
    on:click={onClick} 
    style:background-color={background}
    style:color={color}
>
    <i class="fa {icon}"></i>
    <div>{label}</div>
</div>
<script>
import {createEventDispatcher} from 'svelte';
export let icon = "";
export let label = "";
export let selected = false;
export let clickable = true;
export let background = "#fff";
export let color = "black";
const dispatch = createEventDispatcher();
function onClick(event) {
    dispatch('tap');
}
</script>
<style>
.menu_item{
    width: 100%;
    height: 40px;
    display: flex;
    justify-content: flex-start;
    align-items: center;
    font-size: 16px;
    font-weight: normal;
    cursor: pointer;
    padding-left: 20px;
}
.menu_item i{
    margin-right: 15px;
}
.clickable.menu_item:hover{
    background-color: #ddd;
}
.clickable.selected{
    background-color: crimson !important; 
}
.clickable.selected *{
    color: #fff !important;
}
</style>