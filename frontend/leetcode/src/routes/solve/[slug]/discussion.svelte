<script context="module">
	import Discussions from '../../../components/discussions.svelte';

	export async function load({ page }) {
		const fullPath = page.path;
		const slugArray = fullPath.split('/');
		const slug = slugArray[2];

		const url = `https://goleetcode.xyz/backend/discussions/${slug}`;

		const options = {
			method: 'GET'
		};

		try {
			const res = await fetch(url, options);
			const discussions = await res.json();
			return { props: { slug, discussions } };
		} catch (err) {
			alert(err);
		}
	}
</script>

<script>
	export let slug;
	export let discussions;
</script>

<div>
	<Discussions {slug} {discussions} />
</div>
