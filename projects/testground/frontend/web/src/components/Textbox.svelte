<div class="form_field">
    <div class="form_label">{ label }</div>
    <input type="text" placeholder={placeholder} value={value} on:keyup={onTextChange}>
</div>
<script>
export let label = "Input Label";
export let placeholder = "Input Placeholder...";
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
export let value = "";
function onTextChange(event){
    dispatch('change', event.target.value);
}
</script>
<style>
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