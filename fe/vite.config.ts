import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
	server:{
		port: 9000,
		strictPort:false,
	},
	preview:{
		port:9001,
		strictPort:false,
	}	
});
