# C++
```
// template file
#include <fstream>
#include <iostream>
#include <vector>
#include "json.hpp"

using namespace std;
using json = nlohmann::json;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

void insertNode(ListNode** head, int i) {
    ListNode* newNode = new ListNode(i, *head);
    *head = newNode;
}

ListNode* createLinkedList(vector<int> vecNums) {
    ListNode* head = nullptr;
    for (int i = vecNums.size()-1; i >= 0; i--) {
        insertNode(&head, vecNums.at(i));
    }
    
    return head;
}

vector<int> linkedListToVector(ListNode* node) {
    vector<int> vecInt;
    while (node != nullptr) {
        vecInt.push_back(node->val);
        node = node->next;
    }

    return vecInt;
}

// insert Solution class here


int main() {
    Solution sol;

    ifstream i("../testcase-2-add-two-numbers.json");
    json j;
    i >> j;

    vector<vector<int>> vecNums1 = j["input"]["l1"];
    vector<vector<int>> vecNums2 = j["input"]["l2"];
    vector<vector<int>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (size_t i = 0; i < vecNums1.size(); i++) {
        ListNode* l1 = createLinkedList(vecNums1.at(i));
        ListNode* l2 = createLinkedList(vecNums2.at(i));

        ListNode* solutionNode = sol.addTwoNumbers(l1, l2);
        vector<int> vecSolution = linkedListToVector(solutionNode);
        for (size_t j = 0; j < vecSolution.size(); j++) {
            if (vecSolution.at(j) != vecExpected.at(i).at(j)) {
                isOk = false;
                break;
            }
        }
    }
    if (isOk) {
        json output = {
            {"result", "OK"}
        };
        cout << output.dump(4) << endl;
    }
    cout << "test completed" << endl;
    i.close();
}
```

# Java
```
```

# Javascript
```
```

# Python3
```
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


with open("../testcase-2-add-two-numbers.json", "r") as read_file:
    data = json.load(read_file)

sol = Solution()
isOk = True

listL1 = data["input"]["l1"]
listL2 = data["input"]["l2"]

for i, nums in enumerate(listL1):
    l1 = createLinkedList(listL1[i])
    l2 = createLinkedList(listL2[i])

    solutionNode = sol.addTwoNumbers(l1, l2)
    solutionList = linkedListToVector(solutionNode)
    for solution in enumerate(solutionList):
        if solutionList != data["expected"][i]:
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

```