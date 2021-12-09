# List of feature and refactoring requests

- Refactoring frontend code
  - If we are using SvelteKit, we should try to use its SSR capabilities. AKA, use it how it's meant to be.
  - Right now there is a /solve page that basically loads all the subpages via Svelte components.
  - This makes each subpage components such as `discussions.svelte` not reusable, which goes against the purpose of reusable components.
  - There are ways to separate frontend code without relying on dynamically loading components in SvelteKit.
    - Nested layouts, JS endpoints