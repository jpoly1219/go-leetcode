# template file
import json
from typing import List


# insert Solution class here



with open("../testcase-4-median-of-two-sorted-arrays.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

listNums1 = data["input"]["nums1"]
listNums2 = data["input"]["nums2"]
listExpected = data["expected"]

for i, element in enumerate(listNums1):
    num1 = listNums1[i]
    num2 = listNums2[i]
    solution = sol.findMedianSortedArrays(num1, num2)
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