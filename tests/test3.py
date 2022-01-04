# template file
import json
from typing import List


# insert Solution class here
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0
        res = 1
        slow = 0
        for i in range(1, len(s)):
            if s[i] in s[slow : i]:
                slow = s[slow:i].index(s[i]) + slow + 1
                #print(slow)
            res = max(res, i - slow + 1)
        return res

with open("tc3.json", "r") as read_file:
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