# template file
import json
from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def createLinkedList(nums):
    head = None
    for num in reversed(nums):
        newNode = ListNode(num, head)
        head = newNode
    
    return head


def linkedListToVector(node):
    intList = []
    while (node != None):
        intList.append(node.val)
        node = node.next
    
    return intList


# insert Solution class here
class Solution:
    def addTwoNumbers(self, l1, l2):
        dummy = cur = ListNode(0)
        carry = 0
        while l1 or l2 or carry:
            if l1:
                carry += l1.val
                l1 = l1.next
            if l2:
                carry += l2.val
                l2 = l2.next
            cur.next = ListNode(carry%10)
            cur = cur.next
            carry //= 10
        return dummy.next


with open("tc2.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

listL1 = data["input"]["l1"]
listL2 = data["input"]["l2"]

for j, nums in enumerate(listL1):
    l1 = createLinkedList(listL1[j])
    l2 = createLinkedList(listL2[j])

    solutionNode = sol.addTwoNumbers(l1, l2)
    solutionList = linkedListToVector(solutionNode)
    
    if solutionList != data["expected"][j]:
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
