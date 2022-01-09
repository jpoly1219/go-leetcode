# template file
import json
from typing import List


# insert Solution class here



with open("../testcase-5-longest-palindromic-substring.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

listInput = data["input"]
listExpected = data["expected"]

for i, element in enumerate(listInput):
    expected = listExpected[i]
    solution = sol.longestPalindrome(element)
    if solution not in expected:
        isOk = False
        break

if isOk:
    output = {
        "result": "OK"
    }
    json_object = json.dumps(output, indent=4)
    print(json_object)

print("test completed")