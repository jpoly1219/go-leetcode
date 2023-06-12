<script>
	import { onMount } from 'svelte';

	export let comment;

	let userData = {
		username: '',
		fullname: '',
		email: '',
		profilePic: ''
	};
	onMount(async () => {
		const url = `https://backend/users/${comment.author}`;
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
			console.log(data);
		} catch (err) {
			console.log(err);
		}
	});
</script>

<div class="">
	<div class="">
		<div class="flex flex-row items-center mb-2">
			<img src={userData.profilePic} class="rounded-full w-10 h-10 mr-2" />
			<p class="text-sm mr-2">{comment.author}</p>
			<p class="text-sm">{comment.created}</p>
		</div>
		<div class="flex flex-row">
			<div class="w-10 mr-2" />
			<p>{comment.description}</p>
		</div>
	</div>
</div>
