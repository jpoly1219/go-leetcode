# template file
import json
from typing import List


# insert Solution class here
class Solution:
    # @return a float
    def findMedianSortedArrays(self, A, B):
        l=len(A)+len(B)
        return self.findKth(A,B,l//2) if l%2==1 else (self.findKth(A,B,l//2-1)+self.findKth(A,B,l//2))/2.0
            
            
    def findKth(self,A,B,k):
        if len(A)>len(B):
            A,B=B,A
        if not A:
            return B[k]
        if k==len(A)+len(B)-1:
            return max(A[-1],B[-1])
        i=len(A)//2
        j=k-i
        if A[i]>B[j]:
            #Here I assume it is O(1) to get A[:i] and B[j:]. In python, it's not but in cpp it is.
            return self.findKth(A[:i],B[j:],i)
        else:
            return self.findKth(A[i:],B[:j],j)


with open("tc4.json", "r") as read_file:
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