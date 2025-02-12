<div class="menu_bar">
    <div class="user_account">
        <i class="fa fa-user"></i>
    </div>
</div>
<style>
.user_account{
	width: 40px;
	height: 40px;
	display: flex;
	justify-content: center;
	align-items: center;
	cursor: pointer;
	color: #fff;
	background-color: transparent;
	transition: 0.3s;
}
.user_account:hover{
	background-color: #333;
}
.menu_bar{
	height: 40px;
	width: 100%;
	display: flex;
	position: fixed;
	top: 0;
	right: 0;
	z-index: 10000001;
	background-color: rgba(0, 0, 0, 0.7);
	justify-content: flex-end;
}
</style>