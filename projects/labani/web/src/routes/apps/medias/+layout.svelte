<div class="app">
	<slot />
	<SvelteToast {toastOptions} />
</div>
<style>
.app {
	min-height: 100vh;
}
:global(._toastContainer) {
	z-index: 10001 !important;
}
</style>
<script>
import { SvelteToast } from '@zerodevx/svelte-toast';
const toastOptions = {};
</script>