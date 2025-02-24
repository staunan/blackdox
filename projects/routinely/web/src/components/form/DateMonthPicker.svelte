<div class="datemonthpicker_container">
    {#if label }
        <FormLabel label={label}></FormLabel>
    {/if}
    <div class="datemonthpicker_input">
        <div class="month_input">
            <Dropdown placeholder="Month" items={Months} currentitem={selectedMonth} on:change={monthChangeHandler} />
        </div>
        <div class="day_input">
            <Dropdown placeholder="Day" items={Days} currentitem={selectedDay} on:change={dayChangeHandler} />
        </div>
    </div>
</div>
<script>
import Dropdown from 'components/form/Dropdown.svelte';
import FormLabel from 'components/form/FormLabel.svelte';
export let label = "";

let Months = [
	{"label": "January", "value": 1},
    {"label": "February", "value": 2},
    {"label": "March", "value": 3},
    {"label": "April", "value": 4},
    {"label": "May", "value": 5},
    {"label": "June", "value": 6},
    {"label": "July", "value": 7},
    {"label": "August", "value": 8},
    {"label": "September", "value": 9},
    {"label": "October", "value": 10},
    {"label": "November", "value": 11},
    {"label": "December", "value": 12}   
];
let Days = [];

let selectedMonth = null;
let selectedDay = null;

function monthChangeHandler(event){
    selectedMonth = event.detail;
    if(["January", "March", "May", "July", "August", "October", "December"].includes(selectedMonth.label)){
        Days = calculateDays(31);
    }else if(["April", "June", "September", "November"].includes(selectedMonth.label)){
        Days = calculateDays(30);
    }else if(["February"].includes(selectedMonth.label)){
        Days = calculateDays(29);
    }
}
function calculateDays(length){
    let arr = [];
    for(let i=1; i<=length; i++){
        arr.push({label: i, value: i});
    }
    return arr;
}
function dayChangeHandler(event){
    selectedDay = event.detail;
}
</script>
<style>
.datemonthpicker_input{
    display: flex;
    width: 100%;
}
.month_input, .day_input{
    flex: 1;
}
</style>