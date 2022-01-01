<script context="module">
    import { get } from "svelte/store"
    import { accessTokenStore } from "../stores/stores.js"

    export async function load() {
        let accessToken = get(accessTokenStore)
        let username
        if (accessToken != "") {
            const payloadB64 = accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }

        const url = "http://jpoly1219devbox.xyz:8090/problemsets/all"
        const options = {
            method: "POST",
            body: JSON.stringify({username: username})
        }

        try {
            const res = await fetch(url, options)
            const data = await res.json()
            console.log(data, typeof(data))
            const loadedProblem = data.map((data, index) => {
                return {
                    num: index + 1,
                    title: data.title,
                    slug: data.slug,
                    difficulty: data.difficulty,
                    result: data.result
                }
            })
            return {props: {problems: loadedProblem}, username}
        } catch(err) {
            alert(err)
        }
    }

    // Call attempts table. Send username to API.
    // Check username, slug and result column.
    // SELECT DISTINCT title, problems.slug, difficulty, result FROM problems LEFT JOIN attempts ON problems.slug = attempts.slug AND userame = $1 AND result = 'OK' ORDER BY title;
    // Get back slug/result JSON.
    // for each loadedProblem, if result = OK, mark the problem as solved.
</script>

<script>
    import { onMount } from "svelte"
    import { problemsListStore } from "../stores/stores.js"

    export let problems
    export let username
    
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
        username: username,
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
            const loadedProblem = data.map((data, index) => {
                return {
                    num: index + 1,
                    title: data.title,
                    slug: data.slug,
                    difficulty: data.difficulty,
                    result: data.result
                }
            })

            problems = loadedProblem
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
                    <p class="text-base">{problem.result}</p>
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