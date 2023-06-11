<script context="module">
	import { get } from 'svelte/store';
	import { accessTokenStore } from '../stores/stores.js';

	export async function load() {
		let accessToken = get(accessTokenStore);
		let username;
		if (accessToken != '') {
			const payloadB64 = accessToken.split('.')[1];
			username = JSON.parse(window.atob(payloadB64)).username;
		}

		const url = `http://54.145.220.238:8090/users/${username}`;
		const options = {
			method: 'GET'
		};

		try {
			const res = await fetch(url, options);
			const userData = await res.json();

			return { props: { user: userData } };
		} catch (err) {
			console.log(err);
		}
	}
</script>

<script>
	import { timeToExpireStore } from '../stores/stores.js';
	export let user;
	let newUsername = user.username;
	let newFullname = user.fullname;
	let newEmail = user.email;
	let newPassword = '';
	let confirmPassword = '';

	async function submit() {
		const newProfileDetails = {
			oldUsername: user.username,
			newUsername: newUsername,
			newFullname: newFullname,
			newEmail: newEmail,
			newPassword: newPassword
		};

		const options = {
			method: 'POST',
			body: JSON.stringify(newProfileDetails),
			headers: {
				'Content-Type': 'application/json'
			}
		};
		const url = 'http://54.145.220.238:8090/auth/editprofile';

		try {
			const res = await fetch(url, options);
			const accessToken = await res.json();
			console.log(accessToken);
			accessTokenStore.set(accessToken);
			const payloadB64 = accessToken.split('.')[1];
			timeToExpireStore.set(JSON.parse(window.atob(payloadB64)).exp);
			console.log(timeToExpireStore);
		} catch (err) {
			console.log(err);
			alert(err);
		}
	}
</script>

<svelte:head>Edit Profile - go-leetcode</svelte:head>

<div class="container items-center flex flex-col">
	<div class="rounded-lg shadow-lg w-1/3 my-44">
		<div class="m-4">
			<p class="text-2xl font-medium text-center">Edit your account information.</p>
		</div>
		<form on:submit|preventDefault={submit} class="flex flex-col mx-4">
			<div class="flex flex-col items-stretch">
				<input
					bind:value={newUsername}
					type="text"
					placeholder="New username"
					required
					class="border-b-2 py-4"
				/>
				<input
					bind:value={newFullname}
					type="text"
					placeholder="New name"
					required
					class="border-b-2 py-4"
				/>
				<input
					bind:value={newEmail}
					type="email"
					placeholder="New email address"
					required
					class="border-b-2 py-4"
				/>
				<input
					bind:value={newPassword}
					type="password"
					placeholder="New password"
					required
					class="border-b-2 py-4"
				/>
				<input
					bind:value={confirmPassword}
					type="password"
					placeholder="Confirm password"
					required
					pattern={newPassword}
					class="border-b-2 py-4"
				/>
			</div>
			<div class="flex flex-col items-center my-10">
				<button type="submit" class="border bg-blue-400 rounded-lg p-3">
					<a href="/" class="mx-3 my-2 text-white">Edit Profile</a>
				</button>
			</div>
		</form>
	</div>
</div>
