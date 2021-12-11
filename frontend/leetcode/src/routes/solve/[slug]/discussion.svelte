<script>
import Discussions from "src/components/discussions.svelte";
import { load } from "./solution.svelte";

</script>
<script context="module">
    export async function load({page}) {
        const fullPath = page.path
        const slugArray = fullPath.split("/")
        const slug = slugArray[2]

        const url = `http://jpoly1219devbox.xyz:8090/discussions/${slug}`

        const options = {
            method: "GET"
        }

        try {
            const res = await fetch(url, options)
            const discussions = await res.json()
            return {props: {slug, discussions}}
        } catch(err) {
            console.log(err)
        }
    }
</script>

<script>
    export let slug
    export let discussions
</script>

<div>
    <Discussions slug={slug} discussions={discussions}/>
</div>