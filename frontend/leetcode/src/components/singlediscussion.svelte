<script>
    import { createEventDispatcher } from "svelte";

    export let discussion

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Discussioncard'
        })
    }
</script>

<div>
    <div class="flex flex-row divide-x divide-solid divide-gray-300">
        <p on:click={switchComponent} class="text-sm mx-3 my-2">&lt; Back</p>
        <p class="font-bold text-lg mx-3 my-2">{discussion.title}</p>
    </div>
    <div>
        <p class="text-sm mb-2">{discussion.author}</p>
        <p class="prose max-w-max">{@html snarkdown(discussion.description)}</p>
    </div>
</div>
