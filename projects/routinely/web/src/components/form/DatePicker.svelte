<div class="date_picker">
    {#if label }
        <FormLabel label={label}></FormLabel>
    {/if}
    <DateInput placeholder="Select a date" format="yyyy-MM-dd" bind:value={month_date} />
</div>
<script>
import FormLabel from 'components/form/FormLabel.svelte';
import { DateInput } from 'date-picker-svelte';
export let label = "";

let month_date = null;
</script>
<style>
:global(.date-time-field input) {
	font-family: monospace;
    height: 36px;
    font-size: 18px;
    border-radius: 0 !important;
    width: 100% !important;
}
:global(.picker) {
	font-family: monospace;
    font-size: 16px;
}
</style>