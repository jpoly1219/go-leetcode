# List of feature and refactoring requests

## 2023

- Make HTTPS work.
  - For this, we are going to have to do some overhaul with the routing.
  - nginx can handle locations, but these locations come after the url.
  - So location / is for site.com/, and /backend is for site.com/backend.
  - The plan is to wrap all the backend to /backend, and coderunner to /coderunner.
  - Frontend won't need this, and I don't think the database needs it?
    - Although if necessary, I will wrap database to /database.

## Past

- Refactoring frontend code (DONE)

  - If we are using SvelteKit, we should try to use its SSR capabilities. AKA, use it how it's meant to be.
  - Right now there is a /solve page that basically loads all the subpages via Svelte components.
  - This makes each subpage components such as `discussions.svelte` not reusable, which goes against the purpose of reusable components.
  - There are ways to separate frontend code without relying on dynamically loading components in SvelteKit.
    - Nested layouts, JS endpoints

- Refactoring `coderunner` code

  - All the code is inside `main.go`. This makes the file unnecessarily long and difficult to maintain.
  - There should be a structure where the utils and the controllers are separated into their own files.
  - Keep the `cpp`, `java`, `js`, `py` directories.
  - `utils.go` should hold functions such as `FileToLines`, `LinesFromFile`, `WriteCodeToFile`, etc.
  - `main.go` should hold the db connection and mux.
  - `interfaces.go` sounds like a good idea, but I need to research this a bit more.
  - `models.go` should hold the struct definitions.
  - `controllers.go` should hold the controller.

- Separate `auth` into its own service?
  - Many microservices have an auth server running separately from the rest of the program.
  - Separation of concerns.
  - The auth server will have a Redis cache to store session id's.
  - Or just use Auth0.
