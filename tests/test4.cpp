// template file
#include <fstream>
#include <iostream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann:json;

// insert Solution class here


int main() {
    Solution sol;

    ifstream i("tc4.json")
    json j;
    i >> j;

    vector<vector<int>> vecNums1 = j["input"]["nums1"];
    vector<vector<int>> vecNums2 = j["input"]["nums2"];
    vector<double> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (size_t i = 0; i < vecNums1.size(); i++) {
        double solution  = sol.findMedianSortedArrays(*vecNums1.at(i), *vecNums2.at(i));
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