<script>
	import Nav from '../components/nav.svelte';
	import { accessTokenStore, timeToExpireStore } from '../stores/stores';

	async function refresh() {
		const options = {
			method: 'GET',
			credentials: 'include'
		};

		try {
			const res = await fetch('https://backend/auth/silentrefresh', options);
			const accessToken = await res.json();
			const payloadB64 = accessToken.split('.')[1];
			timeToExpireStore.set(JSON.parse(window.atob(payloadB64)).exp);
		} catch (err) {
			alert(err);
		}
	}

	function refreshTimer() {
		if ($timeToExpireStore != '') {
			var i = Date.now() / 1000;
			var timer = setInterval(() => {
				console.log($timeToExpireStore);
				if (i >= Number($timeToExpireStore)) {
					refresh();
					clearInterval(timer);
				}
				console.log(Date.now() / 1000);
				i++;
			}, 1000);
		}
	}

	$: $timeToExpireStore, refreshTimer();
	$: $accessTokenStore, console.log($accessTokenStore);
</script>

<div class="h-screen">
	<Nav />
	<div class="px-8 pt-24 pb-8 h-full">
		<slot />
	</div>
</div>

<style>
	@tailwind base;
	@tailwind components;
	@tailwind utilities;
</style>
