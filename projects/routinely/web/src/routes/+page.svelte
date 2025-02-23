<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Card>
		<div class="card_body">
			<h1 class="text-center mb10">Create Routine</h1>
			<TextBox label="Title" placeholder="Write a title for your routine..." on:change={titleChangedHandler} value={routine_title}></TextBox>
			<div class="form_gap"></div>
			<TextArea label="Description" placeholder="Write a description for your routine (Optional)..." on:change={routineDetailsChangeHandler} value={routine_details}></TextArea>
			<div class="form_gap"></div>
			<Dropdown label="Routine Mode" placeholder="Select a execution mode of this routine " items={all_routine_modes} currentitem={selected_routine_mode} on:change={routineModeChangeHandler} />
			<div class="text-center mt10">
				<SubmitButton title="Create Routine" on:tap={createRoutineHandler}></SubmitButton>
			</div>
		</div>
	</Card>
</section>
<script>
import Card from 'components/Card.svelte';
import TextBox from 'components/TextBox.svelte';
import TextArea from 'components/TextArea.svelte';
import Dropdown from 'components/Dropdown.svelte';
import SubmitButton from 'components/SubmitButton.svelte';
import { createRoutine } from 'apis/apis.js';

let routine_title = "";
let routine_details = "";
let all_routine_modes = [
	{"label": "Daily", "value": "Daily"},
	{"label": "Weekly", "value": "Weekly"},
	{"label": "Monthly", "value": "Monthly"},
	{"label": "Yearly", "value": "Yearly"},
];
let selected_routine_mode = null;

function titleChangedHandler(event){
    routine_title = event.detail;
}
function routineModeChangeHandler(event){
    selected_routine_mode = event.detail;
}
async function createRoutineHandler(event){
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
function routineDetailsChangeHandler(event){
    selected_routine_mode = event.detail;
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
