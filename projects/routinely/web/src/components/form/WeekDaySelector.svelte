<div class="days_selector">
    {#if label }
        <FormLabel label={label}></FormLabel>
    {/if}
    <div class="days">
        {#if days.length > 0}
            {#each days as day}
                <!-- svelte-ignore a11y_click_events_have_key_events -->
                 <!-- svelte-ignore a11y-no-static-element-interactions -->
                <div class="day" class:selected={selectedDay && day.id == selectedDay.id } on:click={() => dayClickedHandler(day)} title={day.title}>{day.short}</div>
            {/each}
        {/if}
    </div>
</div>
<script>
import FormLabel from 'components/form/FormLabel.svelte';
import {createEventDispatcher} from 'svelte';

export let label = "";
export let value = "";

const dispatch = createEventDispatcher();
let selectedDay = null;
let days = [
    {id: 1, short: "Sun", title: "Sunday"},
    {id: 2, short: "Mon", title: "Monday"},
    {id: 3, short: "Tue", title: "Tuesday"},
    {id: 4, short: "Wed", title: "Wednesday"},
    {id: 5, short: "Thu", title: "Thursday"},
    {id: 6, short: "Fri", title: "Friday"},
    {id: 7, short: "Sat", title: "Saturday"},
];
$: {
    if(value){
        selectedDay = days.filter((day)=>day.short == value)[0];
    }else{
        selectedDay = null;
    }
}
function dayClickedHandler(day){
    if(selectedDay && day.id === selectedDay.id){
        selectedDay = null;
        dispatch('change', null);
    }else{
        selectedDay = day;
        dispatch('change', selectedDay.short);
    }
}
</script>
<style>
.days{
    display: flex;
    align-items: center;
    justify-content: left;
}
.day{
    display: flex;
    justify-content: center;
    align-items: center;
    width: 40px;
    height: 40px;
    cursor: pointer;
    border-radius: 50%;
    border: 1px solid #ccc;
    background-color: #fff;
    margin-right: 10px;
    margin-bottom: 10px;
    font-family: monospace;
    font-size: 14px;
    font-weight: 500;
}
.day:hover{
    background-color: #ccc;
}
.day.selected{
    color: #fff;
    box-shadow: rgba(0, 0, 0, 0.35) 0px 5px 15px;
    background: #0F2027;  /* fallback for old browsers */
    background: -webkit-linear-gradient(to right, #2C5364, #203A43, #0F2027);  /* Chrome 10-25, Safari 5.1-6 */
    background: linear-gradient(to right, #2C5364, #203A43, #0F2027); /* W3C, IE 10+/ Edge, Firefox 16+, Chrome 26+, Opera 12+, Safari 7+ */
}
</style>