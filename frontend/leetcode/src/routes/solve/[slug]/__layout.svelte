<script context="module">
    import hljs from "highlight.js/lib/core"
    import javascript from "highlight.js/lib/languages/javascript"
    import { get } from "svelte/store"
    import { accessTokenStore } from "../../../stores/stores"

    hljs.registerLanguage("js", javascript)
    const highlight = (code, syntax) =>
        hljs.highlight(code, {
            language: syntax,
        }).value
    
    export function load({page}) {
        const slug = page.params.slug

        let accessToken = get(accessTokenStore)
        let username
        if (accessToken != "") {
            const payloadB64 = accessToken.split(".")[1]
            username = JSON.parse(window.atob(payloadB64)).username
        }

        return {props: {slug, username}}
    }
</script>

<script>
    import { beforeUpdate, onMount } from "svelte"
    import { problemsListStore, submitCodeStore } from "../../../stores/stores.js"
    import { goto } from "$app/navigation"
    import Tabs from "../../../components/tabs.svelte";

    // Props from the module context script
    export let slug
    export let username

    // CodeJar editor
    let CodeJar
    export let value = "console.log('hello world')"
    let languages = ["cpp", "java", "js", "py"]
    let selected = "cpp"

    // Page UI
    onMount(async () => {
        ({CodeJar} = await import("@novacbn/svelte-codejar"));
    });
    
    $: currentIndex = $problemsListStore.indexOf(slug)
    let prevSlug
    let nextSlug
    beforeUpdate(() => {
        const prevIndex = currentIndex - 1
        const nextIndex = currentIndex + 1
    
        if (prevIndex >= 0) {
            prevSlug = $problemsListStore[prevIndex] + "/description"
        } else {
            prevSlug = slug + "/description"
        }
    
        if (nextIndex < $problemsListStore.length) {
            nextSlug = $problemsListStore[nextIndex] + "/description"
        } else {
            nextSlug = slug + "/description"
        }
        
    })

    let randSlug = "1-two-sum/description"

    function generateRandSlug() {
        const randIndex = Math.floor(Math.random() * $problemsListStore.length)
        if (randIndex === currentIndex) {
            generateRandSlug()
        } else {
            randSlug = $problemsListStore[randIndex] + "/description"
        }
    }

    // On code submit
    function submitCode() {
        alert("code submitted!")

        submitCodeStore.set({
            username: username,
            slug: slug,
            lang: selected,
            code: value
        })

        goto(`/solve/${slug}/submissions`)
    }
</script>

<div class="grid grid-rows-16 h-full">
    <div class="row-span-15 grid grid-cols-2 gap-4">
        <div class="overflow-auto border border-gray-300 p-4">
            <Tabs slug={slug}/>
            <slot></slot>
        </div>
        <div class="flex flex-col border border-gray-300 overflow-hidden">
            <div class="overflow-auto">
                {#if CodeJar}
                    <svelte:component this={CodeJar} class="hljs" addClosing={true} indentOn={/{$/} spellcheck={false} tab={"\t"} withLineNumbers={true} syntax="js" {highlight} {value}/>
                {:else}
                    <pre><code>{value}</code></pre>
                {/if}
            </div>
        </div>
    </div>
    <div class="row-span-1 grid grid-cols-2 gap-4 content-center">
        <div class="flex flex-row">
            <div class="flex-1 flex">
                <a href="/problemset">
                    <button class="border border-gray-300 rounded-lg px-3 py-2">Problems</button>
                </a>
            </div>
            <a href={`/solve/${randSlug}`}>
                <button on:click={generateRandSlug} class="border border-gray-300 rounded-lg px-3 py-2">Pick One</button>
            </a>
            <a href={`/solve/${prevSlug}`}>
                <button class="border border-gray-300 rounded-lg px-3 py-2 mx-4">Prev</button>
            </a>
            <div class="mx-2 flex">
                <p class="self-center">1/1977</p>
            </div>
            <a href={`/solve/${nextSlug}`}>
                <button class="border border-gray-300 rounded-lg px-3 py-2 ml-4">Next</button>
            </a>
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
            <button on:click={submitCode} class="border border-gray-300 rounded-lg px-3 py-2 ml-2">Submit</button>
        </div>
    </div>
</div>