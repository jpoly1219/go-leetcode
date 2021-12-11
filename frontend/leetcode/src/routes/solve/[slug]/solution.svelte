<script context="module">
    import snarkdown from "snarkdown"

    export async function load({page}) {
        const fullPath = page.path
        const slugArray = fullPath.split("/")
        const slug = slugArray[2]

        const url = `http://jpoly1219devbox.xyz:8090/solutions/${slug}`

        const options1 = {
            method: "GET"
        }

        try {
            const res = await fetch(url, options1)
            const solution = await res.json()
            return {props: {solution}}
        } catch(err) {
            console.log(err)
        }
    }
</script>

<script>
    export let solution
</script>

<div>
    <p class="text-lg font-bold mb-3">Solution</p>
    <hr class="my-4">
    <p class="prose max-w-max">{@html snarkdown(solution.solution)}</p>
</div>