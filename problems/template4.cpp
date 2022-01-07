// template file
#include <fstream>
#include <iostream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

// insert Solution class here


int main() {
    Solution sol;

    ifstream i("../testcase-4-median-of-two-sorted-arrays.json");
    json j;
    i >> j;

    vector<vector<int>> vecNums1 = j["input"]["nums1"];
    vector<vector<int>> vecNums2 = j["input"]["nums2"];
    vector<double> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (size_t i = 0; i < vecNums1.size(); i++) {
        vector<int> num1 = vecNums1.at(i);
        vector<int> num2 = vecNums2.at(i);
        double solution  = sol.findMedianSortedArrays(num1, num2);
        if (solution != vecExpected.at(i)) {
            isOk = false;
            break;
        }
    }
    if (isOk) {
        json output = {
            {"result", "OK"}
        };
        cout << output.dump(4) << endl;
    }
    cout << "test completed" << endl;
    i.close();
}