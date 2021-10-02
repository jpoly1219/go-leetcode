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