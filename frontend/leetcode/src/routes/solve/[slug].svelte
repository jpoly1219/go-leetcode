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
    <div class="row-span-1 border border-gray-300 rounded-lg">
        <div class="grid grid-cols-2">
            <div class="flex flex-row">
                <button>Problems</button>
                <button>Pick One</button>
                <button>Prev</button>
                <p>1/1977</p>
                <button>Next</button>
            </div>
            <div class="flex flex-row">
                <button>Console</button>
                <button>Run Code</button>
                <button>Submit</button>
            </div>
        </div>
    </div>
</div>
