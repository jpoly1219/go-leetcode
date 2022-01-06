# C++
```
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
```

# Java
```
```

# Javascript
```
// template file
const tc = require("../testcase-3-longest-substring-without-repeating-characters.json")

// insert Solution class here


let isOk = true
const arrInput = tc.input
const arrExpected = tc.expected
for (let i = 0; i < arrInput.length; i++) {
    const solution = lengthOfLongestSubstring(arrInput[i])
    if (solution != arrExpected[i]) {
        isOk = false
        break
    }

}

if (isOk) {
    const output = {
        "result": "OK"
    }
    const data = JSON.stringify(output)
    console.log(data)
}

console.log("test completed")
```

# Python
```
# template file
import json
from typing import List


# insert Solution class here



with open("../testcase-3-longest-substring-without-repeating-characters.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

listInput = data["input"]
listExpected = data["expected"]

for i, element in enumerate(listInput):
    solution = sol.lengthOfLongestSubstring(element)
    if solution != listExpected[i]:
        isOk = False
        break

if isOk:
    output = {
        "result": "OK"
    }
    json_object = json.dumps(output, indent=4)
    print(json_object)

print("test completed")
```