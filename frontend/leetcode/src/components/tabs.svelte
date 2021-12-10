<script>
    export let slug
</script>

<div class="mb-3">
    <ul class="flex justify-start p-0 list-none">
        <li class="mr-4 text-sm cursor-pointer">
            <a href={`/solve/${slug}/description`}>
                <p>Description</p>
            </a>
        </li>
        <li class="mr-4 text-sm cursor-pointer">
            <a href={`/solve/${slug}/solution`}>
                <p>Solution</p>
            </a>
        </li>
        <li class="mr-4 text-sm cursor-pointer">
            <a href={`/solve/${slug}/discussion`}>
                <p>Discussion</p>
            </a>
        </li>
        <li class="mr-4 text-sm cursor-pointer">
            <a href={`/solve/${slug}/submissions`}>
                <p>Submissions</p>
            </a>
        </li>
    </ul>
</div>