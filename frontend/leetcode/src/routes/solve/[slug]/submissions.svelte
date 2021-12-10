<script>

</script>

<div>
    <p class="text-lg font-bold mb-3">Submissions</p>
    {#if submissionsData}
        <div class="w-full">
            <table class="table-fixed items-center w-full border-collapse">
                <thead>
                    <tr>
                        <th class="w-1/5 px-4 py-2 bg-gray-200 border border-solid border-gray-100 border-r-0 text-sm text-gray-700 text-left">Result</th>
                        <th class="w-4/5 px-4 py-2 bg-gray-200 border border-solid border-gray-100 border-l-0 text-sm text-gray-700 text-left">Output</th>
                    </tr>
                </thead>
                {#each submissionsData as submissionsDatum}
                    <tbody>
                        <tr>
                            <td class="px-4 py-2 text-sm text-left {submissionsDatum.result === 'OK' ? 'text-green-600' : 'text-red-600'}">
                                {submissionsDatum.result}
                            </td>
                            <td class="px-4 py-2 text-sm text-left break-words">{submissionsDatum.output}</td>
                        </tr>
                    </tbody>
                {/each}
            </table>
        </div>
    {/if}
</div>