<script>
    import Nav from "../components/nav.svelte"
    import { timeToExpire } from "../stores/timer"

    async function refresh() {
		const options = {
			method: "POST",
			credentials: "include"
		}

		try {
			const res = await fetch("http://jpoly1219devbox.xyz:8090/auth/silentrefresh", options)
			const accessToken = await res.json()
			const payloadB64 = accessToken.split(".")[1]
			timeToExpire.set(JSON.parse(window.atob(payloadB64)).exp)
		} catch(err) {
			alert(err)
		}
	}

    function refreshTimer() {
		if ($timeToExpire != "") {
			var i = Date.now()/1000;
			var timer = setInterval(() => {
				//console.log($timeToExpire)
				if (i >= Number($timeToExpire)) {
                    refresh()
                    clearInterval(timer)
				}	
				//console.log(Date.now()/1000)
				i++
			}, 1000)
		}
	}

	$: $timeToExpire, refreshTimer()
</script>

<div class="p-8 h-screen">
    <Nav/>
    <slot></slot>
</div>

<style>
    @tailwind base;
    @tailwind components;
    @tailwind utilities;
</style>