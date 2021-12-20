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

    // Menu bar
    let filterObject = {
        difficulty: "all"
    }

    async function filter() {
        // fetch API to return only the problems of the selected difficulty
        const url = "http://jpoly1219devbox.xyz:8090/problemsets/filter"
        const options = {
            method: "POST",
            body: JSON.stringify(filterObject)
        }

        try {
            const res = await fetch(url, options)
            const data = await res.json()
            problems = data
            console.log(problems)
        } catch(err) {
            console.log(err)
        }
    }
</script>

<svelte:head>
    <title>Problems - go-leetcode</title>
</svelte:head>

<h1 class="text-4xl text-center my-8">Problem Sets</h1>
<div class="flex flex-row mb-4">
    <div class="flex flex-row items-center">
        <p class="text-base mr-2">Difficulty:</p>
        <select bind:value={filterObject.difficulty} on:change={filter} class="border border-gray-300 rounded-lg">
            <option value="all">
                <p class="text-green-500">all</p>
            </option>
            <option value="easy">
                <p class="text-green-500">easy</p>
            </option>
            <option value="medium">
                <p class="text-yellow-500">medium</p>
            </option>
            <option value="hard">
                <p class="text-red-500">hard</p>
            </option>
        </select>
    </div>
</div>
<div class="w-full">
    <table class="table-fixed w-full items-start">
        <tr class="border-b border-gray-200">
            <th class="w-1/12 text-left text-lg">Status</th>
            <th class="w-8/12 text-left text-lg">Title</th>
            <th class="w-3/12 text-left text-lg">Difficulty</th>
        </tr>
        {#each problems as problem}
            <tr class="{problem.num % 2 === 0 ? "bg-gray-100" : "bg-white"} py-4">
                <td>
                    <p class="text-base">Status</p>
                </td>
                <td>
                    <p class="text-base"><a href={`/solve/${problem.slug}/description`}>{problem.title}</a></p>
                </td>
                <td>
                    <p class="text-base font-extralight my-2 {problem.difficulty === "easy" ? "text-green-500" : (problem.difficulty === "medium" ? "text-yellow-500" : "text-red-500")}">
                        {problem.difficulty}
                    </p>
                </td>
            </tr>
        {/each}
    </table>
</div>

<!--
    {#each problems as problem}
        <div class="mb-5">
            <p class="text-lg my-2"><a href={`/solve/${problem.slug}/description`}>{problem.title}</a></p>
            <p class="text-sm font-extralight my-2 {problem.difficulty === "easy" ? "text-green-500" : (problem.difficulty === "medium" ? "text-yellow-500" : "text-red-500")}">
                {problem.difficulty}
            </p>
        </div>
    {/each}
-->