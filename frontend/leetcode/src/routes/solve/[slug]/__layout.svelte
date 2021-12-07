<script>
    
</script>

<div class="grid grid-rows-16 h-full">
    <div class="row-span-15 grid grid-cols-2 gap-4">
        <div class="overflow-auto border border-gray-300 p-4">
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
            <button on:click={submit} class="border border-gray-300 rounded-lg px-3 py-2 ml-2">Submit</button>
        </div>
    </div>
</div>