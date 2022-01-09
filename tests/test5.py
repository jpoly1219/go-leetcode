# template file
import json
from typing import List


# insert Solution class here
class Solution:
    def longestPalindrome(self, s):
        lenS = len(s)
        if lenS <= 1: return s
        minStart, maxLen, i = 0, 1, 0
        while i < lenS:
            if lenS - i <= maxLen / 2: break
            j, k = i, i
            while k < lenS - 1 and s[k] == s[k + 1]: k += 1
            i = k + 1
            while k < lenS - 1 and j and s[k + 1] == s[j - 1]:  k, j = k + 1, j - 1
            if k - j + 1 > maxLen: minStart, maxLen = j, k - j + 1
        return s[minStart: minStart + maxLen]


with open("tc5.json", "r") as read_file:
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