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
</script>

<svelte:head>Profile - go-leetcode</svelte:head>

<div class="container">
    <p>This is the profile page of the user {user}.</p>
    <p>Data to be added soon.</p>
</div>