<script>
    import { accessTokenStore } from "../stores/stores";
    import { beforeUpdate } from "svelte";

    let username = ""
    beforeUpdate(() => {
        if ($accessTokenStore != "") {
            const payloadB64 = $accessTokenStore.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }
    })
</script>
<nav class="flex justify-center w-full">
    <a href="/" class="mx-4 text-lg">Home</a>
    <a href="/problemset" class="mx-4 text-lg">Problems</a>
    {#if $accessTokenStore != ""}
        <a href="#" class="mx-4 text-lg">Welcome, {username}</a>
    {:else}
        <a href="/signup" class="mx-4 text-lg">Sign Up</a>
        <a href="/login" class="mx-4 text-lg">Log In</a>
    {/if}
</nav>