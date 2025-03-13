<script>
	export let data = null;
</script>

<div class="routine_details">
	{#if data}
		<div class="routine_title">{data.Title}</div>
	{/if}
</div>
