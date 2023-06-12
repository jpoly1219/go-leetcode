<script>
	import { createEventDispatcher, onMount } from 'svelte';
	import { accessTokenStore } from '../stores/stores';
	import Comment from './comment.svelte';
	import snarkdown from 'snarkdown';

	let username;
	if ($accessTokenStore != '') {
		const payloadB64 = $accessTokenStore.split('.')[1];
		username = JSON.parse(window.atob(payloadB64)).username;
	}

	export let discussion;

	const dispatch = createEventDispatcher();
	const switchComponent = () => {
		dispatch('switch', {
			component: 'Discussioncard'
		});
	};

	let comments = [];
	onMount(async () => {
		const url = `https://backend/discussions/${discussion.slug}/${discussion.id}`;
		const options = {
			method: 'GET'
		};

		try {
			const res = await fetch(url, options);
			const data = await res.json();
			comments = data.map((data) => {
				return {
					id: data.id,
					author: data.author,
					discussionId: data.discussionId,
					description: data.description,
					created: data.created
				};
			});
		} catch (err) {
			console.log(err);
		}
	});

	let newComment = '';
	const postComment = async () => {
		console.log('running postComment');
		const url = `https://backend/discussions/${discussion.slug}/${discussion.id}/newcomment`;
		const newCommentData = {
			author: username,
			description: newComment
		};
		const options = {
			method: 'POST',
			body: JSON.stringify(newCommentData)
		};
		const res = await fetch(url, options);
		const data = await res.json();
		comments = [...comments, data];
	};
</script>

<div class="overflow-auto">
	<div class="flex flex-row divide-x divide-solid divide-gray-300 items-center">
		<p on:click={switchComponent} class="text-sm mr-3 my-2 cursor-pointer">&lt; Back</p>
		<p class="font-bold text-lg px-3 my-2">{discussion.title}</p>
	</div>
	<div>
		<p class="text-sm mb-2">{discussion.author}</p>
		<p class="prose max-w-max">{@html snarkdown(discussion.description)}</p>
	</div>
	<div class="my-5 border-t border-b border-gray-200">
		<p class="text-base my-2">Comments:</p>
	</div>
	<div class="mb-5">
		<textarea
			bind:value={newComment}
			class="w-full h-24 border border-gray-200 rounded"
			placeholder="Type comment here... (Markdown is supported)"
		/>
		<span
			on:click={postComment}
			class="bg-gray-600 px-2 py-1 mt-3 text-sm text-white cursor-pointer rounded">Post</span
		>
	</div>
	{#each comments as comment}
		<Comment {comment} />
	{/each}
</div>
