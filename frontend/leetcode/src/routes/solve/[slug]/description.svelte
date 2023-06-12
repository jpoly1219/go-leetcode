<script context="module">
	import { get } from 'svelte/store';
	import { accessTokenStore } from '../../../stores/stores.js';
	import snarkdown from 'snarkdown';
	import { goto } from '$app/navigation';

	export async function load({ page }) {
		const fullPath = page.path;
		const slugArray = fullPath.split('/');
		const slug = slugArray[2];

		const url = `http://backend/solve/${slug}`;

		let accessToken = get(accessTokenStore);
		if (accessToken == '') {
			goto('/login');
		}

		const options1 = {
			method: 'GET',
			headers: {
				Authorization: 'Bearer ' + accessToken
			},
			credentials: 'include'
		};

		try {
			const res = await fetch(url, options1);
			const problem = await res.json();
			return { props: { problem } };
		} catch (err) {
			console.log(err);
		}
	}
</script>

<script>
	export let problem;
</script>

<div>
	<p class="text-lg font-bold mb-3">{problem.title}</p>
	<p
		class="text-sm font-light {problem.difficulty === 'easy'
			? 'text-green-500'
			: problem.difficulty === 'medium'
			? 'text-yellow-500'
			: 'text-red-500'}"
	>
		{problem.difficulty}
	</p>
	<hr class="my-4" />
	<p class="prose max-w-max">{@html snarkdown(problem.description)}</p>
</div>
