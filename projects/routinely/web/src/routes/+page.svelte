<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>

<section>
	<Card>
		<div class="card_body">
			<h1 class="text-center mb10">Create Routine</h1>
			<TextBox label="Title" placeholder="Write a title for your routine..." on:change={titleChangedHandler} value={routine_title}></TextBox>
			<div class="text-center mt10">
				<SubmitButton title="Create Routine" on:tap={createRoutineHandler}></SubmitButton>
			</div>
		</div>
	</Card>
</section>
<script>
import Card from 'components/Card.svelte';
import TextBox from 'components/TextBox.svelte';
import SubmitButton from 'components/SubmitButton.svelte';
import { createRoutine } from 'apis/apis.js';

let routine_title = "";

function titleChangedHandler(event){
    routine_title = event.detail;
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
</script>
<style>
	section {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		flex: 0.6;
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
