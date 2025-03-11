<div class="crossmark_icons">
	<svg
		class="crossmark animateElement"
		xmlns="http://www.w3.org/2000/svg"
		viewBox="0 0 52 52"
		><circle
			class="crossmark__circle animateElement"
			cx="26"
			cy="26"
			r="25"
			fill="none"
		/>
		<path
			class="cross__path cross__path--right animateElement"
			fill="none"
			d="M16,16 l20,20"
		/>
		<path
			class="cross__path cross__path--left animateElement"
			fill="none"
			d="M16,36 l20,-20"
		/>
	</svg>
</div>

<style>
	@keyframes stroke {
		100% {
			stroke-dashoffset: 0;
		}
	}
	@keyframes scale {
		0%,
		100% {
			transform: none;
		}
		50% {
			transform: scale3d(1.1, 1.1, 1);
		}
	}
	@keyframes fill {
		100% {
			box-shadow: inset 0px 0px 0px 30px #7ac142;
		}
	}
	.crossmark_icons {
		height: 80px;
		width: 80px;
		position: relative;
	}
	.crossmark__circle.animateElement {
		stroke-dasharray: 166;
		stroke-dashoffset: 166;
		stroke-width: 2;
		stroke-miterlimit: 10;
		stroke: #ff0000;
		fill: none;
		animation: stroke 0.6s cubic-bezier(0.65, 0, 0.45, 1) forwards;
	}

	.crossmark.animateElement {
		width: 56px;
		height: 56px;
		border-radius: 50%;
		display: block;
		stroke-width: 2;
		stroke: #ff0000;
		stroke-miterlimit: 10;
		margin: 10% auto;
		box-shadow: inset 0px 0px 0px #ff0000;
		animation:
			fillred 0.4s ease-in-out 0.4s forwards,
			scale 0.3s ease-in-out 0.9s both;
	}

	.cross__path.animateElement {
		stroke: #ffffff;
		stroke-dasharray: 48;
		stroke-dashoffset: 48;
		transform-origin: 50% 50% 0;
	}
	.cross__path--right.animateElement {
		animation: 0.3s ease 0.8s normal forwards 1 running stroke;
	}
	.cross__path--left.animateElement {
		animation: 1s ease 0.8s normal forwards 1 running stroke;
	}
	@keyframes fillred {
		100% {
			box-shadow: inset 0px 0px 0px 30px #ff0000;
		}
	}
	@keyframes stroke {
		100% {
			stroke-dashoffset: 0;
		}
	}
	@-webkit-keyframes rotate {
		100% {
			-webkit-transform: rotate(360deg);
			transform: rotate(360deg);
		}
	}

	@keyframes rotate {
		100% {
			-webkit-transform: rotate(360deg);
			transform: rotate(360deg);
		}
	}
	@-webkit-keyframes dash {
		0% {
			stroke-dasharray: 1, 200;
			stroke-dashoffset: 0;
		}
		50% {
			stroke-dasharray: 89, 200;
			stroke-dashoffset: -35;
		}
		100% {
			stroke-dasharray: 89, 200;
			stroke-dashoffset: -124;
		}
	}
	@keyframes dash {
		0% {
			stroke-dasharray: 1, 200;
			stroke-dashoffset: 0;
		}
		50% {
			stroke-dasharray: 89, 200;
			stroke-dashoffset: -35;
		}
		100% {
			stroke-dasharray: 89, 200;
			stroke-dashoffset: -124;
		}
	}
	@-webkit-keyframes color {
		0% {
			stroke: #ffffff;
		}
		40% {
			stroke: #ffffff;
		}
		66% {
			stroke: #ffffff;
		}
		80%,
		90% {
			stroke: #ffffff;
		}
	}
	@keyframes color {
		0% {
			stroke: #ffffff;
		}
		40% {
			stroke: #ffffff;
		}
		66% {
			stroke: #ffffff;
		}
		80%,
		90% {
			stroke: #ffffff;
		}
	}
</style>
