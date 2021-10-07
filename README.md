# go-leetcode

A clone of Leetcode written in Go and SvelteKit.

## Languages and Tools Used:
- Go
- SvelteKit
- Tailwind CSS
- PostgreSQL
- Docker Compose

## App Workflow
- User writes a code within the online editor provided by the frontend.
- The code is then sent to the backend once the user presses the submit button. The data transfer is done via REST API.
- The API gateway receives the code and checks what problem the user is solving. Then it queries the database for the appropriate template code and testcases.
- The backend then hands it over to a separate "code-run" container which provides an isolated environment for code to be run in.
- The code-run container generates files with the user code and template code inside it, then compiles/runs it to compare outputs with the testcase.
  - It handles concurrent runs by generating unique file names using Google's `uuid` package.
- Once the run completes, the output will either be a compile error, runtime error, or a successful run. The output is sent back to the API gateway.
- The result is marshalled into JSON and is sent back to the frontend, which renders it as a Javascript alert.
- User authentication is a token-based auth system using JWT. 