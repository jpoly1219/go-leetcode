<script>
    import { beforeUpdate, createEventDispatcher } from "svelte";
    import { accessTokenStore } from "../stores/stores.js"

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Discussioncard'
        })
    }

    let username = ""
    beforeUpdate(() => {
        if ($accessTokenStore != "") {
            const payloadB64 = $accessTokenStore.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }
    })

    let title = ""
    let description = ""

    const postDiscussion = async () => {
        console.log("running postDiscussion")
        const url = `http://jpoly1219devbox.xyz:8090/discussions/${discussion.slug}/newdiscussion`
        const newDiscussionData = {
            author: username,
            title: title,
            description: description
        }
        const options = {
            method: "POST",
            body: JSON.stringify(newDiscussionData)
        }
        const res = await fetch(url, options)
        const data = await res.json()
        console.log("new discussion posted: ", data)
    }
</script>

<div class="overflow-auto">
    <div class="flex flex-row divide-x divide-solid divide-gray-300 items-center ">
        <p on:click={switchComponent} class="text-sm mr-3 my-2 cursor-pointer">&lt; Back</p>
        <p class="font-bold text-lg px-3 my-2">Create New Discussion</p>
    </div>
    <div class="mb-5">
        <textarea bind:value={title} class="w-full h-9 border border-gray-200 rounded p-1 mb-5" placeholder="Title"></textarea>
        <textarea bind:value={description} class="w-full h-96 border border-gray-200 rounded p-1 mb-5" placeholder="Description (Markdown is supported)"></textarea>
        <span on:click={postDiscussion} class="bg-gray-600 px-2 py-1 mt-3 text-sm text-white cursor-pointer rounded">Post</span>
    </div>
</div>