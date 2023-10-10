<button type="button" class="btn blue" on:click={onClick}><span>{ title }</span></button>
<script>
export let title = "Submit";
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
function onClick(event) {
    dispatch('tap');
}
</script>
<style>
.btn {
    position: relative;
    display: block;
    padding: 0;
    overflow: hidden;
    border-width: 0;
    outline: none;
    border-radius: 2px;
    transition: background-color .3s;
    cursor: pointer;
    font-size: 14px;
}
.btn > * {
    position: relative;
}
.btn span {
    display: block;
    padding: 12px 24px;
}
.btn:before {
    content: "";
    position: absolute;
    top: 50%;
    left: 50%;
    display: block;
    width: 0;
    padding-top: 0;
    border-radius: 100%;
    background-color: rgba(236, 240, 241, .3);
    -webkit-transform: translate(-50%, -50%);
    -moz-transform: translate(-50%, -50%);
    -ms-transform: translate(-50%, -50%);
    -o-transform: translate(-50%, -50%);
    transform: translate(-50%, -50%);
}
.btn:active:before {
    width: 120%;
    padding-top: 120%;
    transition: width .2s ease-out, padding-top .2s ease-out;
}
/* Colors */
.btn.blue{
    box-shadow: 0 1px 4px rgba(0, 0, 0, .6);
    background-color: #673ab7;
    color: #fff;
}
.btn.blue:hover, .btn.blue:focus {
    background-color: #3f51b5;
}
</style>