<script context="module">
    import { get } from "svelte/store"
    import { accessTokenStore } from "../stores/stores.js"

    export async function load() {
        let accessToken = get(accessTokenStore)
        let username
        if (accessToken != "") {
            const payloadB64 = accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }

        const url = `http://localhost:8090/users/${username}`
        const options = {
            method: "GET"
        }

        try {
            const res = await fetch(url, options)
            const userData = await res.json()

            return {props: {user: userData}}
        } catch (err) {
            console.log(err)
        }
    }

</script>

<script>
    export let user
    let username = ""
    let fullname = ""
    let email = ""
    let password = ""

    async function submit() {
        const editProfileDetails = {
            username: username,
            fullname: fullname,
            email: email,
            password: password
        }

        const options = {
            method: "POST",
            body: JSON.stringify(editProfileDetails),
            headers: {
                "Content-Type": "application/json"
            }
        }
        const url = "http://localhost:8090/auth/editprofile"

        try {
            const res = await fetch(url, options)
            const message = await res.json()
            alert(message.message)
        } catch(err) {
            alert(err)
        }
    }
</script>

<svelte:head>Profile - go-leetcode</svelte:head>

<div class="container items-center flex flex-col">
    <div class="rounded-lg shadow-lg w-1/3 my-44">
        <div class="m-4">
            <p class="text-2xl font-medium text-center">Welcome, {user.fullname}</p>
        </div>
        <div class="grid grid-cols-2 grid-rows-4">
            <div class="m-4">
                <p>Username</p>
            </div>
            <div class="m-4">
                <input bind:value={username} type="text" placeholder={user.username}>
            </div>
            <div class="m-4">
                <p>Full Name</p>
            </div>
            <div class="m-4">
                <input bind:value={fullname} type="text" placeholder={user.fullname}>
            </div>
            <div class="m-4">
                <p>Email</p>
            </div>
            <div class="m-4">
                <input bind:value={email} type="text" placeholder={user.email}>
            </div>
            <div class="m-4">
                <p>New Password</p>
            </div>
            <div class="m-4">
                <input bind:value={password} type="text">
            </div>
            <div class="m-4 flex justify-center col-span-2">
                <button type="submit" class="border bg-blue-400 rounded-lg p-3">
                    <a href="/" class="mx-3 my-2 text-white">Save Changes</a>
                </button>
            </div>
        </div>
    </div>
</div>

<!-- 
<div class="container mx-auto w-1/3 shadow">
    <div class="bg-white rounded-t-lg">
        <div class="px-7 py-5">
            <h3 class="text-lg leading-6 font-medium text-gray-900">
                User Profile
            </h3>
            <p class="mt-1 text-sm text-gray-500">
                Edit your profile and account information here.
            </p>
        </div>
    </div>
    <div class="border-t border-gray-200">
        <form action="submit">
            <div class="bg-gray-50 px-7 py-5 flex flex-row justify-between">
                <h4 class="text-base text-gray-500 font-medium mr-52 self-center">
                    Full Name
                </h4>
                <input type="text" value={user.fullname} class="text-base py-2 flex-grow max-w-sm">
            </div>
            <div class="bg-white px-7 py-5 flex flex-row justify-between">
                <h4 class="text-base text-gray-500 font-medium mr-52 self-center">
                    Username
                </h4>
                <input type="text" value={user.username} class="text-base py-2 flex-grow max-w-sm">
            </div>
            <div class="bg-gray-50 px-7 py-5 flex flex-row justify-between">
                <h4 class="text-base text-gray-500 font-medium mr-52 self-center">
                    Email
                </h4>
                <input type="password" value={user.email} class="text-base py-2 flex-grow max-w-sm">
            </div>
            <div class="bg-white px-7 py-5 flex flex-row justify-between">
                <h4 class="text-base text-gray-500 font-medium mr-52 self-center">
                    Password
                </h4>
                <input type="text" class="text-base py-2 flex-grow max-w-sm">
            </div>
            <div class="bg-gray-50 px-7 py-5 flex flex-row justify-end">
                <button type="submit" class="inline-flex items-center bg-blue-400 border rounded-lg text-base">
                    <span class="mx-3 my-2 text-white">Save changes</span>
                </button>
                <button type="button" class="inline-flex items-center ml-5 border border-gray-500 rounded-lg">
                    <span class="mx-3 my-2">Return to home</span>
                </button>
            </div>
        </form>
    </div>
</div>
-->
