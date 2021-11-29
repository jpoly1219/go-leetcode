<script>
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Newdiscussion'
        })
    }

    
</script>

<p>create new discussion here</p>