#include <iostream>
#include <vector>

using namespace std;

// insert Solution class from the backend using Go's file write functions
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> solution;

        for (auto& it : nums) {
            solution.push_back(it + target);
        }
        return solution;
    }
};

struct Testcases {
    vector<vector<int>> vecNums;
    vector<int> vecTargets;
};

int main() {
    struct Testcases tc;
    Solution sol;

    // add testcases from the backend using Go's file write functions
    // tc.vecNums = ...
    // tc.vecTarget = ...
    tc.vecNums = {{1, 2, 3, 4, 5}, {1, 2, 3}};
    tc.vecTargets = {1, 2};

    for (int i = 0; i < tc.vecNums.size(); i++) {
        vector<int> vecSolution = sol.twoSum(tc.vecNums.at(i), tc.vecTargets.at(i));
        for (auto it : vecSolution) {
            cout << it << " ";
        }
        cout << endl;
    }
}