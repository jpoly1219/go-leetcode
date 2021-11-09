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

<p on:click={switchComponent}>Title</p>