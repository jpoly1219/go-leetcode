import json

# insert Solution class here
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        pass

with open("testcase.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True
# for each input pair, check if Solution.twoSum() returns a matching expected answer
for i in range(len(data["input"]["nums"])):
    answer = sol.twoSum(data["input"]["num"][i], data["input"]["target"][i])
    if answer != data["expected"][i]:
        # write json to file
        isOk = False
        break

if isOk:
    # write json to file
    print("ok")

print("done")