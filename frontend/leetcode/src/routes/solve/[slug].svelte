<script>

import { problems } from "src/stores/problemstore";
import { loop_guard } from "svelte/internal";

</script>
<script context="module">
    export async function load({page}) {
        const slug = page.params.slug
        const url = `http://jpoly1219devbox.xyz:8090/${slug}`
        const res = await fetch(url)
        const problem = await res.json()
        return {props: {problem}}
    }
</script>

<script>
    export let problem
</script>

<h1>{problems.title}</h1>