import node from '@sveltejs/adapter-node';

/** @type {import('@sveltejs/kit').Config} */

const config = {
	kit: {
		adapter: node(),
		// hydrate the <div id="svelte"> element in src/app.html
		target: '#svelte',
		vite: {
			server: {
				hmr: {
					host: 'localhost',
					port: 15000,
					protocol: 'ws'
				}
			}
		}
	}
};

export default config;
