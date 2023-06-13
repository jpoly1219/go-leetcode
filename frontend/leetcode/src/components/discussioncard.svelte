<script>
	import { createEventDispatcher, onMount } from 'svelte';

	export let discussion;

	const dispatch = createEventDispatcher();
	const switchComponent = () => {
		dispatch('switch', {
			component: 'Singlediscussion',
			props: {
				discussion: discussion
			}
		});
	};

	let userData = {
		username: '',
		fullname: '',
		email: '',
		profilePic: ''
	};
	onMount(async () => {
		const url = `https://goleetcode.xyz:8090/backend/users/${discussion.author}`;
		const options = {
			method: 'GET'
		};
		try {
			const res = await fetch(url, options);
			const data = await res.json();
			userData = {
				username: data.username,
				fullname: data.fullname,
				email: data.email,
				profilePic: data.profilePic
			};
		} catch (err) {
			console.log(err);
		}
	});
</script>

<div>
	<div class="flex flex-row items-center py-2">
		<img src={userData.profilePic} class="rounded-full w-12 h-12 mr-4" />
		<div class="flex flex-col">
			<p on:click={switchComponent} class="text-base cursor-pointer">{discussion.title}</p>
			<p class="text-xs text-gray-500">{discussion.author} created at: {discussion.created}</p>
		</div>
	</div>
</div>
