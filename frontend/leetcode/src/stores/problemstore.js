import { writable } from "svelte/store"

export const problems = writable([])

const fetchProblem = async () => {
    const url = "http://jpoly1219devbox.xyz:8090/problemsets"
    const res = await fetch(url)
    const data = await res.json()
    const loadedProblem = data.map((data, index) => {
        return {
            num: index + 1,
            title: data.title,
            slug: data.slug,
            difficulty: data.difficulty,
            description: data.description,
            created: data.created
        }
    })
    problems.set(loadedProblem)
}

fetchProblem()