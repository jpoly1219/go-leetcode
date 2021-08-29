#include <iostream>
#include <iterator>
#include <sstream>
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
    vector<vector<int>> vecExpected;
};

int main() {
    struct Testcases tc;
    Solution sol;

    // add testcases from the backend using Go's file write functions
    // tc.vecNums = ...
    // tc.vecTarget = ...
    // tc.vecExpected = ...
    tc.vecNums = {{1, 2, 3, 4, 5}, {1, 2, 3}};
    tc.vecTargets = {1, 2};
    tc.vecExpected = {{2, 3, 4, 5, 6}, {3, 4, 5}};

    ostringstream oss;

    // test
    for (int i = 0; i < tc.vecNums.size(); i++) {
        vector<int> vecSolution = sol.twoSum(tc.vecNums.at(i), tc.vecTargets.at(i));
        if (vecSolution != tc.vecExpected.at(i)) {
            cout << "Wrong Answer" << endl << "input: nums=";
            copy(tc.vecNums.at(i).begin(), tc.vecNums.at(i).end()-1, ostream_iterator<int>(oss, ","));
            oss << tc.vecNums.at(i).back() << ", target=" << tc.vecTargets.at(i);
            cout << oss.str() << endl << "expected: ";
            oss.str("");
            copy(tc.vecExpected.at(i).begin(), tc.vecExpected.at(i).end()-1, ostream_iterator<int>(oss, ","));
            oss << tc.vecExpected.at(i).back();
            cout << oss.str() << endl << "output: ";
            oss.str("");
            copy(vecSolution.begin(), vecSolution.end()-1, ostream_iterator<int>(oss, ","));
            oss << vecSolution.back();
            cout << oss.str() << endl;
            break;
        }
        else {
            cout << "OK" << endl;
        }
    }
}