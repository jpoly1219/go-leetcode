<script>
    import Discussioncard from "./discussioncard.svelte"
    import Singlediscussion from "./singlediscussion.svelte"

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

{#if currentComponent === "Singlediscussion"}
    <Singlediscussion discussion={props.discussion} on:switch={handleSwitch}/>
{:else}
    <p class="text-lg font-bold mb-3">Discussion Board for {slug}</p>
    <div class="divide-y divide-solid divide-gray-300">
    {#each discussions as discussion}
        <Discussioncard discussion={discussion} on:switch={handleSwitch}/>
    {/each}
    </div>
{/if}
