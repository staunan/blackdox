<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Card>
		<div class="card_body">
			<h1 class="text-center mb10">Create Routine</h1>
			<TextBox label="Title" placeholder="Write a title for your routine..." on:change={titleChangedHandler} value={routine_title} hasError={routineTitleHasError} errorMessage={routineTitleErrorMessage}></TextBox>
			<div class="form_gap"></div>
			<TextArea label="Description" placeholder="Write a description for your routine (Optional)..." on:change={routineDetailsChangeHandler} value={routine_details} hasError={routineDetailsHasError} errorMessage={routineDetailsErrorMessage}></TextArea>
			<div class="form_gap"></div>
			<Dropdown label="Routine Mode" placeholder="Select a execution mode of this routine " items={all_routine_modes} currentitem={selected_routine_mode} on:change={routineModeChangeHandler} hasError={routineModeHasError} errorMessage={routineModeErrorMessage} />
			{#if selected_routine_mode && selected_routine_mode.value === "Daily"}
				<div class="form_gap"></div>
				<DaysSelector selectedItems={selected_days} label="[Optional] You can also deselect following day buttons to exclude from the execution days"></DaysSelector>
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

			<div class="form_gap"></div>
			<TimePicker label="Choose a time of the day"></TimePicker>
			
			<div class="text-center mt10">
				<SubmitButton title="Create Routine" on:tap={createRoutineHandler}></SubmitButton>
			</div>
		</div>
	</Card>
</section>
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

import { createRoutine } from 'apis/apis.js';

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

let month_date = new Date();

let selected_days = ["Sun", "Mon", "Tues", "Wed", "Thu", "Fri", "Sat"];
let selected_week_day = "Sun";

function titleChangedHandler(event){
    routine_title = event.detail;
}
function routineDetailsChangeHandler(event){
    routine_details = event.detail;
}
function routineModeChangeHandler(event){
    selected_routine_mode = event.detail;
	if(selected_routine_mode.value == "Daily"){
		selected_days = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
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
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
	}
	.form_gap{
		width: 100%;
		height: 20px;
	}
	.card_body{
		padding: 20px;
	}
	.text-center{
		display: flex;
		justify-content: center;
		align-items: center;
	}
	.mb10{
		margin-bottom: 10px;
	}
	.mt10{
		margin-top: 10px;
	}
</style>
