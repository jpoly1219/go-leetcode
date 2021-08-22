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

<div class="grid grid-rows-16 h-full">
    <div class="row-span-15 grid grid-cols-2">
        <div class="prose border border-gray-300 rounded-lg max-w-max overflow-auto">
            <p class="">{problem.title}</p>
            <p class="">{problem.difficulty}</p>
            <p class="">{@html snarkdown(problem.description)}</p>
        </div>
        <div class="border border-gray-300 rounded-lg">
            <p>Your code here!</p>
        </div>
    </div>
    <div class="row-span-1 grid grid-cols-2 content-center border border-gray-300 rounded-lg">
        <div class="flex flex-row">
            <div class="flex-1 flex">
                <button class="border border-gray-300 rounded-lg px-3 py-2">Problems</button>
            </div>
            <button class="border border-gray-300 rounded-lg px-3 py-2">Pick One</button>
            <button class="border border-gray-300 rounded-lg px-3 py-2 mx-4">Prev</button>
            <div class="mx-2 flex">
                <p class="self-center">1/1977</p>
            </div>
            <button class="border border-gray-300 rounded-lg px-3 py-2 mx-4">Next</button>
        </div>
        <div class="flex flex-row">
            <button>Console</button>
            <button>Run Code</button>
            <button>Submit</button>
        </div>
    </div>
</div>
