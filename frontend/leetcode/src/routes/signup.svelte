<script>
	import { accessTokenStore, timeToExpireStore } from '../stores/stores';
	let username = '';
	let fullname = '';
	let email = '';
	let password = '';

	async function submit() {
		const signupDetails = {
			username: username,
			fullname: fullname,
			email: email,
			password: password
		};

		const options = {
			method: 'POST',
			body: JSON.stringify(signupDetails),
			headers: {
				'Content-Type': 'application/json'
			}
		};
		const url = 'http://backend/auth/signup';

		try {
			const res = await fetch(url, options);
			const accessToken = await res.json();
			console.log(accessToken);
			accessTokenStore.set(accessToken);
			const payloadB64 = accessToken.split('.')[1];
			timeToExpireStore.set(JSON.parse(window.atob(payloadB64)).exp);
		} catch (err) {
			alert(err);
		}
	}
</script>

<svelte:head>
	<title>Sign Up - go-leetcode</title>
</svelte:head>

<div class="container mx-auto items-center flex flex-col text-gray-900">
	<div class="bg-white p-16 rounded-lg border-gray-900 shadow-lg my-20">
		<h2 class="text-2xl font-medium text-center mb-20">Join go-leetcode.</h2>
		<form on:submit|preventDefault={submit} class="flex flex-col">
			<div class="flex flex-col items-stretch">
				<input bind:value={username} type="text" placeholder="Username" class="border-b-2 py-4" />
				<input bind:value={fullname} type="text" placeholder="Full Name" class="border-b-2 py-4" />
				<input bind:value={email} type="text" placeholder="Email address" class="border-b-2 py-4" />
				<input
					bind:value={password}
					type="password"
					placeholder="Password"
					class="border-b-2 py-4"
				/>
			</div>
			<div class="flex flex-col items-center my-10">
				<button type="submit" class="border bg-blue-400 rounded-lg p-3">
					<a href="/" class="mx-3 my-2 text-white">Sign up</a>
				</button>
			</div>
			<p class="text-sm font-extralight">
				Already have an account? <a href="/login" class="text-blue-400 cursor-pointer">Log in</a>.
			</p>
		</form>
	</div>
</div>
