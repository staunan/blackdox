<script>
	import { createEventDispatcher } from "svelte";
	const dispatch = createEventDispatcher();

	export let checked = false;

	function checkedChangeHandler(event) {
		if (!checked) {
			dispatch("change", true);
			checked = true;
		}
	}
</script>

<div class:editdisabled={checked}>
	<label class="field">
		<input
			class="field__checkbox"
			type="checkbox"
			{checked}
			on:change={checkedChangeHandler}
		/>
		<span class="field__label">
			<svg class="check" viewBox="0 0 64 64" aria-hidden="true">
				<g transform="translate(32,32)">
					<g stroke-linecap="round" stroke-width="3">
						<polyline
							class="check__stroke-offset check__stroke-offset--1"
							stroke="var(--purple)"
							points="-30 -30,-42 -42"
							stroke-dasharray="17 17"
							stroke-dashoffset="17"
						/>
						<polyline
							class="check__stroke-offset check__stroke-offset--2"
							stroke="var(--primary)"
							points="38 -38,54 -54"
							stroke-dasharray="22.63 22.63"
							stroke-dashoffset="22.63"
						/>
						<polyline
							class="check__stroke-offset check__stroke-offset--3"
							stroke="var(--green)"
							points="-28 28,-40 40"
							stroke-dasharray="17 17"
							stroke-dashoffset="17"
						/>
						<polyline
							class="check__stroke-offset check__stroke-offset--4"
							stroke="var(--red)"
							points="32 32,44 44"
							stroke-dasharray="17 17"
							stroke-dashoffset="17"
						/>
					</g>
					<g>
						<circle
							class="check__move-fade check__move-fade--1"
							fill="var(--red)"
							r="3"
							cx="4"
							cy="-44"
							opacity="0"
						/>
						<circle
							class="check__move-fade check__move-fade--2"
							fill="var(--primary)"
							r="3"
							cx="-44"
							cy="-8"
							opacity="0"
						/>
						<circle
							class="check__move-fade check__move-fade--3"
							fill="var(--green)"
							r="3"
							cx="52"
							cy="12"
							opacity="0"
						/>
						<circle
							class="check__move-fade check__move-fade--4"
							fill="var(--purple)"
							r="2"
							cx="-2"
							cy="40"
							opacity="0"
						/>
						<circle
							class="check__move-fade check__move-fade--5"
							fill="var(--primary)"
							r="3"
							cx="-12"
							cy="46"
							opacity="0"
						/>
					</g>
					<g
						class="check__scale-out"
						fill="none"
						stroke="var(--check-outline)"
						stroke-width="2"
					>
						<circle r="30" />
						<polygon
							points="-10 -4,-16 2,-4 14,16 -6,10 -12,-4 2"
						/>
					</g>
					<g class="check__fade" opacity="0">
						<circle
							class="check__scale-in check__scale-in--1"
							fill="var(--check-bubble)"
							r="30.9"
						/>
						<circle
							class="check__scale-in check__scale-in--2"
							fill="var(--primary)"
							r="31"
						/>
						<polygon
							class="check__scale-in check__scale-in--3"
							fill="var(--white)"
							stroke="var(--primary)"
							stroke-width="2"
							points="-10 -4,-16 2,-4 14,16 -6,10 -12,-4 2"
						/>
					</g>
				</g>
			</svg>
			<span class="field__sr-only">Check circle</span>
		</span>
	</label>
</div>

<style>
	:root {
		--hue: 223;
		--sat: 10%;
		--bg: hsl(var(--hue), var(--sat), 90%);
		--fg: hsl(var(--hue), var(--sat), 10%);
		--primary: hsl(var(--hue), 90%, 50%);
		--green: hsl(150, 90%, 40%);
		--red: hsl(0, 90%, 50%);
		--purple: hsl(270, 90%, 60%);
		--white: hsl(0, 0%, 100%);
		--light-gray1: hsl(var(--hue), var(--sat), 85%);
		--light-gray2: hsl(var(--hue), var(--sat), 65%);
		--dark-gray1: hsl(var(--hue), var(--sat), 35%);
		--dark-gray2: hsl(var(--hue), var(--sat), 15%);
		--check-bubble: var(--light-gray1);
		--check-outline: var(--light-gray2);
		--trans-dur: 0.3s;
		font-size: clamp(1rem, 0.95rem + 0.25vw, 1.25rem);
	}
	.check {
		overflow: visible;
		pointer-events: none;
		width: 3em;
		height: auto;
	}
	.editdisabled {
		pointer-events: none;
		cursor: not-allowed;
	}
	.check circle,
	.check polygon {
		transition:
			fill var(--trans-dur),
			stroke var(--trans-dur);
	}
	.check__move-fade,
	.check__scale-in,
	.check__scale-out,
	.check__stroke-offset {
		animation-duration: 1s;
		animation-timing-function: cubic-bezier(0.37, 0, 0.63, 1);
		animation-fill-mode: forwards;
	}
	.check__fade {
		transition: opacity var(--trans-dur);
	}
	.check__move-fade {
		animation-timing-function: cubic-bezier(0.61, 1, 0.88, 1);
	}

	.field {
		display: flex;
		margin: auto;
	}
	.field,
	.field__checkbox {
		-webkit-tap-highlight-color: transparent;
	}
	.field__checkbox {
		-webkit-appearance: none;
		appearance: none;
	}
	.field__label {
		cursor: pointer;
		display: flex;
		gap: 0.5em;
		align-items: center;
	}
	.field__checkbox:checked + .field__label .check__fade {
		opacity: 1;
		transition-duration: 0s;
	}
	.field__checkbox:checked + .field__label .check__move-fade--1 {
		animation-name: move-fade1;
	}
	.field__checkbox:checked + .field__label .check__move-fade--2 {
		animation-name: move-fade2;
	}
	.field__checkbox:checked + .field__label .check__move-fade--3 {
		animation-name: move-fade3;
	}
	.field__checkbox:checked + .field__label .check__move-fade--4 {
		animation-name: move-fade4;
	}
	.field__checkbox:checked + .field__label .check__move-fade--5 {
		animation-name: move-fade5;
	}
	.field__checkbox:checked + .field__label .check__scale-in {
		transform: scale(1);
	}
	.field__checkbox:checked + .field__label .check__scale-in--1 {
		animation-name: scale-in1;
	}
	.field__checkbox:checked + .field__label .check__scale-in--2 {
		animation-name: scale-in2;
	}
	.field__checkbox:checked + .field__label .check__scale-in--3 {
		animation-name: scale-in3;
	}
	.field__checkbox:checked + .field__label .check__scale-out {
		animation-name: scale-out;
	}
	.field__checkbox:checked + .field__label .check__stroke-offset--1 {
		animation-name: stroke-offset1;
	}
	.field__checkbox:checked + .field__label .check__stroke-offset--2 {
		animation-name: stroke-offset2;
	}
	.field__checkbox:checked + .field__label .check__stroke-offset--3 {
		animation-name: stroke-offset3;
	}
	.field__checkbox:checked + .field__label .check__stroke-offset--4 {
		animation-name: stroke-offset4;
	}
	.field__sr-only {
		overflow: hidden;
		position: absolute;
		width: 1px;
		height: 1px;
	}

	/* Dark theme */
	@media (prefers-color-scheme: dark) {
		:root {
			--bg: hsl(var(--hue), var(--sat), 10%);
			--fg: hsl(var(--hue), var(--sat), 90%);
			--check-bubble: var(--dark-gray2);
			--check-outline: var(--dark-gray1);
		}
	}
	/* Animations */
	@keyframes move-fade1 {
		from {
			animation-timing-function: steps(1, end);
			opacity: 0;
			transform: translate(0, 16px);
		}
		31% {
			opacity: 1;
			transform: translate(0, 16px);
		}
		75% {
			opacity: 1;
			transform: translate(0, 0);
		}
		92.75%,
		to {
			opacity: 0;
			transform: translate(0, 0);
		}
	}
	@keyframes move-fade2 {
		from {
			animation-timing-function: steps(1, end);
			opacity: 0;
			transform: translate(22px, 0);
		}
		35.25% {
			opacity: 1;
			transform: translate(22px, 0);
		}
		75% {
			opacity: 1;
			transform: translate(0, 0);
		}
		92.75%,
		to {
			opacity: 0;
			transform: translate(0, 0);
		}
	}
	@keyframes move-fade3 {
		from {
			animation-timing-function: steps(1, end);
			opacity: 0;
			transform: translate(-27px, 0);
		}
		44% {
			opacity: 1;
			transform: translate(-27px, 0);
		}
		82.25% {
			opacity: 1;
			transform: translate(0, 0);
		}
		to {
			opacity: 0;
			transform: translate(0, 0);
		}
	}
	@keyframes move-fade4 {
		from {
			animation-timing-function: steps(1, end);
			opacity: 0;
			transform: translate(0, -11px);
		}
		44% {
			opacity: 1;
			transform: translate(0, -11px);
		}
		82.25% {
			opacity: 1;
			transform: translate(0, 0);
		}
		to {
			opacity: 0;
			transform: translate(0, 0);
		}
	}
	@keyframes move-fade5 {
		from {
			animation-timing-function: steps(1, end);
			opacity: 0;
			transform: translate(0, -23px);
		}
		22% {
			opacity: 1;
			transform: translate(0, -23px);
		}
		61.75% {
			opacity: 1;
			transform: translate(0, 0);
		}
		79.5%,
		to {
			opacity: 0;
			transform: translate(0, 0);
		}
	}
	@keyframes scale-in1 {
		from {
			transform: scale(0);
		}
		26.5%,
		to {
			transform: scale(1);
		}
	}
	@keyframes scale-in2 {
		from,
		13.25% {
			transform: scale(0);
		}
		44% {
			transform: scale(1.05);
		}
		48.5%,
		to {
			transform: scale(1);
		}
	}
	@keyframes scale-in3 {
		from,
		35.25% {
			transform: scale(0);
		}
		66.25% {
			transform: scale(1.05);
		}
		70.5%,
		to {
			transform: scale(1);
		}
	}
	@keyframes scale-out {
		from {
			transform: scale(1);
		}
		31%,
		to {
			transform: scale(0);
		}
	}
	@keyframes stroke-offset1 {
		from,
		20.5% {
			stroke-dashoffset: 17;
		}
		41% {
			stroke-dashoffset: 0;
		}
		61.75%,
		to {
			stroke-dashoffset: -17;
		}
	}
	@keyframes stroke-offset2 {
		from,
		29.5% {
			stroke-dashoffset: 22.63;
		}
		50% {
			stroke-dashoffset: 0;
		}
		70.5%,
		to {
			stroke-dashoffset: -22.63;
		}
	}
	@keyframes stroke-offset3 {
		from,
		38% {
			stroke-dashoffset: 17;
		}
		58.75% {
			stroke-dashoffset: 0;
		}
		79.5%,
		to {
			stroke-dashoffset: -17;
		}
	}
	@keyframes stroke-offset4 {
		from,
		22% {
			stroke-dashoffset: 17;
		}
		42.75% {
			stroke-dashoffset: 0;
		}
		63.25%,
		to {
			stroke-dashoffset: -17;
		}
	}
</style>
