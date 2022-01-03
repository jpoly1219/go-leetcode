// template file
#include <fstream>
#include <iostream>
#include <vector>
#include <string>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

// insert Solution class here
class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.length();
        
        unordered_map<int, int>m;
        int c = 0, maxlen = -1;
        
        for (int i = 0; i < n; i++) {
            if (m.find(s[i]) == m.end() || (i - m[s[i]] > c))
                c++;
            else {
                maxlen = max(maxlen, c);
                c = i-m[s[i]];
            }
            m[s[i]] = i;
        }
        
        return max(maxlen, c);
    }
};

int main() {
    Solution sol;

    ifstream i("tc3.json");
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