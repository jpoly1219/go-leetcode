<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Discussioncard'
        })
    }

    let title = ""
    let description = ""

    const postDiscussion = () => {

    }
</script>

<div class="overflow-auto">
    <div class="flex flex-row divide-x divide-solid divide-gray-300 items-center">
        <p on:click={switchComponent} class="text-sm mr-3 my-2 cursor-pointer">&lt; Back</p>
        <p class="font-bold text-lg px-3 my-2">Create New Discussion</p>
    </div>
    <div class="mb-5">
        <textarea bind:value={title} class="w-full h-8 border border-gray-200 rounded" placeholder="Title"></textarea>
        <textarea bind:value={description} class="w-full h-150 border border-gray-200 rounded" placeholder="Description (Markdown is supported)"></textarea>
        <span on:click={postDiscussion} class="bg-gray-600 px-2 py-1 mt-3 text-sm text-white cursor-pointer rounded">Post</span>
    </div>
</div>