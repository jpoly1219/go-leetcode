<script>
    import { problems } from "../stores/problemstore"
    console.log($problems)
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
{#each $problems as problem}
    <p>{problem.title}</p>
{/each}