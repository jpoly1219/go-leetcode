<script>
    import Discussioncard from "./discussioncard.svelte"
    import Singlediscussion from "./singlediscussion.svelte"
    import Newdiscussion from "./newdiscussion.svelte";

    // props
    export let slug
    export let discussions

    // variable to show which component to render
    let currentComponent
    let props

    // event listener to change currentView
    const handleSwitch = (event) => {
        currentComponent = event.detail.component
        props = event.detail.props
    }
</script>

<div class="container w-full">
    {#if currentComponent === "Singlediscussion"}
        <Singlediscussion discussion={props.discussion} on:switch={handleSwitch}/>
    {:else if currentComponent === "Newdiscussion"}
        <Newdiscussion on:switch={handleSwitch}/>
    {:else}
        <div class="flex flex-row items-center relative">
            <p class="text-lg font-bold mb-3">Discussion Board for {slug}</p>
            <a href="#">
                <button on:click={currentComponent = "Newdiscussion"} class="absolute right-0 rounded-md bg-gray-800 px-2 py-1 text-sm text-white">New +</button>
            </a>
        </div>
        <div class="divide-y divide-solid divide-gray-300">
        {#each discussions as discussion}
            <Discussioncard discussion={discussion} on:switch={handleSwitch}/>
        {/each}
        </div>
    {/if}
</div>
