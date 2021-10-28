<script context="module">
    import { get } from "svelte/store"
    import { accessTokenStore } from "../../stores/stores"
    export async function load({page}) {
        const slug = page.params.slug
        const url1 = `http://jpoly1219devbox.xyz:8090/solve/${slug}`
        const url2 = `http://jpoly1219devbox.xyz:8090/submissions`

        let accessToken = get(accessTokenStore)
        let username
        if (accessToken != "") {
            const payloadB64 = accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }
        const options1 = {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + accessToken,
            },
            credentials: "include"
        }
        const options2 = {
            method: "POST",
            body: JSON.stringify({username: username, slug: slug})
        }
        try {
            const res1 = await fetch(url1, options1)
            const res2 = await fetch(url2, options2)
            const problem = await res1.json()
            const submissions = await res2.json()
            return {props: {problem, submissions, username}}
        } catch(err) {
            console.log(err)
        }
    }
</script>

<script>
    import { onMount } from "svelte"
    import snarkdown from "snarkdown"
    import Tabs from "../../components/tabs.svelte";
    
    export let problem
    export let submissions
    export let username
    
    let CodeJar
    let submissionsData = []
    onMount(async () => {
        ({CodeJar} = await import("@novacbn/svelte-codejar"));
        submissions.map((data) => {
            submissionsData.push(data)
        })
    });
    export let value = ""

    let languages = ["cpp", "java", "js", "py"]
    let selected

    let resultData
    async function submit() {
        alert("code submitted!")
        activeTab = "Submissions"
        const userInput = {
            username: username,
            slug: problem.slug,
            lang: selected,
            code: value
        }

        const options = {
            method: "POST",
            body: JSON.stringify(userInput)
        }
        const res = await fetch(`http://jpoly1219devbox.xyz:8090/check/${problem.slug}`, options)
        resultData = await res.json()
        console.log(resultData)
        submissionsData.push(resultData)
    }

    let tabs = ["Description", "Solution", "Discussion", "Submissions"]
    let activeTab = "Description"
    const tabChange = (e) => {
        activeTab = e.detail
        
    }
</script>

<svelte:head>
    <title>{problem.title} - go-leetcode</title>
</svelte:head>

<div class="grid grid-rows-16 h-full">
    <div class="row-span-15 grid grid-cols-2 gap-4">
        <div class="overflow-auto border border-gray-300 p-4">
            <Tabs {tabs} {activeTab} on:tabChange={tabChange} />
            {#if activeTab === "Description"}
            <p class="text-lg font-bold">{problem.title}</p>
            <p class="text-sm text-green-600 font-light mt-2">{problem.difficulty}</p>
            <hr class="my-4">
            <p class="prose max-w-max">{@html snarkdown(problem.description)}</p>
            {:else if activeTab === "Solution"}
            <p class="text-lg font-bold">Solution</p>
            {:else if activeTab === "Discussion"}
            <p class="text-lg font-bold">Discussion</p>
            {:else if activeTab === "Submissions"}
            <p class="text-lg font-bold">Submissions</p>
                {#if submissionsData}
                <div class="w-full">
                    <table class="items-center w-full border-collapse">
                        <thead>
                            <tr>
                                <th class="px-4 py-2 bg-gray-300 border border-solid border-gray-400 border-r-0 text-sm text-gray-700">Result</th>
                                <th class="px-4 py-2 bg-gray-300 border border-solid border-gray-400 border-l-0 text-sm text-gray-700">Output</th>
                            </tr>
                        </thead>
                        {#each submissionsData as submissionsDatum}
                        <tbody>
                            <tr>
                                <td class="text-red-600" class:text-green-600="{submissionsDatum.result === 'OK'}">{submissionsDatum.result}</td>
                                <td class="break-words">{submissionsDatum.output}</td>
                            </tr>
                        </tbody>
                        {/each}
                    </table>
                </div>
                {/if}
            {/if}
        </div>
        <div class="flex flex-col border border-gray-300 overflow-hidden">
            <div class="overflow-auto">
                {#if CodeJar}
                <CodeJar addClosing={true} indentOn={/{$/} spellcheck={false} tab={"\t"} withLineNumbers={true} bind:value/>
                {:else}
                <pre><code>{value}</code></pre>
                {/if}
            </div>
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
                <option value={languages[0]}>C++</option>
                <option value={languages[1]}>Java</option>
                <option value={languages[2]}>Javascript</option>
                <option value={languages[3]}>Python</option>
            </select>
            <button class="border border-gray-300 rounded-lg px-3 py-2 mx-2">Run Code</button>
            <button on:click={submit} class="border border-gray-300 rounded-lg px-3 py-2 ml-2">Submit</button>
        </div>
    </div>
</div>
