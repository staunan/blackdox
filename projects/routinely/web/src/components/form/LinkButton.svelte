<!-- svelte-ignore a11y_click_events_have_key_events -->
<!-- svelte-ignore a11y_no_static_element_interactions -->
<!-- svelte-ignore a11y_missing_attribute -->
<a class="link_button btn" on:click={onClick}> <span>{label}</span> </a>
<script>
export let label = "";
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
function onClick(event) {
    dispatch('tap');
}
</script>
<style>
.link_button{
    font-family: monospace;
    font-size: 14px;
    font-weight: bold;
}
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
    color: #2196f3;
}
.btn:hover{
    text-decoration: none;
}
.btn > * {
    position: relative;
}
.btn span {
    display: block;
    padding: 5px 5px;
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
</style>