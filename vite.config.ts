import devtoolsJson from 'vite-plugin-devtools-json';
import tailwindcss from '@tailwindcss/vite';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import { resolve } from 'node:path';

export default defineConfig({
	plugins: [tailwindcss(), sveltekit(), devtoolsJson()],
	resolve: {
		alias: {
			'/content': resolve(process.cwd(), 'content')
		}
	},
	// Add this server configuration
	server: {
		fs: {
			// Allow serving files from the project root
			allow: ['.']
		}
	}
});