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

<p on:click={switchComponent}>{discussion.title}</p>
<p>{discussion.author}</p>
<p>{discussion.description}</p>