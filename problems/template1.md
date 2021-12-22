# C++
```
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

    ifstream i("../testcase-1-two-sum.json");
    json j;
    i >> j;

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
            cout << output.dump(4) << endl;
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
// template file
// insert Solution class here
class Solution {
    public int[] twoSum(int[] nums, int target) {
        int[nums.length] output;
        for (int i = 0; i < nums.length; i++) {
            output[i] = nums[i] + target;
        }
        return output;
    }
}

public class Template {
    public static void main(String[] args) {
        Solution sol = new Solution();
        boolean isOk = true;
        for (int i = 0; i < )
    }
}
```

# Javascript
```
// template file
const tc = require("../testcase-1-two-sum.json")

// insert Solution class here

let isOk = true
for (let i = 0; i < tc.input.nums.length; i++) {
    const answer = twoSum(tc.input.nums[i], tc.input.target[i])
    if (answer != tc.expected[i]) {
        const output = {
            "result": "wrong",
            "input": tc.input.nums[i],
            "expected": tc.expected[i],
            "output": answer
        }
        const data = JSON.stringify(output)
        console.log(data)
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

# Python3
```
# template file
import json
from typing import List


# insert Solution class here


with open("../testcase-1-two-sum.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True
# for each input pair, check if Solution.twoSum() returns a matching expected answer
for i in range(len(data["input"]["nums"])):
    answer = sol.twoSum(data["input"]["nums"][i], data["input"]["target"][i])
    if answer != data["expected"][i]:
        # write json to file
        output = {
            "result": "wrong",
            "input": data["input"]["nums"][i],
            "expected": data["expected"][i],
            "output": answer
        }
        json_object = json.dumps(output, indent=4)
        print(json_object)
        
        isOk = False
        break

if isOk:
    # write json to file
    output = {
        "result": "OK"
    }
    json_object = json.dumps(output, indent=4)
    print(json_object)

print("test completed")

```