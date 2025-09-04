<script>
	import { onMount } from "svelte";
	import { goto } from "$app/navigation";
	import TrashRoutineItem from "components/pages/routines/trash/TrashRoutineItem.svelte";
	import NoItemInTrash from "components/pages/routines/trash/NoItemInTrash.svelte";
	import { getTrashedRoutines } from "apis/apis.js";

	let routines = $state([]);

	let page = 1; // Current page to fetch
	let loading = $state(false); // Loading state
	let hasMore = $state(true); // Check if there's more data

	const fetchTrashedRoutines = async () => {
		if (loading || !hasMore) return; // Avoid duplicate requests
		loading = true;

		try {
			let formData = {
				page: page,
			};
			const response = await getTrashedRoutines(formData);
			const data = response.Data;

			if (!data || (data && data.length === 0)) {
				hasMore = false; // No more data
			} else {
				routines = [...routines, ...data];
				page += 1; // Move to the next page
			}
		} catch (error) {
			console.error("Error fetching data:", error);
		}
		loading = false;
	};

	const onScroll = () => {
		const bottom =
			window.innerHeight + window.scrollY >=
			document.body.offsetHeight - 100;
		if (bottom && !loading && hasMore) {
			fetchTrashedRoutines(); // Trigger data fetch when close to bottom of the page
		}
	};

	onMount(() => {
		fetchTrashedRoutines(); // Fetch initial data
		window.addEventListener("scroll", onScroll);

		return () => {
			window.removeEventListener("scroll", onScroll); // Cleanup on unmount
		};
	});

	function routineRestoredHandler(r) {
		routines = routines.filter((routine) => {
			return routine.ID !== r.ID;
		});
	}
	function routineDeletedHandler(r) {
		routines = routines.filter((routine) => {
			return routine.ID !== r.ID;
		});
	}
</script>

<div class="routine_container">
	{#if routines && routines.length == 0}
		<NoItemInTrash></NoItemInTrash>
	{:else}
		{#each routines as routine}
			<TrashRoutineItem
				{routine}
				on:restored={() => {
					routineRestoredHandler(routine);
				}}
				on:deleted={() => {
					routineDeletedHandler(routine);
				}}
			></TrashRoutineItem>
		{/each}
	{/if}

	{#if loading}
		<div class="loading">Loading...</div>
	{/if}

	{#if !hasMore}
		<div class="loading">No more items in trash</div>
	{/if}
</div>

<style>
	.routine_container {
		padding-top: 20px;
	}
	.loading {
		text-align: center;
		margin: 20px 0;
	}
</style>
