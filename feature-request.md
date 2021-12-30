# List of feature and refactoring requests

- Refactoring frontend code
  - If we are using SvelteKit, we should try to use its SSR capabilities. AKA, use it how it's meant to be.
  - Right now there is a /solve page that basically loads all the subpages via Svelte components.
  - This makes each subpage components such as `discussions.svelte` not reusable, which goes against the purpose of reusable components.
  - There are ways to separate frontend code without relying on dynamically loading components in SvelteKit.
    - Nested layouts, JS endpoints

- Refactoring `coderunner` code
  - All the code is inside `main.go`. This makes the file unnecessarily long and difficult to maintain.
  - There should be a structure where the utils and the controllers are separated into their own files.