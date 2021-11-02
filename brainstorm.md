# Structure

## Workflow
- User types up code in the frontend.
  - The code needs to accept CLI input (something like int argc, char argv) to accept test case inputs and return outputs.
  - The code needs to have a class Solution (for languages that support OOP), then have it run Solution.method().
- Code is sent to the backend.
- The backend saves the code to the submissions database. (columns = username, question number, language, code, runtime, result, output)
  - Database is run inside a container that is attached to a volume.
- The backend checks what problem it is, then creates a container from the image of that problem.
  - The backend sends POST requests to the container to give the username.
  - Container is created using the Docker API
    - https://pkg.go.dev/github.com/docker/docker/client
    - https://docs.docker.com/engine/api/sdk/examples/
- The container queries the submissions databases.
  - Test cases are stored locally, for each language. This makes it easier for data to be read.
  - Making each test separate from each other, instead of running a for loop inside one file to go over all the cases, might allow me to leverage the power of goroutines to basically... run all tests concurrently?
- The container runs the code.
- Once the run is complete, the output or error is sent to the frontend.
  - Results and outputs are also saved to the submissions database.

## Frontend
- home page (before login)
- personalized dashboard page (after login)
- problems list page
  - rank by lists, difficulty, status (todo, solved, attempted), tags
  - list columns: status, title, solution, acceptance, difficulty, frequency
  - search bar for looking up questions
  - button to randomly select a problem
  - pagination (group 50 problems per page)
  - your progress on the right
- problem solving page
  - problem.md on the left, editor on the right
  - left side tabs: description, solution, discuss, submissions
  - right side: editor, drop down menu for languages, editor settings (font size, theme, tab size)
  - bottom bar: problems, randomly pick a problem, prev/next, run code (single test case, multiple test cases), submit
  - console tabs: test case, run code result, debugger
- sign up, log in, user settings page
- learn topics page
- online interview page
- resources:
  - https://kit.svelte.dev/docs
  - https://www.sitepoint.com/a-beginners-guide-to-sveltekit/
  - https://prismic.io/blog/svelte-sveltekit-tutorial
  - https://dev.to/sharu725/sveltekit-tutorial-2e5d

## Backend
- router to the pages listed above and below
  - home, dashboard, problems list, problem solving, sign up/login/user settings, learn topics, online interview
- controller functions for each route
- REST API to communicate between frontend and backend
- on code submit:
  - the code is sent over to the backend
  - the backend then generates a goroutine that creates a unique file with an extension that matches the user's choice of language
  - user's code is then inserted into the unique file and is saved
  - the goroutine will run a pre-loaded script that is used for each compile/run of the program
    - the script will hold CLI commands to build and run the code
    - the script should also have a slot where the goroutine can "plug in" the file name
  - once the script is done running, the goroutine will check run time, returned output, stdout, and error messages
  - the info mentioned above will be parsed into a JSON, then be sent back to the user
  - after sending the JSON back to the user, the unique file will be deleted and the goroutine will gracefully shut down
- user authentication with a third party auth service (it's not really worth it to implement auth myself by scratch IMO...)
  - probably Auth0? if I can wrap my head around how third party auth services work, it would be great.
  - this might be a good way to learn Firebase or Supabase, but they seem like more of a shortcut for frontend developers to scaffold together a backend...
- just realized that letting code compile and run on the backend could be a huge backdoor and a security threat
  - instead of just goroutines, we can have goroutines spawn a docker container?
  - could be resource intensive this way, but a docker container with alpine should be small enough...
  - I can just throw away any processes that take too long. timing happens outside the container, and the running happens inside the container.
  - this is a great opportunity to learn docker... I've been putting it off to the side for too long at this point
- resources:
  - https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/

## Goals
- Learn more about Go
- Try using Go's cool features: goroutines and interfaces
  - Interfaces:
    - Pretty neat if you have a lot of similar parts doing similar things but have to be implemented in a different way
    - Struct for each language might make sense. Each language has a compile/run phase that are handles differently
  - Goroutines:
    - Each goroutine should run the code and return its output.
      - The output also has to be compared to the answer and return false if it doesn't match with the expected answer.
    - Because letting people run code natively on the server is a huge security risk, the goroutines should spawn a docker container to run the code.
  - Testing:
    - `go test`
    - So `*_test.go` files need to be in the same directory as files that they are testing, but I want a better way to organize the test files.
    - https://medium.com/@matryer/5-simple-tips-and-tricks-for-writing-unit-tests-in-golang-619653f90742
    - https://tutorialedge.net/golang/intro-testing-in-go/
    - https://golang.org/doc/tutorial/add-a-test
- Learn Docker:
  - Containerized environment for each users will be good for security reasons.
  - https://dev.to/narasimha1997/building-a-secure-sandboxed-environment-for-executing-untrusted-code-7e8
  - So... create a file named `uuid.filetype`, create a `Dockerfile` with a dynamically generated `RUN` statement, get output and errors
  - How to communicate between container and host?
  - Good idea:
    - Create a docker image for each problem and spin up a container for each attempt.
    - The container will contain a Go program that basically tests the user's code.
    - The user's code should be passed into the container for each attempt.
    - This will be done using a volume or a Dockerfile? I still need to figure this out.
- Learn a different database such as PostgreSQL:
  - `users` table for storing user info
    - ID, username, password, etc.
  - `submissions` table for storing user submissions
    - Submission number, output, run time, success/fail, user (foreign key)
  - `problems` table for storing problems
    - CREATE TABLE problems (
      id serial PRIMARY KEY,
      title VARCHAR(50) UNIQUE NOT NULL,
      difficulty VARCHAR(10) NOT NULL,
      description TEXT NOT NULL,
    );
  - How should I handle test cases? Should I store them in the DB or just do it inside the backend code?
- Learn Test Driven Development:

- Frontend:
  - Use `svelte-codejar` for embedding a code editor to the frontend.
    - `codejar` itself is the only editor that I found that is installed via npm and is maintained properly.
    - Other editors such as `ace` or `codemirror` is good but are difficult to implement with Svelte.
  - https://novacbn.github.io/svelte-codejar/
  - I will look into the code later to see how it is implemented...

### TODO:
- Load `problems` database with problem sets.
  - This is done in memory at the moment, will load database later.
- Create a query in `controllers.go` to select all problems. More detailed queries will be implemented later.
- Call the Go API with `fetch` inside `problemstore.js`.
- Load problems as a markdown.
  - Save each problem as markdown, then parse this using `snarkdown`.

- 21.08.24
  - Backend needs to process the code sent from the frontend.
    - Experimentation with Docker is required, but it is not a priority.
  - Frontend needs syntax highlighting via PrismJS.
  - Frontend needs a dropdown menu with langauge selections.
    - Syntax highlighting should switch according to the language chosen.

- Inserting code to line number x
  - https://siongui.github.io/2017/01/30/go-insert-line-or-string-to-file/
  - need A Better Way To Handle Different Language Because It Is Way Too Repetitive At The Moment.
  - For a successful insert, the code must be formatted into a single line. For some reason, multi-line inserts don't seem to work...
  - Or make it so that the insert function loops over each line of user code...

- User authentication
  - Token based authentication.
  - https://hasura.io/blog/best-practices-of-using-jwt-with-graphql/
  - https://hackernoon.com/creating-a-middleware-in-golang-for-jwt-based-authentication-cx3f32z8
  - https://medium.com/swlh/building-a-user-auth-system-with-jwt-using-golang-30892659cc0
  - User sends requests to a protected route, with a bearer token.
  - Middleware extracts the token from the HTTP header, then uses `jwt-go` library to check if the token is valid.
  - If token is valid, the middleware creates a ctx variable with claims stored inside it, then hands over the request to the protected endpoint.
    - `ctx := context.WithValue(r.Context(), <key>, claims)`, `next.ServeHTTP(w, r.WithContext(ctx))` from the middleware
    - `props, _ := r.Context().Value(<key>).(jwt.MapClaims)` from the protected endpoint
  
  - Token has the following claims: `userid`, `username`, `expire`
  - Tokens won't need uuids for now.

  - Token generation:
    - User submits credentials to the backend. The backend then generates tokens.
    - Tokens will be signed by a secret key saved inside the server's environment variable.
    - Access and refresh tokens will have different secret keys.
    - Access token expires every 15min and refresh token expires every 24hrs.
    - Access token and access token expiry time will be sent as JSON payload, and refresh token will be saved inside an HttpOnly cookie.
  
  - Frontend behavior:
    - User submits credentials to the backend and receives access token and access token expiry time as JSON, and refresh token as an HttpOnly cookie.
    - After the access token expiry time passes, the frontend will send a request for silent refresh, with the refresh token inside the HttpOnly cookie.

  - Database
    - https://www.calhoun.io/inserting-records-into-a-postgresql-database-with-gos-database-sql-package/
    - `users` table with five columns: `userid`, `username`, `fullname`, `email`, and `password`
    - `password` will hold hashes instead of string
    - Table relationships:
      - `users`, `problems`, `templates`, `testcases`, `attempts`
      - `templates` has a foreign key that references the primary key of `problems`
      - `testcases` has a foreign key that references the primary key of `problems`
      - `attempts` has two foreign keys that references the primary key of `users` and `problems`
      - `CREATE TABLE users (id SERIAL PRIMARY KEY, username VARCHAR (50) UNIQUE NOT NULL, fullname VARCHAR (100) NOT NULL, email VARCHAR (255) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL)`
      - `CREATE TABLE problems (id SERIAL PRIMARY KEY, title VARCHAR (100) UNIQUE NOT NULL, slug VARCHAR (100) UNIQUE NOT NULL, difficulty VARCHAR (10) NOT NULL, description TEXT NOT NULL, created TIMESTAMP NOT NULL DEFAULT NOW())`
      - `CREATE TABLE templates (id SERIAL PRIMARY KEY, slug VARCHAR (100) NOT NULL, lang VARCHAR (10) NOT NULL, template TEXT NOT NULL, FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE)`
      - `CREATE TABLE testcases (id SERIAL PRIMARY KEY, slug VARCHAR (100) NOT NULL, testcase TEXT NOT NULL UNIQUE, FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE)`
      - `CREATE TABLE attempts (id SERIAL PRIMARY KEY, username VARCHAR (50) NOT NULL, slug VARCHAR (100) NOT NULL, lang VARCHAR (10) NOT NULL, code TEXT NOT NULL, result VARCHAR (50) NOT NULL, output TEXT NOT NULL, created TIMESTAMP NOT NULL DEFAULT NOW(), FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE, FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE)`

      - `INSERT INTO templates (slug, lang, template) VALUES ('1-two-sum', 'cpp', *paste template.cpp*)`
      - `INSERT INTO testcases (slug, testcase) VALUES ('1-two-sum', *paste template.json*)`
  
  - CORS
    - https://flaviocopes.com/golang-enable-cors/

  - SvelteKit and Reactivity, Stores
    - https://github.com/sveltejs/svelte/issues/4306
  
  - Current Issues:
    - Error loading 1-two-sum when switching tabs.
    - For some reason, SvelteKit is reloading itself every 5 seconds without me asking.
    - This is throwing off a lot of functionality.
    - https://github.com/sveltejs/kit/issues/1134
  
  - Thoughts:
    - Per-user containers might be a better idea than per-problem containers.
    - Popular problems may have issues with being overly used and thus hogging server resources.
    - Also, having a per user container would be better for security as users cannot access other users' data.
    - Even if we need to manage a lot of containers as our userbase grows, they won't be running calculations 24/7.
    - Templates for each problems should be stored elsewhere, such as inside a database.
    - Testcases should also be stored inside a database.
    - `main.go` inside the container will receive a struct of username, problem number, usercode, language.
      - Then `main.go` will pull the required data from the database by executing a query such as `SELECT template, testcase FROM problems WHERE problemNumber = 1 AND langauge = cpp"`
    - NEW TOUGHT: just have a single container for running programs and spawn goroutines within the container to handle concurrent runs of programs.
      - Load balancing will be handled via goroutines and container orchestration tools.
      - Create a docker compose with four containers: frontend, backend (API gateway), backend (running code), database
    - Because there is only going to be one container at the moment, the container needs to be able to handle concurrent requests.
      - This can be an issue because all codes that the user submits will yield a `file.*` format, making it very difficult for the program to distinguish between the two. The program may even overwrite the file with new code input.
      - One way to solve this is to generate UUIDs for each file.
      - The resulting JSON file should also follow this scheme... how?
        - Create `uuid-result.json` first. Then the template code should just print out resulting JSON as a string into stdout. `main.go` inside the container will catch that output and just insert it into `uuid-result.json`.
        - OR, just have each goroutine have its own channel that waits for its own result.
        - So because each request is handled as a goroutine, I don't need to create goroutines within a http handler, so the use of channels is not very attractive. However, what I can do is to just return the output of `exec.Command()`.

    - Does each user container have its own attempts database? Or is there going to be a separate, more central database for all user attempts?
    - I'll start off with a central database, but this won't scale too well I think.

    - Frontend is very bare-bones at the moment. It needs to render output inside console window that will toggle on and off when the user click on the console button.
    - Run code should run tests on three most lenient testcases, and not count towards your attempts.

    - Feature idea: group submissions by their steps taken
      - go-leetcode can tell you how other people have attempted the problem, and how many people have used similar steps as you.
      - Docker Compose
    
    - Dockerize the apps first before developing them!
      - Because this is a multi-container app, it is better if I can separate these into their own directory,

    - Frontend needs more polish, especially the submissions tab.
      - Right now the tab renders extremely basic table without any styling. Also, the og Leetcode has a feature where one can click on a submission entry to view detailed results.
      - This will need an update to the backend.
      - Code run time and memory usage should be measured.
      - Table style: https://tailwindcomponents.com/component/table-visits
      - Answers and discussion tabs need to be populated.
    
    - Discussions tab:
      - The concept is very similar to a blog app. It will have posts, and each post will have a title, author, time created, and comments.
      - This would be nice if it was a component. The logic is complicated, and the `[slug].svelte` file will be unreadable.