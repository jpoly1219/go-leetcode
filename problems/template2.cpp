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

vector<ListNode> createLinkedList(vector<vector<int>> vecNums) {
    vector<ListNode> vecLL;

    for (i = 0; i < vecNums.size(); i++) {
        if (i == 0) {
            ListNode(vecNums.at(vecNums.size() - i)) node;
        } else {
            ListNode(vecNums.at(vecNums.size() - i), vecLL.back()) node;
        }

        vecLL.push_back(node);
    }

    return vecLL;
}

vector<int> linkedListToVector(ListNode* node) {
    vector<int> vecInt;
    while (node != nullptr) {
        vecInt.push_back(node->val);
        node = node->next
    }

    return vecInt;
}

// insert Solution class here

int main() {
    Solution sol;

    ifstream i("../testcase-2-add-two-numbers.json")
    json j;
    i >> j;

    vector<vector<int>> vecNums1 = j["input"]["l1"];
    vector<vector<int>> vecNums2 = j["input"]["l2"];
    vector<vector<int>> vecExpected = j["expected"];

    bool isOk = true;
    // test
    for (int i = 0; i < vecNums1.size(); i++) {
        vector<ListNode> vecL1 = createLinkedList(vecNums1.at(i));
        vector<ListNode> vecL2 = createLinkedList(vecNums2.at(i));

        ListNode* solutionNode = sol.addTwoNumbers(vecL1.front(), vecL2.front());
        vector<int> vecSolution = linkedListToVector(solutionNode)
        for (int j = 0; j < vecSolution.size(); j++) {
            if (vecSolution.at(i) != vecExpected.at(i)) {
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