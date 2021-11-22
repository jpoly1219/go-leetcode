<script>
    import { accessTokenStore } from "../stores/stores";
    import { beforeUpdate } from "svelte";
    import { goto } from "$app/navigation"

    let username = ""
    beforeUpdate(() => {
        if ($accessTokenStore != "") {
            const payloadB64 = $accessTokenStore.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
            goto("/")
        }
    })

    let currentPage = "home"
</script>
<nav class="flex justify-start w-full bg-gray-800 py-2 items-center sticky top-0">
    <p class="mx-4 px-3 py-2 rounded-md text-lg text-white font-extrabold">go-leetcode</p>
    <a href="/" on:click={() => {currentPage = "home"}} class="mx-4 px-3 py-2 rounded-md text-lg text-white font-medium {currentPage === "home" ? "bg-gray-900" : "hover:bg-gray-700"}">Home</a>
    <a href="/problemset" on:click={() => {currentPage = "problems"}} class="mx-4 px-3 py-2 rounded-md text-lg text-white font-medium {currentPage === "problems" ? "bg-gray-900" : "hover:bg-gray-700"}">Problems</a>
    {#if $accessTokenStore != ""}
        <a href="#" on:click={() => {currentPage = "profile"}} class="mx-4 text-lg text-white {currentPage === "profile" ? "bg-gray-900" : "hover:bg-gray-700"}">Welcome, {username}</a>
    {:else}
    <div class="absolute right-0">
        <a href="/signup" on:click={() => {currentPage = "signup"}} class="mx-4 px-3 py-2 rounded-md text-lg text-white font-medium {currentPage === "signup" ? "bg-gray-900" : "hover:bg-gray-700"}">Sign Up</a>
        <a href="/login" on:click={() => {currentPage = "login"}} class="mx-4 px-3 py-2 rounded-md text-lg text-white font-medium {currentPage === "login" ? "bg-gray-900" : "hover:bg-gray-700"}">Log In</a>
    </div>
    {/if}
</nav>