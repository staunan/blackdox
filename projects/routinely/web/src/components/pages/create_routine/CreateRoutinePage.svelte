<script>
	import CreateRoutineForm from "components/pages/create_routine/CreateRoutineForm.svelte";
	import FormHeadingTitle from "components/form/FormHeadingTitle.svelte";
	import Center from "components/layouts/Center.svelte";
	import Card from "components/Card.svelte";
</script>

<div class="create_routine_form">
	<Card padding={6}>
		<Center>
			<FormHeadingTitle title="Create Routine"></FormHeadingTitle>
		</Center>
		<div class="form">
			<CreateRoutineForm></CreateRoutineForm>
		</div>
	</Card>
</div>

<style>
	.create_routine_form {
		width: 70%;
		margin: auto;
		padding-top: 25px;
	}
	.form {
		padding-top: 30px;
	}
</style>
