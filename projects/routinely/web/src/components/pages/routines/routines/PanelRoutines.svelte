<script>
	import { onMount } from "svelte";
	import RoutineItem from "components/pages/routines/routines/RoutineItem.svelte";
	import NoItemRoutineList from "components/pages/routines/routines/NoItemRoutineList.svelte";
	import { goto } from "$app/navigation";
	import { getAllRoutines } from "apis/apis.js";

	export let search = "";
	let routines = [];

	function routineClickedHandler(event) {
		let slug = event.detail.Slug;
		goto("/routine/" + slug);
	}

	let page = 1; // Current page to fetch
	let loading = false; // Loading state
	let hasMore = true; // Check if there's more data

	const fetchRoutines = async () => {
		if (loading || !hasMore) return; // Avoid duplicate requests
		loading = true;

		try {
			let formData = {
				page: page,
				search: search,
			};
			const response = await getAllRoutines(formData);
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
			fetchRoutines(); // Trigger data fetch when close to bottom of the page
		}
	};

	$: {
		if (search) {
			hasMore = true;
			page = 1;
			routines = [];
			fetchRoutines();
		} else {
			hasMore = true;
			page = 1;
			routines = [];
			fetchRoutines();
		}
	}

	onMount(() => {
		fetchRoutines(); // Fetch initial data
		window.addEventListener("scroll", onScroll);

		return () => {
			window.removeEventListener("scroll", onScroll); // Cleanup on unmount
		};
	});
</script>

<div class="routine_container">
	{#if routines && routines.length == 0}
		<NoItemRoutineList></NoItemRoutineList>
	{:else}
		{#each routines as routine}
			<RoutineItem {routine} on:click={routineClickedHandler}
			></RoutineItem>
		{/each}
	{/if}

	{#if loading}
		<div class="loading">Loading...</div>
	{/if}

	{#if !hasMore}
		<div class="loading">No more Routines to load</div>
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
