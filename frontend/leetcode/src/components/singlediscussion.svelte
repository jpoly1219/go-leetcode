<script>
    import { createEventDispatcher } from "svelte";
    import snarkdown from "snarkdown"

    export let discussion

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Discussioncard'
        })
    }
</script>

<div class="overflow-auto">
    <div class="flex flex-row divide-x divide-solid divide-gray-300">
        <p on:click={switchComponent} class="text-sm mx-3 my-2">&lt; Back</p>
        <p class="font-bold text-lg mx-3 my-2">{discussion.title}</p>
    </div>
    <div>
        <p class="text-sm mb-2">{discussion.author}</p>
        <p class="prose max-w-max">{@html snarkdown(discussion.description)}</p>
    </div>
    <div class="my-2 border-t-2 border-b-2 border-gray-400">
        <p class="text-base">Comments:</p>
    </div>
    <div class="">
        <div class="">
            <div class="flex flex-row">
                <img 
                    src="https://img.freepik.com/free-photo/pleasant-looking-serious-man-stands-profile-has-confident-expression-wears-casual-white-t-shirt_273609-16959.jpg?size=626&ext=jpg"
                    class="rounded-full w-12 h-12 mr-2"
                >
                <p class="text-sm mr-2">username</p>
                <p class="text-sm">time created</p>
            </div>
            <div class="flex flex-row">
                <div class="w-12 mr-2"></div>
                <p>This is a comment.</p>
            </div>
        </div>
    </div>
</div>
