# List of known bugs
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

- Panic after RunTest() is reached
  - Tested
  - Output:
    ```
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
  - Suggested fix:
    - Check the db query statement for any typos.