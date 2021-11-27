<script context="module">
    export async function load() {
        const url = "http://jpoly1219devbox.xyz:8090/problemsets"
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
    }
</script>

<script>
    import { problemsListStore } from "../stores/stores.js"

    export let problems
    
    console.log($problemsListStore)
    Object.entries(problems).forEach(([key, value]) => {
        $problemsListStore = [...$problemsListStore, value.slug]
    })
    console.log($problemsListStore)
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
{#each problems as problem}
<p class="text-lg my-2"><a href={`/solve/${problem.slug}`}>{problem.title}</a></p>
<p class="text-sm font-extralight text-gray-500 my-2">{problem.difficulty}</p>
{/each}