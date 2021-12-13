<script context="module">
    import { get } from "svelte/store"
    import { accessTokenStore } from "../../../stores/stores"

    export async function load({page}) {
        const fullPath = page.path
        const slugArray = fullPath.split("/")
        const slug = slugArray[2]

        const url = `http://jpoly1219devbox.xyz:8090/submissions`

        let accessToken = get(accessTokenStore)
        let username
        if (accessToken != "") {
            const payloadB64 = accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }

        const options = {
            method: "POST",
            body: JSON.stringify({username: username, slug: slug})
        }

        try {
            const res = await fetch(url, options)
            const data = await res.json()
            const loadedSubmission = data.map((data, index) => {
                return {
                    num: index + 1,
                    title: data.title,
                    slug: data.slug,
                    lang: data.lang,
                    code: data.code,
                    result: data.result,
                    input: data.input,
                    expected: data.expected,
                    output: data.output
                }
            })

            return {props: {submissions: loadedSubmission}}
        } catch(err) {
            console.log(err)
        }
    }
</script>

<script>
    import { beforeUpdate } from "svelte"
    import { submitCodeStore } from "../../../stores/stores.js"

    export let submissions

    beforeUpdate(() => {
        if (Object.keys($submitCodeStore).length !== 0) {
            alert($submitCodeStore.username)
            alert($submitCodeStore.slug)
            alert($submitCodeStore.lang)
            alert($submitCodeStore.code)
            submitCodeStore.set({})
        } else {
            console.log("else")
        }
    })
</script>

<div>
    <p class="text-lg font-bold mb-3">Submissions</p>
    {#if submissions}
        <div class="w-full">
            <table class="table-fixed items-center w-full border-collapse">
                <thead>
                    <tr>
                        <th class="w-1/5 px-4 py-2 bg-gray-200 border border-solid border-gray-100 border-r-0 text-sm text-gray-700 text-left">Result</th>
                        <th class="w-4/5 px-4 py-2 bg-gray-200 border border-solid border-gray-100 border-l-0 text-sm text-gray-700 text-left">Output</th>
                    </tr>
                </thead>
                {#each submissions as submission}
                    <tbody>
                        <tr>
                            <td class="px-4 py-2 text-sm text-left {submission.result === 'OK' ? 'text-green-600' : 'text-red-600'}">
                                {submission.result}
                            </td>
                            <td class="px-4 py-2 text-sm text-left break-words">{submission.output}</td>
                        </tr>
                    </tbody>
                {/each}
            </table>
        </div>
    {/if}
</div>