<script>
    let username = ""
    let password = ""

    async function submit() {
        const loginDetails = {
            username: username,
            password: password
        }

        const options = {
            method: "POST",
            body: JSON.stringify(loginDetails),
            headers: {
                "Content-Type": "application/json"
            },
            credentials: "include"
        }
        const url = "http://jpoly1219devbox.xyz:8090/auth/login"
        const res = await fetch(url, options)
        const tokenPair = await res.json()
    }
</script>

<svelte:head>
    <title>Log In - go-leetcode</title>
</svelte:head>

<div class="container mx-auto items-center flex flex-col text-gray-900">
    <div class="bg-white p-16 rounded-lg border-gray-900 shadow-lg my-44">
        <h2 class="text-2xl font-medium text-center mb-20">Welcome back.</h2>
        <form on:submit|preventDefault={submit} class="flex flex-col px-30">
            <div class="flex flex-col items-stretch px-30">
                <input bind:value={username} type="text" placeholder="Username" class="border-b-2 py-4">
                <input bind:value={password} type="password" placeholder="Password" class="border-b-2 py-4">
            </div>
            <div class="flex flex-col items-center my-10">
                <button type="submit" class="border border-gray-500 rounded-lg p-3">
                    <span class="mx-3 my-2">Log in</span>
                </button>
            </div>
            <p class="text-sm font-extralight">
                Don't have an account yet? <a href="/signup" class="text-blue-400 cursor-pointer">Sign up</a> today.
            </p>
        </form>
    </div>
</div>