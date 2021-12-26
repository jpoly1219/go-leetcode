# template file
import json
from typing import List


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def insertNode(head, i):
    newNode = ListNode(i, head)
    head = newNode


def createLinkedList(nums):
    head = None
    for num in nums:
        insertNode(head, num)
    
    return head


def linkedListToVector(node):
    intList = []
    while (node != None):
        intList.append(node.val)
        node = node.next
    
    return intList


# insert Solution class here


with open("../testcase-2-add-two-numbers.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

for i in range(len(data["input"]["l1"])):
    l1 = createLinkedList(data["input"]["l1"])
    l2 = createLinkedList(data["input"]["l2"])

    solutionNode = sol.addTwoNumbers(l1[0], l2[0])
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
