# List of known bugs
- RunTest() receives poorly formatted JSON from the frontend
  - Tested
  - The frontend parses languages as `C++`, `Java`, etc., while the backend expects `cpp`, `java`, etc.
  - Suggested fix:
    - Properly format the request JSON body.