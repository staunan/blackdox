<Card>
    <div class="card_body">
        <h1 class="center mb10 form_heading">Create Routine</h1>
        <TextBox label={routine_title_label} placeholder="Write a title for your routine..." on:change={titleChangedHandler} value={routine_title} hasError={routineTitleHasError} errorMessage={routineTitleErrorMessage}></TextBox>
        {#if advanceSettings}
        <div class="advanceSettings">
            <div class="form_gap"></div>
            <TextArea label="Routine Description" placeholder="Write a description for your routine (Optional)..." on:change={routineDetailsChangeHandler} value={routine_details} hasError={routineDetailsHasError} errorMessage={routineDetailsErrorMessage}></TextArea>
            <div class="form_gap"></div>
            <Dropdown label="Routine Mode" placeholder="Select a execution mode of this routine " items={all_routine_modes} currentitem={selected_routine_mode} on:change={routineModeChangeHandler} hasError={routineModeHasError} errorMessage={routineModeErrorMessage} />
            {#if selected_routine_mode && selected_routine_mode.value === "Daily"}
                <div class="form_gap"></div>
                <div class="form_row">
                    <div class="form_column">
                        <DaysSelector selectedItems={selected_days} label="Select Days"></DaysSelector>
                    </div>
                    <div class="form_column">
                        <TimePicker label="Choose a time of the day" selectedTime={selectedTime}></TimePicker>
                    </div>
                </div>
            {:else if selected_routine_mode && selected_routine_mode.value === "Weekly"}
                <div class="form_gap"></div>
                <WeekDaySelector selectedItem={selected_week_day} label="Choose a day on which you want execute your routine"></WeekDaySelector>
            {:else if selected_routine_mode && selected_routine_mode.value === "Monthly"}
                <div class="form_gap"></div>
                <DayPicker label="Choose a Day"></DayPicker>
            {:else if selected_routine_mode && selected_routine_mode.value === "Yearly"}
                <div class="form_gap"></div>
                <DateMonthPicker label="Choose a date (Year-Month-Day)"></DateMonthPicker>
            {/if}
            <div class="advance_settings_hide">
                <LinkButton label="Hide Advance Settings" on:tap={advanceSettings = !advanceSettings}></LinkButton>
            </div>
        </div>
        {:else}
            <div class="advance_settings_trigger">
                <LinkButton label="Show Advance Settings" on:tap={advanceSettings = !advanceSettings}></LinkButton>
                <div class="advance_settings_trigger_icon">
                    <ArrowDown></ArrowDown>
                </div>
            </div>
        {/if}
        
        <div class="center mt10 submit_button_container">
            <SubmitButton title="Create Routine" on:tap={createRoutineHandler}></SubmitButton>
        </div>
    </div>
</Card>
<script>
import Card from 'components/Card.svelte';
import TextBox from 'components/form/TextBox.svelte';
import TextArea from 'components/form/TextArea.svelte';
import Dropdown from 'components/form/Dropdown.svelte';
import DaysSelector from 'components/form/DaysSelector.svelte';
import WeekDaySelector from 'components/form/WeekDaySelector.svelte';
import DayPicker from 'components/form/DayPicker.svelte';
import DateMonthPicker from 'components/form/DateMonthPicker.svelte';
import TimePicker from 'components/form/TimePicker.svelte';
import SubmitButton from 'components/form/SubmitButton.svelte';
import LinkButton from 'components/form/LinkButton.svelte';
import ArrowDown from 'components/svg/ArrowDown.svelte';

import { createRoutine } from 'apis/apis.js';

let advanceSettings = false;
let routine_title_label = "";

let routine_title = "";
let routineTitleHasError = false;
let routineTitleErrorMessage = "";
let routine_details = "";
let routineDetailsHasError = false;
let routineDetailsErrorMessage = "";
let all_routine_modes = [
    {"label": "Daily", "value": "Daily"},
    {"label": "Weekly", "value": "Weekly"},
    {"label": "Monthly", "value": "Monthly"},
    {"label": "Yearly", "value": "Yearly"},
];
let selected_routine_mode = null;
let routineModeHasError = false;
let routineModeErrorMessage = "";
let selectedTime = "00:00";

let selected_days = ["Sun", "Mon", "Tues", "Wed", "Thu", "Fri", "Sat"];
let selected_week_day = "Sun";

$: {
    if(advanceSettings === true){
        routine_title_label = "Routine Title";
    }else{
        routine_title_label = "";
    }
}

function titleChangedHandler(event){
    routine_title = event.detail;
    if(routine_title && routineTitleHasError){
        routineTitleHasError = false;
        routineTitleErrorMessage = "";
    }
}
function routineDetailsChangeHandler(event){
    routine_details = event.detail;
    if(routine_details && routineDetailsHasError){
        routineDetailsHasError = false;
        routineDetailsErrorMessage = "";
    }
}
function routineModeChangeHandler(event){
    selected_routine_mode = event.detail;
    if(selected_routine_mode && routineModeHasError){
        routineModeHasError = false;
        routineModeErrorMessage = "";
    }
    if(selected_routine_mode.value == "Daily"){
        selected_days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
        selectedTime = "00:00"
    }else if(selected_routine_mode.value == "Weekly"){
        selected_week_day = "Sun";
    }else if(routine_details && routineDetailsHasError){
        routineDetailsHasError = false;
        routineDetailsErrorMessage = "";
    }
}
function validateRoutineForm(){
    let hasError = false;
    if( routine_title == ""){
        hasError = true;
        routineTitleHasError = true;
        routineTitleErrorMessage = "Please provide a title";
    }else{
        routineTitleHasError = false;
        routineTitleErrorMessage = "";
    }
    if( routine_details == ""){
        hasError = true;
        routineDetailsHasError = true;
        routineDetailsErrorMessage = "Please provide a description";
    }else{
        routineDetailsHasError = false;
        routineDetailsErrorMessage = "";
    }
    if( selected_routine_mode == null){
        hasError = true;
        routineModeHasError = true;
        routineModeErrorMessage = "Please select an execution mode";
    }else{
        routineModeHasError = false;
        routineModeErrorMessage = "";
    }
    return !hasError;
}
async function createRoutineHandler(event){
    if(!validateRoutineForm()){
        return;
    }
    let formData = {
        'routine_title': routine_title
    };
    try{
        let response = await createRoutine(formData);
        console.log(response);
    }catch(error){
        console.log(error);
    }
}
</script>
<style>
.form_heading{
    font-family: monospace;
    font-size: 24px;
    font-weight: bold;
    margin-top: 10px;
    padding-bottom: 40px;
}
.form_gap{
    width: 100%;
    height: 20px;
}
.card_body{
    padding: 20px;
}
.center{
    display: flex;
    justify-content: center;
    align-items: center;
}
.advance_settings_trigger{
    display: flex;
    justify-content: center;
    align-items: center;
    flex-direction: column;
    padding-top: 20px;
    position: relative;
}
.advance_settings_trigger_icon{
    position: relative;
    bottom: 10px;
    cursor: pointer;
}
.advance_settings_hide{
    padding-top: 20px;
    padding-bottom: 20px;
}
.submit_button_container{
    margin-bottom: 10px;
    display: flex;
    justify-content: center;
    align-items: center;
    padding-bottom: 10px;
}
.form_row{
    display: flex;
    flex-wrap: nowrap;
    box-sizing: border-box;
}
.form_column{
    flex: 1;
}
.mb10{
    margin-bottom: 10px;
}
.mt10{
    margin-top: 10px;
}
</style>