<div class="timepicker_container">
    {#if label }
        <FormLabel label={label}></FormLabel>
    {/if}
    <div class="timepicker_input">
        <div class="hour_input">
            <Dropdown placeholder="Hours" items={Hours} currentitem={selectedHour} on:change={hourChangeHandler} />
        </div>
        <div class="minute_input">
            <Dropdown placeholder="Minute" items={Minutes} currentitem={selectedMinute} on:change={minuteChangeHandler} />
        </div>
    </div>
</div>
<script>
import Dropdown from 'components/form/Dropdown.svelte';
import FormLabel from 'components/form/FormLabel.svelte';
export let label = "";

let Hours = [
	{"label": "01", "value": 1},
    {"label": "02", "value": 2},
    {"label": "03", "value": 3},
    {"label": "04", "value": 4},
    {"label": "05", "value": 5},
    {"label": "06", "value": 6},
    {"label": "07", "value": 7},
    {"label": "08", "value": 8},
    {"label": "09", "value": 9},
    {"label": "10", "value": 10},
    {"label": "11", "value": 11},
    {"label": "12", "value": 12},
];
let Minutes = [
	{"label": "00", "value": 0},
    {"label": "05", "value": 5},
    {"label": "10", "value": 10},
    {"label": "15", "value": 15},
    {"label": "20", "value": 20},
    {"label": "25", "value": 25},
    {"label": "30", "value": 30},
    {"label": "35", "value": 35},
    {"label": "40", "value": 40},
    {"label": "45", "value": 45},
    {"label": "50", "value": 50},
    {"label": "55", "value": 55},
];

let selectedHour = null;
let selectedMinute = null;

function hourChangeHandler(event){
    selectedHour = event.detail;
}
function minuteChangeHandler(event){
    selectedMinute = event.detail;
}
</script>
<style>
.timepicker_input{
    display: flex;
    width: 100%;
}
.hour_input, .minute_input{
    flex: 1;
}
</style>