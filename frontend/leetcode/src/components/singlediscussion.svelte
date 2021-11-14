<script>
    import { createEventDispatcher, onMount } from "svelte";
    import snarkdown from "snarkdown"

    export let discussion

    const dispatch = createEventDispatcher();
    const switchComponent = () => {
        dispatch('switch', {
            component: 'Discussioncard'
        })
    }

    let comments = []
    onMount(async () => {
        const url = `http://jpoly1219devbox.xyz:8090/discussions/${discussion.slug}/${discussion.id}`
        const options = {
            method: "GET"
        }
        const res = await fetch(url, options)
        const data = await res.json()
        comments = data.map((data) => {
            return {
                id: data.id,
                author: data.author,
                discussionId: data.discussionId,
                description: data.description,
                created: data.created
            }
        })
    })

    const postComment = () => {
        console.log("running postComment")
    }
</script>

<div class="overflow-auto">
    <div class="flex flex-row divide-x divide-solid divide-gray-300 items-center">
        <p on:click={switchComponent} class="text-sm mr-3 my-2 cursor-pointer">&lt; Back</p>
        <p class="font-bold text-lg px-3 my-2">{discussion.title}</p>
    </div>
    <div>
        <p class="text-sm mb-2">{discussion.author}</p>
        <p class="prose max-w-max">{@html snarkdown(discussion.description)}</p>
    </div>
    <div class="my-5 border-t border-b border-gray-200">
        <p class="text-base my-2">Comments:</p>
    </div>
    <div class="mb-5 border border-gray-200 rounded">
        <textarea class="w-full h-24" placeholder="Type comment here... (Markdown is supported)"></textarea>
        <span on:click={postComment}>Post</span>
    </div>
    {#each comments as comment}
        <div class="">
            <div class="">
                <div class="flex flex-row items-center mb-2">
                    <img 
                        src="https://img.freepik.com/free-photo/pleasant-looking-serious-man-stands-profile-has-confident-expression-wears-casual-white-t-shirt_273609-16959.jpg?size=626&ext=jpg"
                        class="rounded-full w-10 h-10 mr-2"
                    >
                    <p class="text-sm mr-2">{comment.author}</p>
                    <p class="text-sm">{comment.created}</p>
                </div>
                <div class="flex flex-row">
                    <div class="w-10 mr-2"></div>
                    <p>{comment.description}</p>
                </div>
            </div>
        </div>
    {/each}
</div>
