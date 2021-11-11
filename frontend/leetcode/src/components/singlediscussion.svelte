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
    <div class="flex flex-row divide-x divide-solid divide-gray-100">
        <p on:click={switchComponent} class="text-sm px-3 py-2">&gt; Back</p>
        <p class="font-bold text-lg">{discussion.title}</p>
    </div>
    <div>
        <p class="text-sm">{discussion.author}</p>
        <p class="text-base">{discussion.description}</p>
    </div>
</div>
