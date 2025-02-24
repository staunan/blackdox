<div class="form_field">
    <FormLabel label={label}></FormLabel>
    <textarea rows="5" placeholder={placeholder} value={value} on:keyup={onTextChange}></textarea>
    {#if hasError}
        <FormErrorMessage message={errorMessage}></FormErrorMessage>
    {/if}
</div>
<script>
import FormLabel from 'components/form/FormLabel.svelte';
import FormErrorMessage from 'components/form/FormErrorMessage.svelte';
export let label = "Input Label";
export let placeholder = "Input Placeholder...";
export let hasError = false;
export let errorMessage = "";
import {createEventDispatcher} from 'svelte';
const dispatch = createEventDispatcher();
export let value = "";
function onTextChange(event){
    dispatch('change', event.target.value);
}
</script>
<style>
.form_field textarea{
    outline: none;
    border: 1px solid #b5b5b5;
    padding-left: 10px;
    box-shadow: inset 0 0.0625em 0.125em rgba(10,10,10,.05);
    border-radius: .375em;
    width: 100%;
    font-size: 18px;
    font-weight: 300px;
    box-sizing: border-box;
    font-family: monospace;
    border-radius: 0;
}
.form_field textarea:hover,.form_field textarea:focus{
    border-color: #3e8ed0;
}
</style>