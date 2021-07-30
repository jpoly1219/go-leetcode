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
- resources:
  - https://callistaenterprise.se/blogg/teknik/2019/10/05/go-worker-cancellation/