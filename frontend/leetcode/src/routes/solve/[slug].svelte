<script context="module">
    import { get } from "svelte/store"
    import { accessTokenStore } from "../../stores/stores"
    export async function load({page}) {
        const slug = page.params.slug
        const url = `http://jpoly1219devbox.xyz:8090/solve/${slug}`
        let accessToken = get(accessTokenStore)
        const options = {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + accessToken,
            },
            credentials: "include"
        }
        // console.log(`now fetching:\n ${options.headers.Authorization}`)
        try {
            const res = await fetch(url, options)
            const problem = await res.json()
            return {props: {problem}}
        } catch(err) {
            console.log(err)
        }
    }
</script>

<script>
    import {onMount} from "svelte"
    import snarkdown from "snarkdown"
    
    export let problem
    
    let CodeJar;
    onMount(async () => {
        ({CodeJar} = await import("svelte-codejar"));
    });
    export let value = `console.log("Hello World!");`

    let languages = ["C++", "Java", "Javascript", "Python"]
    let selected

    async function submit() {
        const userInput = {
            lang: selected,
            code: value
        }

        const options = {
            method: "POST",
            body: JSON.stringify(userInput)
        }
        const res = await fetch(`http://jpoly1219devbox.xyz:8090/check/${problem.slug}`, options)
        const data = await res.json()
        console.log(data)
        alert("code submitted!")
        alert(`Result: ${data.result}:\nInput: ${data.input}\nExpected: ${data.expected}\nOutput: ${data.output}`)
    }
</script>

<svelte:head>
    <title>{problem.title} - go-leetcode</title>
</svelte:head>

<div class="grid grid-rows-16 h-full">
    <div class="row-span-15 grid grid-cols-2 gap-4">
        <div class="overflow-auto border border-gray-300 p-4">
            <p class="font-bold">{problem.title}</p>
            <p class="text-sm text-green-600 font-light mt-2">{problem.difficulty}</p>
            <hr class="my-4">
            <p class="prose max-w-max">{@html snarkdown(problem.description)}</p>
        </div>
        <div class="border border-gray-300 overflow-auto">
            {#if CodeJar}
            <CodeJar addClosing={true} indentOn={/{$/} spellcheck={false} tab={"\t"} withLineNumbers={true} bind:value/>
            {:else}
            <pre><code>{value}</code></pre>
            {/if}
        </div>
    </div>
    <div class="row-span-1 grid grid-cols-2 gap-4 content-center">
        <div class="flex flex-row">
            <div class="flex-1 flex">
                <button class="border border-gray-300 rounded-lg px-3 py-2">Problems</button>
            </div>
            <button class="border border-gray-300 rounded-lg px-3 py-2">Pick One</button>
            <button class="border border-gray-300 rounded-lg px-3 py-2 mx-4">Prev</button>
            <div class="mx-2 flex">
                <p class="self-center">1/1977</p>
            </div>
            <button class="border border-gray-300 rounded-lg px-3 py-2 ml-4">Next</button>
        </div>
        <div class="flex flex-row">
            <div class="flex-1 flex">
                <button class="border border-gray-300 rounded-lg px-3 py-2">Console</button>
            </div>
            <select bind:value={selected} class="border border-gray-300 rounded-lg px-3 py-2 mx-2">
                {#each languages as language}
                <option value={language}>{language}</option>
                {/each}
            </select>
            <button class="border border-gray-300 rounded-lg px-3 py-2 mx-2">Run Code</button>
            <button on:click={submit} class="border border-gray-300 rounded-lg px-3 py-2 ml-2">Submit</button>
        </div>
    </div>
</div>
