<div class="row">
	<div class="column_left">
		<slot name="left"></slot>
	</div>
	<div class="column_center">
		<slot name="center"></slot>
	</div>
	<div class="column_right">
		<slot name="right"></slot>
	</div>
</div>

<style>
	.row {
		width: 100%;
		display: flex;
		flex-wrap: wrap;
		box-sizing: border-box;
		justify-content: space-between;
		height: inherit;
	}
	.column_left,
	.column_center,
	.column_right {
		flex: 1;
	}
	.column_left {
		display: flex;
		justify-content: flex-start;
		align-items: flex-start;
	}
	.column_center {
		display: flex;
		justify-content: center;
		align-items: flex-start;
	}
	.column_right {
		display: flex;
		justify-content: flex-end;
		align-items: flex-start;
	}
</style>
