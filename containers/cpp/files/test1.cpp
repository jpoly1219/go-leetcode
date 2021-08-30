#include <fstream>
#include <iostream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

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

int main() {
    Solution sol;

    // add testcases from the backend using Go's file write functions
    // tc.vecNums = ...
    // tc.vecTarget = ...
    // tc.vecExpected = ...
    ifstream i("testcase.json");
    json j;
    i >> j;

    ofstream o("result.json", ios::trunc);

    vector<vector<int>> vecNums = j["input"]["nums"];
    vector<int> vecTargets = j["input"]["target"];
    vector<vector<int>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (int i = 0; i < vecNums.size(); i++) {
        vector<int> vecSolution = sol.twoSum(vecNums.at(i), vecTargets.at(i));
        if (vecSolution != vecExpected.at(i)) {
            json output = {
                {"result", "wrong"},
                {"input", vecNums.at(i)},
                {"expected", vecExpected.at(i)},
                {"output", vecSolution}
            };
            o << output << endl;
            isOk = false;
            break;
        }
        else {
            cout << "OK" << endl;
        }
    }
    if (isOk) {
        json output = {
            {"result", "OK"}
        };
        o << output << endl;
    }
    i.close();
    o.close();
}