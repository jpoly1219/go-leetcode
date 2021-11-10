<script>
    import { createEventDispatcher } from "svelte";

    export let discussion

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Singlediscussion',
            props: {
                discussion: discussion
            }
        })
    }
</script>

<p on:click={switchComponent}>{discussion.title}</p>