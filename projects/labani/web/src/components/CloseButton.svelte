<button class="delete" on:click={buttonClick}></button>
<script>
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
function buttonClick(){
    dispatch('tap');
}
</script>
<style>
.delete{
    -webkit-touch-callout: none;
    -webkit-user-select: none;
    -moz-user-select: none;
    user-select: none;
    background-color: #e91e63;
    border: none;
    border-radius: 9999px;
    cursor: pointer;
    pointer-events: auto;
    display: inline-block;
    flex-grow: 0;
    flex-shrink: 0;
    font-size: 0;
    width: 25px;
    height: 25px;
    max-height: 25px;
    max-width: 25px;
    min-height: 25px;
    min-width: 25px;
    outline: 0;
    position: relative;
    vertical-align: top;
}
.delete::before{
    height: 2px;
    width: 50%;
}
.delete::after{
    height: 50%;
    width: 2px;
}
.delete::after, .delete::before{
    background-color: #fff;
    content: "";
    display: block;
    left: 50%;
    position: absolute;
    top: 50%;
    transform: translateX(-50%) translateY(-50%) rotate(45deg);
    transform-origin: center center;
}
</style>