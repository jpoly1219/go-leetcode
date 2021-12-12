import { writable } from "svelte/store"

export const timeToExpireStore = writable("")
export const accessTokenStore = writable("")
export const problemsListStore = writable([])
export const submitCodeStore = writable({})