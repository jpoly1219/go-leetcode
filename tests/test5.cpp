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
    string longestPalindrome(string s) {
        
        int n = s.length();
        
        int mxlen = 0, start = 0;
        
        for (int i = 0; i < n; i++) {
            
            int j = i;
            int k = i;
            
            while (k+1 < n && s[k] == s[k+1]) k++;
            
            while (j-1 >= 0 && k+1 < n && s[j-1] == s[k+1]) {
                j--; k++;
            }
            
            if (mxlen < k-j+1) {
                mxlen = k-j+1;
                start = j;
            }
        }
        
        return s.substr(start, mxlen);
    }
};

int main() {
    Solution sol;

    ifstream i("tc5.json");
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