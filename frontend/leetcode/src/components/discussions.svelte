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

<p class="text-lg font-bold mb-3">Discussion Board for {slug}</p>

{#if currentComponent === "Singlediscussion"}
    <Singlediscussion discussion={props.discussion} on:switch={handleSwitch}/>
{:else}
    {#each discussions as discussion}
        <Discussioncard discussion={discussion} on:switch={handleSwitch}/>
    {/each}
{/if}
