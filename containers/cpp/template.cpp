#include <fstream>
#include <iostream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

// insert Solution class here

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
    }
    if (isOk) {
        json output = {
            {"result", "OK"}
        };
        o << output << endl;
    }
    cout << "done" << endl;
    i.close();
    o.close();
}

