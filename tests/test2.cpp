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

vector<ListNode*> createLinkedList(vector<int> vecNums) {
    vector<ListNode*> vecLL;
    ListNode* node = new ListNode();

    for (size_t i = 0; i < vecNums.size(); i++) {
        if (i == 0) {
            int nodeValue = vecNums.at(vecNums.size() - i);
            node->val = nodeValue;
        } else {
            int nodeValue = vecNums.at(vecNums.size() - i);
            ListNode* lastNode = vecLL.back();
            node->val = nodeValue;
            node->next = lastNode;
        }

        vecLL.push_back(node);
    }

    return vecLL;
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
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
    
        ListNode* dummy = new ListNode(0);
        ListNode* tail = dummy;
        int carry = 0;
        
        ListNode* ptr1 = l1;
        ListNode* ptr2 = l2;
        
        while(ptr1 && ptr2)
        {
            int sum = (ptr1->val + ptr2->val + carry);
            carry = sum/10;
            tail->next = new ListNode(sum%10);
            
            tail = tail->next;
            ptr1 = ptr1->next;
            ptr2 = ptr2->next;
        }
        
        while(ptr1)
        {
            int sum = ptr1->val + carry;
            carry = sum/10;
            tail->next = new ListNode(sum%10);
            
            tail = tail->next;
            ptr1 = ptr1->next;
        }
        
        while(ptr2)
        {
            int sum = ptr2->val + carry;
            carry = sum/10;
            tail->next = new ListNode(sum%10);
            
            tail = tail->next;
            ptr2 = ptr2->next;
        }
        
        if(carry)
            tail->next = new ListNode(carry);
        return dummy->next;
    }
};

int main() {
    Solution sol;

    ifstream i("tc2.json");
    json j;
    i >> j;

    vector<vector<int>> vecNums1 = j["input"]["l1"];
    vector<vector<int>> vecNums2 = j["input"]["l2"];
    vector<vector<int>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (int i = 0; i < vecNums1.size(); i++) {
        vector<ListNode*> vecL1 = createLinkedList(vecNums1.at(i));
        vector<ListNode*> vecL2 = createLinkedList(vecNums2.at(i));

        ListNode* solutionNode = sol.addTwoNumbers(vecL1.front(), vecL2.front());
        vector<int> vecSolution = linkedListToVector(solutionNode);
        for (size_t j = 0; j < vecSolution.size(); j++) {
            if (vecSolution.at(i) != vecExpected.at(i).at(j)) {
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