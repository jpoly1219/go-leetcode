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
    import { beforeUpdate, onMount } from "svelte"
    import snarkdown from "snarkdown"
    import Tabs from "../../components/tabs.svelte";
    
    export let problem
    
    let CodeJar
    onMount(async () => {
        ({CodeJar} = await import("@novacbn/svelte-codejar"));
    });
    export let value = ""

    let languages = ["cpp", "java", "js", "py"]
    let selected

    let username
    beforeUpdate(() => {
        if ($accessTokenStore != "") {
            const payloadB64 = $accessTokenStore.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }
        loadSubmissions()
    })

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
    }

    async function loadSubmissions() {
        const userInput = {
            username: username,
            slug: problem.slug
        }

        const options = {
            method: "POST",
            body: JSON.stringify(userInput)
        }
        const res = await fetch("http://jpoly1219devbox.xyz:8090/submissions", options)
        const data = await res.json()
        console.log(data, typeof(data))
        const submissionsData = data.map((data) => {
            return {
                username: data.username,
                slug: data.slug,
                lang: data.lang,
                code: data.code,
                result: data.result,
                output: data.output
            }
        })
        console.log("submissionsData:", submissionsData)
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
            <p class="font-bold">{problem.title}</p>
            <p class="text-sm text-green-600 font-light mt-2">{problem.difficulty}</p>
            <hr class="my-4">
            <p class="prose max-w-max">{@html snarkdown(problem.description)}</p>
            {:else if activeTab === "Solution"}
            <p class="font-bold">Solution</p>
            {:else if activeTab === "Discussion"}
            <p class="font-bold">Discussion</p>
            {:else if activeTab === "Submissions"}
            <p class="font-bold">Submissions</p>
                {#if resultData}
                <p>{resultData.result}</p>
                <p>{resultData.expected}</p>
                <p>{resultData.output}</p>
                {/if}
                {#if submissionsData}
                <table>
                    <tr>
                        <th>Result</th>
                        <th>Expected</th>
                        <th>Output</th>
                    </tr>
                    {#each submissionsData as submissionsDatum}
                    <tr>
                        <td>{submissionsDatum}</td>
                    </tr>
                    {/each}
                </table>
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
