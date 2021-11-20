<script>
    import { onMount } from "svelte"

    export let comment

    let userData
    onMount(async () => {
        const url = `http://jpoly1219devbox.xyz:8090/users/${comment.author}`
        const options = {
            method: "GET"
        }
        const res = await fetch(url, options)
        const data = await res.json()
        userData = {
            username: data.username,
            fullname: data.fullname,
            email: data.email,
            profilePic: data.profilePic
        }
    })
</script>

<div class="">
    <div class="">
        <div class="flex flex-row items-center mb-2">
            <img 
                src="https://img.freepik.com/free-photo/pleasant-looking-serious-man-stands-profile-has-confident-expression-wears-casual-white-t-shirt_273609-16959.jpg?size=626&ext=jpg"
                class="rounded-full w-10 h-10 mr-2"
            >
            <p class="text-sm mr-2">{comment.author}</p>
            <p class="text-sm">{comment.created}</p>
        </div>
        <div class="flex flex-row">
            <div class="w-10 mr-2"></div>
            <p>{comment.description}</p>
        </div>
    </div>
</div>