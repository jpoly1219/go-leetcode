<script>
    import { createEventDispatcher } from 'svelte'
    
    const dispatch = createEventDispatcher()

    export let tabs
    export let activeTab
</script>

<div>
    <ul>
        {#each tabs as tab}
        <li on:click={() => dispatch("tabChange", tab)}>
            <div>{tab}</div>
        </li>
        {/each}
    </ul>
</div>