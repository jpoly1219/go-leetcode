# Structure

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
  - So... create a file named `uuid.filetype`, create a `Dockerfile` with a dynamically generated `RUN` statement, get output and errors
  - How to communicate between container and host?
- Learn a different database such as PostgreSQL:
  - `users` database for storing user info
    - ID, username, password, etc.
  - `submissions` database for storing user submissions
    - Submission number, output, run time, success/fail, user (foreign key)
- Learn Test Driven Development:

- Frontend:
  - Use CodeMirror for implementing a code editor inside our frontend
  - https://codemirror.net/

### Work Log:
- Created frontend using SvelteKit.
- Created `index` and `problemset` pages.
- Created handler for `/problemset` endpoint in main.go.
- Created an empty controller for `/problemset` endpoint in main.go.
- Created an empty database `problems` using PostgreSQL.
### TODO:
- Load `problems` database with problem sets.
- Create a query in `controllers.go` to select all problems. More detailed queries will be implemented later.
- Call the Go API with `fetch` inside `problemstore.js`.