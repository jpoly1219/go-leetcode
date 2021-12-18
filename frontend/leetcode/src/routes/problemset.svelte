<script context="module">
    export async function load() {
        const url = "http://jpoly1219devbox.xyz:8090/problemsets"

        try {
            const res = await fetch(url)
            const data = await res.json()
            console.log(data, typeof(data))
            const loadedProblem = data.map((data, index) => {
                return {
                    num: index + 1,
                    title: data.title,
                    slug: data.slug,
                    difficulty: data.difficulty,
                    description: data.description,
                    created: data.created
                }
            })
            return {props: {problems: loadedProblem}}
        } catch(err) {
            alert(err)
        }
    }
</script>

<script>
    import { onMount } from "svelte"
    import { problemsListStore } from "../stores/stores.js"

    export let problems
    
    onMount(() => {
        if ($problemsListStore.length != 0) {
            problemsListStore.set([])
        }
        Object.entries(problems).forEach(([key, value]) => {
            $problemsListStore = [...$problemsListStore, value.slug]
        })
    })
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
{#each problems as problem}
    <p class="text-lg my-2"><a href={`/solve/${problem.slug}/description`}>{problem.title}</a></p>
    <p class="text-sm font-extralight my-2 {problem.difficulty === "easy" ? "text-green-500" : (problem.difficulty === "medium" ? "text-yellow-500" : "text-red-500")}">
        {problem.difficulty}
    </p>
{/each}