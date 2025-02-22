import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import path from 'path';

export default defineConfig({
	plugins: [sveltekit()],
	resolve: {
        alias: {
            lib: path.resolve('src/lib/'),
            components: path.resolve('src/components/'),
            apis: path.resolve('src/apis/'),
            config: path.resolve('src/config/'),
        },
    }
});
