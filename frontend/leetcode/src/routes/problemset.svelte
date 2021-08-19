<script>
    import { problems } from "../stores/problemstore"
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
{#each $problems as problem}
<p class="text-lg my-2"><a href="/solve">{problem.title}</a></p>
<p class="text-sm font-extralight text-gray-500 my-2">{problem.difficulty}</p>
{/each}