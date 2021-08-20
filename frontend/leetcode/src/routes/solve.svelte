<script>
    import { problems } from "../stores/problemstore"
    import snarkdown from "snarkdown"
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<div class="h-screen">
    <h1 class="text-4xl text-center my-8">Problem Sets</h1>
    <div class="grid grid-cols-2 gap-6">
        <div class="prose border border-gray-300 rounded-lg max-w-max overflow-auto">
            {#each $problems as problem}
            <p class="">{problem.title}</p>
            <p class="">{problem.difficulty}</p>
            <p class="">{@html snarkdown(problem.description)}</p>
            {/each}
        </div>
        <div class="border border-gray-300 rounded-lg">
            <p>Your code here!</p>
        </div>
    </div>
</div>