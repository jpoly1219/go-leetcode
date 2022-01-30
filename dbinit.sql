CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE='Asia/Seoul';

CREATE TABLE IF NOT EXISTS users (
    user_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(16) UNIQUE NOT NULL,
    fullname VARCHAR(128) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password VARCHAR(128) NOT NULL,
    profile_pic TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS problems (
    problem_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(128) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    difficulty VARCHAR(8) NOT NULL,
    description TEXT UNIQUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS attempts (
    attempt_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(16) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    lang VARCHAR(8) NOT NULL,
    code TEXT NOT NULL,
    result TEXT NOT NULL,
    output TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY(username)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS templates (
    template_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    lang VARCHAR(8) NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    template TEXT UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS testcases (
    testcase_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    testcase TEXT UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS solutions (
    solution_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    slug VARCHAR(128) UNIQUE NOT NULL,
    solution TEXT UNIQUE NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS discussions (
    discussion_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    author VARCHAR(16) UNIQUE NOT NULL,
    slug VARCHAR(128) UNIQUE NOT NULL,
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY(author)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY(slug)
            REFERENCES problems(slug)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    author VARCHAR(16) UNIQUE NOT NULL,
    discussion_id uuid,
    CONSTRAINT fk_user
        FOREIGN KEY(author)
            REFERENCES users(username)
            ON DELETE CASCADE,
    CONSTRAINT fk_discussion
        FOREIGN KEY(discussion_id)
            REFERENCES discussions(discussion_id)
            ON DELETE CASCADE,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- create users

-- create problems
INSERT INTO problems (title, slug, difficulty, description) VALUES (
    '1. Two Sum',
    '1-two-sum',
    'easy',
    $$Given an array of integers `nums` and an integer `target`, *return indices of the two numbers such that they add up to `target`.*

    You may assume that each input would have ***exactly one solution***, and you may not use the same element twice.

    You can return the answer in any order.

    <br>

    **Example 1:**

        Input: nums = [2,7,11,15], target = 9
        Output: [0,1]
        Output: Because nums[0] + nums[1] == 9, we return [0, 1].

    **Example 2:**

        Input: nums = [3,2,4], target = 6
        Output: [1,2]

    **Example 3:**

        Input: nums = [3,3], target = 6
        Output: [0,1]

    <br>

    **Constraints:**

    * `2 <= nums.length <= 104`
    * `-109 <= nums[i] <= 109`
    * `-109 <= target <= 109`
    * **Only one valid answer exists.**

    <br>

    **Follow-up:** Can you come up with an algorithm that is less than <code>O(n<sup>^</sup>2)</code> time complexity?$$
);

INSERT INTO problems (title, slug, difficulty, description) VALUES (
    '2. Add Two Numbers',
    '2-add-two-numbers',
    'medium',
    $$You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

    You may assume the two numbers do not contain any leading zero, except the number 0 itself.

    <br>

    **Example 1:**

    ![](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

        Input: l1 = [2,4,3], l2 = [5,6,4]
        Output: [7,0,8]
        Explanation: 342 + 465 = 807.

    **Example 2:**

        Input: l1 = [0], l2 = [0]
        Output: [0]

    **Example 3:**

        Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
        Output: [8,9,9,9,0,0,0,1]

    <br>

    **Constraints:**

    * The number of nodes in each linked list is in the range `[1, 100]`.
    * `0 <= Node.val <= 9`
    * It is guaranteed that the list represents a number that does not have leading zeros.$$
);

INSERT INTO problems (title, slug, difficulty, description) VALUES (
    '3. Longest Substring Without Repeating Characters',
    '3-longest-substring-without-repeating-characters',
    'medium',
    $$Given a string `s`, find the length of the **longest substring** without repeating characters.

    <br>

    **Example 1:**

        Input: s = "abcabcbb"
        Output: 3
        Explanation: The answer is "abc", with the length of 3.

    **Example 2:**

        Input: s = "bbbbb"
        Output: 1
        Explanation: The answer is "b", with the length of 1.

    **Example 3:**

        Input: s = "pwwkew"
        Output: 3
        Explanation: The answer is "wke", with the length of 3.
        Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

    **Example 4:**

        Input: s = ""
        Output: 0

    <br>

    **Constraints:**

    * <code>0 <= s.length <= 5 * 10<sup>4</sup></code>
    * `s` consists of English letters, digits, symbols and spaces.$$
);

INSERT INTO problems (title, slug, difficulty, description) VALUES (
    '4. Median of Two Sorted Arrays',
    '4-median-of-two-sorted-arrays',
    'hard',
    $$Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

    The overall run time complexity should be `O(log (m+n))`.

    <br>

    **Example 1:**

        Input: nums1 = [1,3], nums2 = [2]
        Output: 2.00000
        Explanation: merged array = [1,2,3] and median is 2.

    **Example 2:**

        Input: nums1 = [1,2], nums2 = [3,4]
        Output: 2.50000
        Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

    **Example 3:**

        Input: nums1 = [0,0], nums2 = [0,0]
        Output: 0.00000

    **Example 4:**

        Input: nums1 = [], nums2 = [1]
        Output: 1.00000

    **Example 5:**

        Input: nums1 = [2], nums2 = []
        Output: 2.00000

    <br>

    **Constraints:**

    * `nums1.length == m`
    * `nums2.length == n`
    * `0 <= m <= 1000`
    * `0 <= n <= 1000`
    * `1 <= m + n <= 2000`
    * <code>-10<sup>6</sup> <= nums1[i], nums2[i] <= 10<sup>6</sup></code>$$
);

INSERT INTO problems (title, slug, difficulty, description) VALUES (
    '5. Longest Palindromic Substring',
    '5-longest-palindromic-substring',
    'medium',
    $$Given a string `s`, return *the longest palindromic substring* in `s`.

    <br>

    **Example 1:**

        Input: s = "babad"
        Output: "bab"
        Note: "aba" is also a valid answer.

    **Example 2:**

        Input: s = "cbbd"
        Output: "bb"

    **Example 3:**

        Input: s = "a"
        Output: "a"

    **Example 4:**

        Input: s = "ac"
        Output: "a"

    <br>

    **Constraints:**

    * `1 <= s.length <= 1000`
    * `s` consist of only digits and English letters.$$
);

-- create templates
-- template 1
INSERT INTO templates (lang, slug, template) VALUES (
    'cpp',
    '1-two-sum',
    $$// template file
    #include <fstream>
    #include <iostream>
    #include <vector>
    #include "json.hpp"

    using namespace std;
    using json = nlohmann::json;

    // insert Solution class here

    int main() {
        Solution sol;

        ifstream i("../testcase-1-two-sum.json");
        json j;
        i >> j;

        vector<vector<int>> vecNums = j["input"]["nums"];
        vector<int> vecTargets = j["input"]["target"];
        vector<vector<int>> vecExpected = j["expected"];

        bool isOk = true;
        // test
        for (int i = 0; i < vecNums.size(); i++) {
            vector<int> vecSolution = sol.twoSum(vecNums.at(i), vecTargets.at(i));
            if (vecSolution != vecExpected.at(i)) {
                json output = {
                    {"result", "wrong"},
                    {"input", vecNums.at(i)},
                    {"expected", vecExpected.at(i)},
                    {"output", vecSolution}
                };
                cout << output.dump(4) << endl;
                isOk = false;
                break;
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
    }$$
);

INSERT INTO templates (lang, slug, template) VALUES (
    'java',
    '1-two-sum',
    $$// template file
    // insert Solution class here
    class Solution {
        public int[] twoSum(int[] nums, int target) {
            int[nums.length] output;
            for (int i = 0; i < nums.length; i++) {
                output[i] = nums[i] + target;
            }
            return output;
        }
    }

    public class Template {
        public static void main(String[] args) {
            Solution sol = new Solution();
            boolean isOk = true;
            for (int i = 0; i < )
        }
    }$$
);

INSERT INTO templates (lang, slug, template) VALUES (
    'js',
    '1-two-sum',
    $$// template file
    const tc = require("../testcase-1-two-sum.json")

    // insert Solution class here

    let isOk = true
    for (let i = 0; i < tc.input.nums.length; i++) {
        const answer = twoSum(tc.input.nums[i], tc.input.target[i])
        if (answer != tc.expected[i]) {
            const output = {
                "result": "wrong",
                "input": tc.input.nums[i],
                "expected": tc.expected[i],
                "output": answer
            }
            const data = JSON.stringify(output)
            console.log(data)
            isOk = false
            break
        }
    }

    if (isOk) {
        const output = {
            "result": "OK"
        }
        const data = JSON.stringify(output)
        console.log(data)
    }

    console.log("test completed")$$
);

INSERT INTO templates (lang, slug, template) VALUES (
    'py',
    '1-two-sum',
    $$# template file
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

    $$
);

INSERT INTO templates (lang, slug, template)

-- template 2
INSERT INTO templates (lang, slug, tempalte) VALUES (
    'cpp',
    '2-add-two-numbers',
    $$// template file
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
    }$$
);

INSERT INTO templates (lang, slug, tempalte) VALUES (
    'js',
    '2-add-two-numbers',
    $$// template file
    const tc = require("../testcase-2-add-two-numbers.json")

    class ListNode {
        constructor(val, next) {
            this.val = (val===undefined ? 0 : val)
            this.next = (next===undefined ? null : next)
        }
    }

    function createLinkedList(nums) {
        let head = null
        nums.forEach(num => {
            let newNode = new ListNode(num, head)
            head = newNode
        });

        return head
    }

    function linkedListToVector(node) {
        let intList = []
        while (node != null) {
            intList.push(node.val)
            node = node.next
        }

        return intList
    }

    // insert Solution class here


    let isOk = true
    const listL1 = tc.input.l1
    const listL2 = tc.input.l2
    for (let i = 0; i < listL1.length; i++) {
        const l1 = createLinkedList(listL1[i])
        const l2 = createLinkedList(listL2[i])

        const solutionNode = addTwoNumbers(l1, l2)
        const solutionList = linkedListToVector(solutionNode)

        for (let j = 0; j < solutionList.length; j++) {
            if (solutionList[j] != tc.expected[i][j]) {
                isOk = false
                break
            }   
        }
    }

    if (isOk) {
        const output = {
            "result": "OK"
        }
        const data = JSON.stringify(output)
        console.log(data)
    }

    console.log("test completed")$$
);

INSERT INTO templates (lang, slug, tempalte) VALUES (
    'py',
    '2-add-two-numbers',
    $$# template file
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
    
    $$
);

-- create testcases