<svelte:head>
	<title>Home</title>
	<meta name="description" content="Svelte demo app" />
</svelte:head>
<section>
	<CreateRoutineForm></CreateRoutineForm>
</section>

<script>
import CreateRoutineForm from 'components/CreateRoutineForm.svelte';
</script>

<style>
section {
	display: flex;
	flex-direction: column;
	justify-content: center;
	align-items: center;
	flex: 0.6;
}
</style>
