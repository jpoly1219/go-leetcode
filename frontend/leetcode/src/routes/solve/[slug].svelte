<script context="module">
    export async function load({page}) {
        const slug = page.params.slug
        const url = `http://jpoly1219devbox.xyz:8090/solve/${slug}`
        const res = await fetch(url)
        const problem = await res.json()
        return {props: {problem}}
    }
</script>

<script>
    import snarkdown from "snarkdown"
    export let problem
</script>

<svelte:head>
    <title>{problem.title} - go-leetcode</title>
</svelte:head>

<div class="grid grid-cols-2 gap-6 h-full">
    <div class="prose border border-gray-300 rounded-lg max-w-max overflow-auto">
        <p class="">{problem.title}</p>
        <p class="">{problem.difficulty}</p>
        <p class="">{@html snarkdown(problem.description)}</p>
    </div>
    <div class="border border-gray-300 rounded-lg">
        <p>Your code here!</p>
    </div>
</div>
