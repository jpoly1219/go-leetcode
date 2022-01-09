// template file
#include <fstream>
#include <iostream>
#include <vector>
#include <string>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

// insert Solution class here


int main() {
    Solution sol;

    ifstream i("../testcase-5-longest-palindromic-substring.json");
    json j;
    i >> j;

    vector<string> vecInput = j["input"];
    vector<vector<string>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (size_t i = 0; i < vecInput.size(); i++) {
        string solution  = sol.longestPalindrome(vecInput.at(i));
        vector<string> expected = vecExpected.at(i);
        if (find(expected.begin(), expected.end(), solution) == expected.end()) {
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