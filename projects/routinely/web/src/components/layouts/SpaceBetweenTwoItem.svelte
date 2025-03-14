<div class="row">
	<div class="column_left">
		<slot name="left"></slot>
	</div>
	<div class="column_right">
		<slot name="right"></slot>
	</div>
</div>

<style>
	.row {
		display: flex;
		flex-wrap: wrap;
		box-sizing: border-box;
		justify-content: space-between;
	}
	.column_left,
	.column_right {
		flex: 1;
	}
	.column_left {
		display: flex;
		justify-content: flex-start;
		align-items: flex-start;
	}
	.column_right {
		display: flex;
		justify-content: flex-start;
		align-items: flex-start;
	}
</style>
