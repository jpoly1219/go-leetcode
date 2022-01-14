# go-leetcode

A clone of Leetcode written in Go and SvelteKit. I wanted to see if I could reverse-engineer a popular web app with complicated internals.
The goal of this side project was to learn how to use Go for backend development, SvelteKit and TailwindCSS for scaffolding the frontend, learn a new database, and understand how Docker and Docker Compose works.

## Languages and Tools Used:
- Go
- SvelteKit
- TailwindCSS
- Postgresql
- Docker, Docker Compose

## Authentication Workflow

## App Workflow
- The frontend service renders these pages:
  - Index page to welcome users.
  - Sign Up / Login page for user authentication.
  - Problems page to list and filter available problems.
  - Problem page that shows the problem description, solution, discussion and submissions, along with an online editor where users can write their code.

- The backend service acts as an API gateway of sort.
  - `main.go` handles incoming requests and relays them to the corresponding controller function in `controllers.go`. It is also responsible for loading the `.env` file and connecting to the database.
  - `auth.go` holds the functions necessary for user authentication, which includes handling CORS requests, generating JWTs, handling user sign ins and logins, and handling silent refreshes.
  - `controllers.go` holds all the controller functions.
  - `middlewares.go` holds the middlewares that wraps certain controller functions.
  - `models.go` holds structs that mirror the database schema.


- User writes a code within the online editor provided by the frontend.
- The code is then sent to the backend once the user presses the submit button. The data is transferred via REST API.
- The API gateway receives the code and checks what problem the user is solving. Then it queries the database for the appropriate template code and testcases.
- The backend then hands it over to a separate "code-run" container which provides an isolated environment for code to be run in.
- The code-run container generates files with the user code and template code inside it, then compiles/runs it to compare outputs with the testcase.
  - It handles concurrent runs by generating unique file names using Google's `uuid` package.
- Once the run completes, the output will either be a compile error, runtime error, or a successful run. The output is sent back to the API gateway.
- The result is marshalled into JSON and is sent back to the frontend, which renders it as a Javascript alert.
- User authentication is a token-based auth system using JWT. 