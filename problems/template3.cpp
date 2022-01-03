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

    ifstream i("../testcase-3-longest-substring-without-repeating-characters.json");
    json j;
    i >> j;

    vector<string> vecInput = j["input"];
    vector<int> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (size_t i = 0; i < vecInput.size(); i++) {
        int solution = sol.lengthOfLongestSubstring(vecInput.at(i));
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