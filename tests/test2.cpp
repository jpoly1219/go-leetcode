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
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int carry = 0;
        int sum = 0;
        
        ListNode* p1 = l1;
        ListNode* p2 = l2;
        
        ListNode* res = nullptr;
        ListNode* currentNode = nullptr;
        
        while (p1 != nullptr || p2 != nullptr) {
            // calculate sum
            int sum = (p1 ? p1->val : 0) + (p2 ? p2->val : 0) + carry;
            if (sum >= 10) {
                sum = sum % 10;
                carry = 1;
            } else {
                carry = 0;
            }
            
            // make the new node
            ListNode* newNode = new ListNode(sum);
            if (!currentNode) { // if the first node hasn't been generated
                res = newNode;
                currentNode = newNode;
            } else {
                currentNode->next = newNode;
                currentNode = newNode;
            }
            
            // proceed to next, if any
            if (p1) p1 = p1->next;
            if (p2) p2 = p2->next;
        }
        
        // finally handle the left carry = 1
        if (1 == carry) {
            ListNode* newNode = new ListNode(1);
            currentNode->next = newNode;
        }
        
        return res;
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