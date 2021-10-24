# List of known bugs
- Docker Compose not reflecting code updates (**FIXED 21.10.24 23:20**)
  - Tested
    - Filesystem inside the container reflects the code updates, but running `docker logs` show no difference between before and after the update.
    - Changes inside `main.go` are applied, but the ones inside `/pkg` are not.
  - Possible reasons:
    - `docker-compose` is using the cached layers instead of rebuilding.
    - *Go isn't building binaries correctly.*
    - *Go isn't importing `pkg` properly.*
  - Suggested fix:
    - `docker-compose up --build`
    - `docker-compose build --no-cache` && `docker-compose up`
    - *Check for any typos in Go files, especially the ones inside `/backend` directory.*
    - *Review how Go modules and packages work.*

- Problemsets() returns an array data type, but Submissions() returns an object data type (**FIXED 21.10.24 23:20**)
  - Tested
  - Possible reasons:
    - *API is sending over different data types.*
    - Frontend is parsing the data weirdly.
  - Suggested fix:
    - *Check the backend API to see if it is sending a JSON or an array of JSONs.*

- Directory structure changes inside Docker container

- RunTest() receives poorly formatted JSON from the frontend
  - Tested
  - Output:
    ```
    RunTest() reached:    C++ class Solution { 
        public: vector<int> twoSum(vector<int>& nums, int target) {
            vector<int> solution; 
            for (auto& it : nums) { 
                solution.push_back(it + target); 
            } 
            
            return solution; 
        }
    };
    ```
  - Possible reasons:
    - The frontend parses languages as `C++`, `Java`, etc., while the backend expects `cpp`, `java`, etc.
  - Suggested fix:
    - Properly format the request JSON body.

- Unable to parse JSON response from backend
  - Tested
  - Output:
    ```
    SyntaxError: Unexpected token M in JSON at position 0
        at JSON.parse (<anonymous>)
        at Response.json (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/install-fetch.js:546:15)
        at processTicksAndRejections (node:internal/process/task_queues:96:5)
        at async load (/root/development/go-leetcode/frontend/leetcode/src/routes/solve/[slug].svelte:30:19)
        at async load_node (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/ssr.js:937:12)
        at async respond$1 (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/ssr.js:1221:15)
        at async render_page (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/ssr.js:1386:19)
        at async resolve (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/ssr.js:1648:10)
        at async respond (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/ssr.js:1629:10)
        at async Immediate.<anonymous> (file:///root/development/go-leetcode/frontend/leetcode/node_modules/@sveltejs/kit/dist/chunks/index.js:3499:22)
    ```
  - Possible reasons:
    - Bad response from backend
  - Suggested fix:
    - Check backend code to see if JSON is properly parsed, or if the frontend is feeding the backend with enough information.

- Panic after RunTest() is reached (**FIXED 21.10.03 16:33**)
  - Tested
  - Output:
    ```
    ./main.go output:
    Error making POST request:  Post "http://jpoly1219devbox.xyz:8091/run": EOF

    ./containers/main.go output:
    2021/10/02 06:32:20 http: panic serving 141.164.41.239:55438: runtime error: invalid memory address or nil pointer dereference
    goroutine 20 [running]:
    net/http.(*conn).serve.func1(0xc0000b2c80)
            /usr/local/go/src/net/http/server.go:1824 +0x153
    panic(0x6f6e00, 0x946920)
            /usr/local/go/src/runtime/panic.go:971 +0x499
    database/sql.(*DB).conn(0x0, 0x7b2290, 0xc0000a8000, 0xc0000ae001, 0x764220, 0xc0000bf710, 0x4dab4e)
            /usr/local/go/src/database/sql/sql.go:1197 +0x41
    database/sql.(*DB).query(0x0, 0x7b2290, 0xc0000a8000, 0x75ca32, 0x42, 0xc0000bf8e0, 0x2, 0x2, 0xc000093001, 0x120, ...)
            /usr/local/go/src/database/sql/sql.go:1623 +0x66
    database/sql.(*DB).QueryContext(0x0, 0x7b2290, 0xc0000a8000, 0x75ca32, 0x42, 0xc0000bf8e0, 0x2, 0x2, 0x693054, 0xc0000b2be0, ...)
            /usr/local/go/src/database/sql/sql.go:1605 +0xd4
    database/sql.(*DB).Query(...)
            /usr/local/go/src/database/sql/sql.go:1619
    main.RunTest(0x7b1ab0, 0xc00010e0e0, 0xc000122700)
            /root/development/go-leetcode/containers/main.go:372 +0x354
    net/http.HandlerFunc.ServeHTTP(0x764278, 0x7b1ab0, 0xc00010e0e0, 0xc000122700)
            /usr/local/go/src/net/http/server.go:2069 +0x44
    github.com/gorilla/mux.(*Router).ServeHTTP(0xc0000fc000, 0x7b1ab0, 0xc00010e0e0, 0xc000122300)
            /root/go/pkg/mod/github.com/gorilla/mux@v1.8.0/mux.go:210 +0xd3
    net/http.serverHandler.ServeHTTP(0xc00010e000, 0x7b1ab0, 0xc00010e0e0, 0xc000122300)
            /usr/local/go/src/net/http/server.go:2887 +0xa3
    net/http.(*conn).serve(0xc0000b2c80, 0x7b2300, 0xc0000d2300)
            /usr/local/go/src/net/http/server.go:1952 +0x8cd
    created by net/http.(*Server).Serve
            /usr/local/go/src/net/http/server.go:3013 +0x39b
    ```
  - Possible reasons:
    - Query statement may have a typo.
    - Scan() parameter may be dereferencing a wrong variable.
    - **main.go has a global declaration `var db *sql.Db`, but `main()` re-declares `db`, `err` using `:=`**
      - Using `:=` inside `main()` effectively shadows the global `db` variable, which makes the value of `db` nil since this makes a new `db` local to `main()`.

  - Suggested fix:
    - Check the db query statement for any typos.
    - **Fix `db` declaration by declaring `err` prior to using `db, err = sql.Open()`**

- Cannot insert into `attempts` table
  - Tested
  - Output:
  ```
  failed to insert attempt:  pq: insert or update on table "attempts" violates foreign key constraint "attempts_username_fkey"
  ```
  - Possible reasons
    - Bad query statement

  - Suggested fix:
    - Check the db query fields

- loadSubmissions() loads infinitely
  - Tested
  - Possible reasons:
   - beforeUpdate vs onMount
   - Lifecycle
   - nested async?

Uncaught (in promise) Error: {#each} only iterates over array-like objects.
    validate_each_argument index.mjs:1970
    create_if_block_5 [slug].svelte:433
    create_if_block_4 [slug].svelte:119
    update [slug].svelte:1119
    update index.mjs:1050
    flush index.mjs:1018
    promise callback*schedule_update index.mjs:993
    make_dirty index.mjs:1752
    ctx index.mjs:1790
    tabChange [slug].svelte:1283
    createEventDispatcher index.mjs:955
    createEventDispatcher index.mjs:954
    click_handler tabs.svelte:13
    click_handler tabs.svelte:47
    listen index.mjs:412
    listen_dev index.mjs:1936
    mount tabs.svelte:81
    mount tabs.svelte:154
    m svelte-hooks.js:197
    mount_component index.mjs:1720
    mount [slug].svelte:1047
    m svelte-hooks.js:197
    mount_component index.mjs:1720
    update root.svelte:107
    update root.svelte:331
    update_slot_base index.mjs:98
    update __layout.svelte:82
    update index.mjs:1050
    flush index.mjs:1018
    promise callback*schedule_update index.mjs:993
    make_dirty index.mjs:1752
    ctx index.mjs:1790
    $$set root.svelte:639
    get proxy.js:83
    $set index.mjs:1885
    key proxy.js:46
    update start.js:627
    _navigate start.js:270
    init_listeners start.js:153
    init_listeners start.js:110
    start start.js:1145
    async* login:1016

Uncaught (in promise) ReferenceError: assignment to undeclared variable data
    loadSubmissions [slug].svelte:83
    instance [slug].svelte:36
    run index.mjs:18
    mount_component index.mjs:1724
    flush index.mjs:1032
    promise callback*schedule_update index.mjs:993
    make_dirty index.mjs:1752
    ctx index.mjs:1790
    $$set root.svelte:639
    get proxy.js:83
    $set index.mjs:1885
    key proxy.js:46
    update start.js:627
    _navigate start.js:270
    init_listeners start.js:153
    init_listeners start.js:110
    start start.js:1145
    async* login:1016
