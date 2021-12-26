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
        def toint(node):
            return node.val + 10 * toint(node.next) if node else 0
        def tolist(n):
            node = ListNode(n % 10)
            if n > 9:
                node.next = tolist(n / 10)
            return node
        return tolist(toint(l1) + toint(l2))


with open("tc2.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

for i in range(len(data["input"]["l1"])):
    listL1 = data["input"]["l1"]
    listL2 = data["input"]["l2"]

    for i, nums in enumerate(listL1):
        l1 = createLinkedList(listL1[i])
        l2 = createLinkedList(listL2[i])

        print("l1: ", l1.val)
        print("l2: ", l2.val)

        solutionNode = sol.addTwoNumbers(l1, l2)
        solutionList = linkedListToVector(solutionNode)
        for j, solution in enumerate(solutionList):
            if solution != data["expected"][j]:
                isOk = False;
                break;

if isOk:
    # write json to file
    output = {
        "result": "OK"
    }
    json_object = json.dumps(output, indent=4)
    print(json_object)

print("test completed")
