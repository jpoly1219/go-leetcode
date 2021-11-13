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

<div>
    <div class="flex flex-row items-center py-2">
        <img src="https://www.catholicsingles.com/wp-content/uploads/2020/06/blog-header-3.png" class="rounded-full w-12 h-12 mr-4">
        <div class="flex flex-col">
            <p on:click={switchComponent} class="text-base">{discussion.title}</p>
            <p class="text-xs text-gray-500">{discussion.author} created at: {discussion.created}</p>
        </div>
    </div>
</div>
