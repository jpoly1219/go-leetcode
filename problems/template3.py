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