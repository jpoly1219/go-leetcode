<script>
    import { createEventDispatcher } from 'svelte'

    const dispatch = createEventDispatcher()

    export let tabs
    export let activeTab
</script>

<div class="mb-3">
    <ul class="flex justify-start p-0 list-none">
        {#each tabs as tab}
        <li on:click={() => dispatch("tabChange", tab)} class="mr-4 text-sm cursor-pointer">
            <div>{tab}</div>
        </li>
        {/each}
    </ul>
</div>