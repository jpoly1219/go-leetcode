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
