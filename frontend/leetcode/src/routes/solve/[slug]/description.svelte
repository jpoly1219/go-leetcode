<script>

</script>

<div>
    <p class="text-lg font-bold mb-3">{problem.title}</p>
    <p class="text-sm text-green-600 font-light">{problem.difficulty}</p>
    <hr class="my-4">
    <p class="prose max-w-max">{@html snarkdown(problem.description)}</p>
</div>