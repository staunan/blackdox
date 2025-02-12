<!-- svelte-ignore a11y-click-events-have-key-events -->
<!-- svelte-ignore a11y-no-static-element-interactions -->
<div class="delete_button_container" on:click={onClick} title={title}>
    <i class="fa fa-trash"></i>
</div>
<script>
import {createEventDispatcher} from 'svelte';
export let title = "";
const dispatch = createEventDispatcher();
function onClick(event) {
    dispatch('tap');
}
</script>
<style>
.delete_button_container{
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
    width: 40px;
    transition: 300ms all;
}
.delete_button_container i{
    color: #ccc;
}
.delete_button_container:hover{
    background-color: #ddd;
}
.delete_button_container:hover i{
    color: red;
}
</style>