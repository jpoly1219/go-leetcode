<script>
    import { createEventDispatcher } from "svelte";

    export let discussion

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Singlediscussion',
            props: {
                title: discussion.title,
                author: discussion.author,
                description: discussion.description,
                created: discussion.created
            }
        })
    }
</script>

<p on:click={switchComponent}>{discussion.title}</p>