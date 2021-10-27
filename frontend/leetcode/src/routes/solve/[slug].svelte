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
        // console.log(`now fetching:\n ${options.headers.Authorization}`)
        try {
            const res1 = await fetch(url1, options1)
            const res2 = await fetch(url2, options2)
            const problem = await res1.json()
            const submissions = await res2.json()
            console.log(problem, submissions)
            return {props: {problem, submissions, username}}
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

    /*
    let username
    beforeUpdate(() => {
        if ($accessTokenStore != "") {
            const payloadB64 = $accessTokenStore.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }
    })
    */

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

    /*
    let submissionsData
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
        submissionsData = data.map((data) => {
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
    */

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
                {#if submissionsData}
                <table class="table-auto">
                    <tr>
                        <th>Result</th>
                        <th>Output</th>
                    </tr>
                    {#if resultData}
                    <tr>
                        <td>{resultData.result}</td>
                        <td>{resultData.output}</td>
                    </tr>
                    {/if}
                    {#each submissionsData as submissionsDatum}
                    <tr>
                        <td>{submissionsDatum.result}</td>
                        <td>{submissionsDatum.output}</td>
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
