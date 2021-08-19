<script>
    import { problems } from "../stores/problemstore"
    import snarkdown from "snarkdown"
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
<div class="grid grid-cols-2 gap-6">
    <div class="border border-gray-300 rounded-lg">
        {#each $problems as problem}
        <p class="text-lg my-2">{problem.title}</p>
        <p class="text-sm font-extralight text-gray-500 my-2">{problem.difficulty}</p>
        <p class="text-sm font-light">{@html snarkdown(problem.description)}</p>
        {/each}
    </div>
    <div class="border border-gray-300 rounded-lg">
        <p>Your code here!</p>
    </div>
</div>