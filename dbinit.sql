CREATE EXTENSION IF NOT EXISTS "pgcrypto";

SET TIMEZONE='Asia/Seoul';

CREATE TABLE IF NOT EXISTS users (
    user_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR (64) UNIQUE NOT NULL,
    fullname VARCHAR (128) NOT NULL,
    email VARCHAR (256) UNIQUE NOT NULL,
    password VARCHAR(256) NOT NULL,
    profile_pic TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS problems (
    problem_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    title VARCHAR (128) UNIQUE NOT NULL,
    slug VARCHAR (128) UNIQUE NOT NULL,
    difficulty VARCHAR (16) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS templates (
    template_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    slug VARCHAR (128) NOT NULL,
    lang VARCHAR (8) NOT NULL,
    template TEXT NOT NULL,
    CONSTRAINT fk_problem 
        FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS testcases (
    testcase_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    slug VARCHAR (128) NOT NULL,
    testcase TEXT NOT NULL UNIQUE,
    CONSTRAINT fk_problem
        FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS attempts (
    attempt_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    username VARCHAR (64) NOT NULL,
    slug VARCHAR (128) NOT NULL,
    lang VARCHAR (16) NOT NULL,
    code TEXT NOT NULL,
    result VARCHAR (64) NOT NULL,
    output TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS solutions (
    solution_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    slug VARCHAR (128) UNIQUE NOT NULL,
    solution TEXT NOT NULL,
    CONSTRAINT fk_problem
        FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS discussions (
    discussion_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    author VARCHAR (64) NOT NULL,
    slug VARCHAR (128) NOT NULL,
    title VARCHAR (128) UNIQUE NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY (author) REFERENCES users (username) ON DELETE CASCADE,
    CONSTRAINT fk_problem
        FOREIGN KEY (slug) REFERENCES problems (slug) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comments (
    comment_id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    author VARCHAR (64) NOT NULL,
    discussion_id UUID DEFAULT gen_random_uuid() NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user
        FOREIGN KEY(author) REFERENCES users(username) ON DELETE CASCADE,
    CONSTRAINT fk_discussion
        FOREIGN KEY (discussion_id) REFERENCES discussions (discussion_id) ON DELETE CASCADE
);

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

-- template 2
INSERT INTO templates (lang, slug, template) VALUES (
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

INSERT INTO templates (lang, slug, template) VALUES (
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

INSERT INTO templates (lang, slug, template) VALUES (
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

-- template 3
INSERT INTO templates (lang, slug, template) VALUES (
    'cpp',
    '3-longest-substring-without-repeating-characters',
    $$// template file
    #include <fstream>
    #include <iostream>
    #include <vector>
    #include <string>
    #include "json.hpp"

    using namespace std;
    using json = nlohmann::json;

    // insert Solution class here


    int main() {
        Solution sol;

        ifstream i("../testcase-3-longest-substring-without-repeating-characters.json");
        json j;
        i >> j;

        vector<string> vecInput = j["input"];
        vector<int> vecExpected = j["expected"];

        bool isOk = true;
        // test
        for (size_t i = 0; i < vecInput.size(); i++) {
            int solution = sol.lengthOfLongestSubstring(vecInput.at(i));
            if (solution != vecExpected.at(i)) {
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
    'js',
    '3-longest-substring-without-repeating-characters',
    $$// template file
    const tc = require("../testcase-3-longest-substring-without-repeating-characters.json")

    // insert Solution class here


    let isOk = true
    const arrInput = tc.input
    const arrExpected = tc.expected
    for (let i = 0; i < arrInput.length; i++) {
        const solution = lengthOfLongestSubstring(arrInput[i])
        if (solution != arrExpected[i]) {
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
    '3-longest-substring-without-repeating-characters',
    $$# template file
    import json
    from typing import List


    # insert Solution class here



    with open("../testcase-3-longest-substring-without-repeating-characters.json", "r") as read_file:
        data = json.load(read_file)

    sol = Solution()
    isOk = True

    listInput = data["input"]
    listExpected = data["expected"]

    for i, element in enumerate(listInput):
        solution = sol.lengthOfLongestSubstring(element)
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

    $$
);

-- template 4
INSERT INTO templates (lang, slug, template) VALUES (
    'cpp',
    '4-median-of-two-sorted-arrays',
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

        ifstream i("../testcase-4-median-of-two-sorted-arrays.json");
        json j;
        i >> j;

        vector<vector<int>> vecNums1 = j["input"]["nums1"];
        vector<vector<int>> vecNums2 = j["input"]["nums2"];
        vector<double> vecExpected = j["expected"];

        bool isOk = true;
        // test
        for (size_t i = 0; i < vecNums1.size(); i++) {
            vector<int> num1 = vecNums1.at(i);
            vector<int> num2 = vecNums2.at(i);
            double solution  = sol.findMedianSortedArrays(num1, num2);
            if (solution != vecExpected.at(i)) {
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
    'js',
    '4-median-of-two-sorted-arrays',
    $$// template file
    const tc = require("../testcase-4-median-of-two-sorted-arrays.json")

    // insert Solution class here


    let isOk = true
    const arrNums1 = tc.input.nums1
    const arrNums2 = tc.input.nums2
    const arrExpected = tc.expected
    for (let i = 0; i < arrNums1.length; i++) {
        const num1 = arrNums1[i]
        const num2 = arrNums2[i]
        const solution = findMedianSortedArrays(num1, num2)
        if (solution != arrExpected[i]) {
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
    '4-median-of-two-sorted-arrays',
    $$# template file
    import json
    from typing import List


    # insert Solution class here



    with open("../testcase-4-median-of-two-sorted-arrays.json", "r") as read_file:
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

    $$
);

-- template 5
INSERT INTO templates (lang, slug, template) VALUES (
    'cpp',
    '5-longest-palindromic-substring',
    $$// template file
    #include <fstream>
    #include <iostream>
    #include <vector>
    #include <string>
    #include "json.hpp"

    using namespace std;
    using json = nlohmann::json;

    // insert Solution class here


    int main() {
        Solution sol;

        ifstream i("../testcase-5-longest-palindromic-substring.json");
        json j;
        i >> j;

        vector<string> vecInput = j["input"];
        vector<vector<string>> vecExpected = j["expected"];

        bool isOk = true;
        // test
        for (size_t i = 0; i < vecInput.size(); i++) {
            string solution  = sol.longestPalindrome(vecInput.at(i));
            vector<string> expected = vecExpected.at(i);
            if (find(expected.begin(), expected.end(), solution) == expected.end()) {
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
    'js',
    '5-longest-palindromic-substring',
    $$// template file
    const tc = require("../testcase-5-longest-palindromic-substring.json")

    // insert Solution class here


    let isOk = true
    const arrInput = tc.input
    const arrExpected = tc.expected
    for (let i = 0; i < arrInput.length; i++) {
        const expected = arrExpected[i]
        const solution = longestPalindrome(arrInput[i])
        if (!expected.includes(solution)) {
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
    '5-longest-palindromic-substring',
    $$# template file
    import json
    from typing import List


    # insert Solution class here



    with open("../testcase-5-longest-palindromic-substring.json", "r") as read_file:
        data = json.load(read_file)

    sol = Solution()
    isOk = True

    listInput = data["input"]
    listExpected = data["expected"]

    for i, element in enumerate(listInput):
        expected = listExpected[i]
        solution = sol.longestPalindrome(element)
        if solution not in expected:
            isOk = False
            break

    if isOk:
        output = {
            "result": "OK"
        }
        json_object = json.dumps(output, indent=4)
        print(json_object)

    print("test completed")$$
);

-- create testcases

INSERT INTO testcases (slug, testcase) VALUES (
    '1-two-sum',
    $${
        "input": {
            "nums": [
                [1, 2, 3, 4, 5],
                [1, 2, 3],
                [5, 5, 5, 5, 5, 5],
                [3, 2, 1],
                [0, 0, 0, 0]
            ],
            "target": [
                1, 2, 3, 4, 5
            ]
        },
        "expected": [
            [2, 3, 4, 5, 6],
            [3, 4, 5],
            [8, 8, 8, 8, 8, 8],
            [7, 6, 5],
            [5, 5, 5, 5]
        ]
    }$$
);

INSERT INTO testcases (slug, testcase) VALUES (
    '2-add-two-numbers',
    $${
        "input": {
            "l1": [
                [2, 4, 3],
                [0],
                [9, 9, 9, 9, 9, 9, 9]
            ],
            "l2": [
                [5, 6, 4],
                [0],
                [9, 9, 9, 9]
            ]
        },
        "expected": [
            [7, 0, 8],
            [0],
            [8, 9, 9, 9, 0, 0, 0, 1]
        ]
    }$$
);

INSERT INTO testcases (slug, testcase) VALUES (
    '3-longest-substring-without-repeating-characters',
    $${
        "input": [
            "abcabcbb",
            "bbbbb",
            "pwwkew"
        ],
        "expected": [
            3,
            1,
            3
        ]
    }$$
);

INSERT INTO testcases (slug, testcase) VALUES (
    '4-median-of-two-sorted-arrays',
    $${
        "input": {
            "nums1": [
                [1, 3],
                [1, 2]
            ],
            "nums2": [
                [2],
                [3, 4]
            ]
        },
        "expected": [
            2.00000,
            2.50000
        ]
    }$$
);

INSERT INTO testcases (slug, testcase) VALUES (
    '5-longest-palindromic-substring',
    $${
        "input": [
            "babad",
            "cbbd"
        ],
        "expected": [
            ["bab", "aba"],
            ["bb"]
        ]
    }$$
);

-- create solutions

INSERT INTO solutions (slug, solution) VALUES (
    '1-two-sum',
    $$### Approach 1: Brute Force

    **Algorithm**

    The brute force approach is simple. Loop through each element *x* and find if there is another value that equals to *target - x*.

    **Implementation**
    ```
    // Java
    class Solution {
        public int[] twoSum(int[] nums, int target) {
            for (int i = 0; i < nums.length; i++) {
                for (int j = i + 1; j < nums.length; j++) {
                    if (nums[j] == target - nums[i]) {
                        return new int[] { i, j };
                    }
                }
            }
            // In case there is no solution, we'll just return null
            return null;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def twoSum(self, nums: List[int], target: int) -> List[int]:
            for i in range(len(nums)):
                for j in range(i + 1, len(nums)):
                    if nums[j] == target - nums[i]:
                        return [i, j]
    ```

    **Complexity Analysis**

    - Time complexity: *O(n^2)*. For each element, we try to find its complement by looping through the rest of the array which takes *O(n)* time. Therefore, the time complexity is *O(n^2)*.

    - Space complexity: *O(1)*. The space required does not depend on the size of the input array, so only constant space is used.

    ---

    ### Approach 2: Two-pass Hash Table

    **Intuition**

    To improve our runtime complexity, we need a more efficient way to check if the complement exists in the array. If the complement exists, we need to get its index. What is the best way to maintain a mapping of each element in the array to its index? A hash table.

    We can reduce the lookup time from *O(n)* to *O(1)* by trading space for speed. A hash table is well suited for this purpose because it supports fast lookup in near constant time. I say "near" because if a collision occurred, a lookup could degenerate to *O(n)* time. However, lookup in a hash table should be amortized *O(1)* time as long as the hash function was chosen carefully.

    **Algorithm**

    A simple implementation uses two iterations. In the first iteration, we add each element's value as a key and its index as a value to the hash table. Then, in the second iteration, we check if each element's complement (*target - nums[i]*) exists in the hash table. If it does exist, we return current element's index and its complement's index. Beware that the complement must not be *nums[i]* itself!

    **Implementation**
    ```
    // Java
    class Solution {
        public int[] twoSum(int[] nums, int target) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int i = 0; i < nums.length; i++) {
                map.put(nums[i], i);
            }
            for (int i = 0; i < nums.length; i++) {
                int complement = target - nums[i];
                if (map.containsKey(complement) && map.get(complement) != i) {
                    return new int[] { i, map.get(complement) };
                }
            }
            // In case there is no solution, we'll just return null
            return null;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def twoSum(self, nums: List[int], target: int) -> List[int]:
            hashmap = {}
            for i in range(len(nums)):
                hashmap[nums[i]] = i
            for i in range(len(nums)):
                complement = target - nums[i]
                if complement in hashmap and hashmap[complement] != i:
                    return [i, hashmap[complement]] 
    ```

    **Complexity Analysis**

    - Time complexity: *O(n)*. We traverse the list containing nnn elements exactly twice. Since the hash table reduces the lookup time to *O(1)*, the overall time complexity is *O(n)*.

    - Space complexity: *O(n)*. The extra space required depends on the number of items stored in the hash table, which stores exactly *n* elements.

    ---

    ### Approach 3: One-pass Hash Table

    **Algorithm**

    It turns out we can do it in one-pass. While we are iterating and inserting elements into the hash table, we also look back to check if current element's complement already exists in the hash table. If it exists, we have found a solution and return the indices immediately.

    **Implementation**
    ```
    // Java
    class Solution {
        public int[] twoSum(int[] nums, int target) {
            Map<Integer, Integer> map = new HashMap<>();
            for (int i = 0; i < nums.length; i++) {
                int complement = target - nums[i];
                if (map.containsKey(complement)) {
                    return new int[] { map.get(complement), i };
                }
                map.put(nums[i], i);
            }
            // In case there is no solution, we'll just return null
            return null;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def twoSum(self, nums: List[int], target: int) -> List[int]:
            hashmap = {}
            for i in range(len(nums)):
                complement = target - nums[i]
                if complement in hashmap:
                    return [i, hashmap[complement]]
                hashmap[nums[i]] = i
    ```

    **Complexity Analysis**

    - Time complexity: *O(n)*. We traverse the list containing nnn elements only once. Each lookup in the table costs only *O(1)* time.

    - Space complexity: *O(n)*. The extra space required depends on the number of items stored in the hash table, which stores at most *n* elements.
    $$
);

INSERT INTO solutions (slug, solution) VALUES (
    '2-add-two-numbers',
    $$### Approach 1: Elementary Math

    **Intuition**

    Keep track of the carry using a variable and simulate digits-by-digits sum starting from the head of list, which contains the least-significant digit.

    ![](https://leetcode.com/problems/add-two-numbers/Figures/2_add_two_numbers.svg)
    *Figure 1. Visualization of the addition of two numbers: 342+465=807342 + 465 = 807342+465=807.*
    *Each node contains a single digit and the digits are stored in reverse order.*

    **Algorithm**

    Just like how you would sum two numbers on a piece of paper, we begin by summing the least-significant digits, which is the head of *l1* and *l2*. Since each digit is in the range of *0...9*, summing two digits may "overflow". For example *5 + 7 = 12*. In this case, we set the current digit to *2* and bring over the *carry = 1* to the next iteration. *carry* must be either *0* or *1* because the largest possible sum of two digits (including the carry) is *9 + 9 + 1 = 19*.

    The pseudocode is as following:

    - Initialize current node to dummy head of the returning list.
    - Initialize carry to *0*.
    - Initialize *p* and *q* to head of *l1* and *l2* respectively.
    - Loop through lists *l1* and *l2* until you reach both ends.
        - Set *x* to node *p*'s value. If *p* has reached the end of *l1*, set to *0*.
        - Set *y* to node *q*'s value. If *q* has reached the end of *l2*, set to *0*.
        - Set *sum = x + y + carry*.
        - Update *carry = sum / 10*.
        - Create a new node with the digit value of (*sum mod 10*) and set it to current node's next, then advance current node to next.
        - Advance both *p* and *q*.
    - Check if *carry = 1*, if so append a new node with digit *1* to the returning list.
    - Return dummy head's next node.

    Note that we use a dummy head to simplify the code. Without a dummy head, you would have to write extra conditional statements to initialize the head's value.

    Take extra caution of the following cases:

    | Test case                         | Explanation                                                                   |
    |-----------------------------------|-------------------------------------------------------------------------------|
    | *l1 = [0, 1]*<br>*l2 = [0, 1, 2]* | When one list is longer than the other.                                       |
    | *l1 = []*<br>*l2 = [0, 1]*        | When one list is null, which means an empty list.                             |
    | *l1 = [9, 9]*<br>*l2 = [1]*       | The sum could have an extra carry of one at the end, which is easy to forget. |

    ```
    // Java
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        ListNode dummyHead = new ListNode(0);
        ListNode p = l1, q = l2, curr = dummyHead;
        int carry = 0;
        while (p != null || q != null) {
            int x = (p != null) ? p.val : 0;
            int y = (q != null) ? q.val : 0;
            int sum = carry + x + y;
            carry = sum / 10;
            curr.next = new ListNode(sum % 10);
            curr = curr.next;
            if (p != null) p = p.next;
            if (q != null) q = q.next;
        }
        if (carry > 0) {
            curr.next = new ListNode(carry);
        }
        return dummyHead.next;
    }
    ```

    **Complexity Analysis**

    - Time complexity: *O(max(m, n))*. Assume that *m* and *n* represents the length of *l1* and *l2* respectively, the algorithm above iterates at most *max(m, n)* times.

    - Space complexity: *O(max(m, n))*. The length of the new list is at most *max(m, n) + 1*.

    **Follow up**

    What if the the digits in the linked list are stored in non-reversed order? For example:

    *(3 → 4 → 2) + (4 → 6 → 5) = 8 → 0 → 7*
    $$
);

INSERT INTO solutions (slug, solution) VALUES (
    '3-longest-substring-without-repeating-characters',
    $$### Approach 1: Brute Force

    **Intuition**

    Check all the substring one by one to see if it has no duplicate character.

    **Algorithm**

    Suppose we have a function `boolean allUnique(String substring)` which will return true if the characters in the substring are all unique, otherwise false. We can iterate through all the possible substrings of the given string `s` and call the function `allUnique`. If it turns out to be true, then we update our answer of the maximum length of substring without duplicate characters.

    Now let's fill the missing parts:

    1. To enumerate all substrings of a given string, we enumerate the start and end indices of them. Suppose the start and end indices are *i* and *j*, respectively. Then we have *0 ≤ i < j ≤ n* (here end index *j* is exclusive by convention). Thus, using two nested loops with *i* from *0* to *n − 1* and *j* from *i + 1* to *n*, we can enumerate all the substrings of *s*.

    2. To check if one string has duplicate characters, we can use a set. We iterate through all the characters in the string and put them into the `set` one by one. Before putting one character, we check if the set already contains it. If so, we return `false`. After the loop, we return `true`.

    ```
    // C++
    class Solution {
    public:
        int lengthOfLongestSubstring(string s) {
            int n = s.length();

            int res = 0;
            for (int i = 0; i < n; i++) {
                for (int j = i; j < n; j++) {
                    if (checkRepetition(s, i, j)) {
                        res = max(res, j - i + 1);
                    }
                }
            }

            return res;
        }

        bool checkRepetition(string& s, int start, int end) {
            vector<int> chars(128);

            for (int i = start; i <= end; i++) {
                char c = s[i];
                chars[c]++;
                if (chars[c] > 1) {
                    return false;
                }
            }

            return true;
        }
    };
    ```

    ```
    // Java
    public class Solution {
        public int lengthOfLongestSubstring(String s) {
            int n = s.length();

            int res = 0;
            for (int i = 0; i < n; i++) {
                for (int j = i; j < n; j++) {
                    if (checkRepetition(s, i, j)) {
                        res = Math.max(res, j - i + 1);
                    }
                }
            }

            return res;
        }

        private boolean checkRepetition(String s, int start, int end) {
            int[] chars = new int[128];

            for (int i = start; i <= end; i++) {
                char c = s.charAt(i);
                chars[c]++;
                if (chars[c] > 1) {
                    return false;
                }
            }

            return true;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def lengthOfLongestSubstring(self, s: str) -> int:
            def check(start, end):
                chars = [0] * 128
                for i in range(start, end + 1):
                    c = s[i]
                    chars[ord(c)] += 1
                    if chars[ord(c)] > 1:
                        return False
                return True

            n = len(s)

            res = 0
            for i in range(n):
                for j in range(i, n):
                    if check(i, j):
                        res = max(res, j - i + 1)
            return res
    ```

    **Complexity Analysis**

    - Time complexity: *O(n^3)*.

    To verify if characters within index range *[i, j)* are all unique, we need to scan all of them. Thus, it costs *O(j - i)* time.    

    For a given `i`, the sum of time costed by each *j ∈ [i+1, n]* is

    *∑(i+1, n) O(j − i)*

    Thus, the sum of all the time consumption is:

    *O(∑(i=0, n−1)(∑(j = i+1, n) (j − i))) = O(∑(i=0, n−1)(1 + n − i)(n − i) / 2) = O(n^3)*

    - Space complexity: *O(min(n, m))*. We need *O(k)* space for checking a substring has no duplicate characters, where *k* is the size of the `Set`. The size of the `Set` is upper bounded by the size of the string *n* and the size of the charset/alphabet *m*.

    ---

    ### Approach 2: Sliding Window

    **Algorithm**

    The naive approach is very straightforward. But it is too slow. So how can we optimize it?

    In the naive approaches, we repeatedly check a substring to see if it has duplicate character. But it is unnecessary. If a substring *s_ij*​ from index *i* to *j − 1* is already checked to have no duplicate characters. We only need to check if *s[j]* is already in the substring *s_ij*.

    To check if a character is already in the substring, we can scan the substring, which leads to an *O(n^2)* algorithm. But we can do better.

    By using HashSet as a sliding window, checking if a character in the current can be done in *O(1)*.

    A sliding window is an abstract concept commonly used in array/string problems. A window is a range of elements in the array/string which usually defined by the start and end indices, i.e. *[i, j)* (left-closed, right-open). A sliding window is a window "slides" its two boundaries to the certain direction. For example, if we slide *[i, j)* to the right by *1* element, then it becomes *[i+1, j+1)* (left-closed, right-open).

    Back to our problem. We use HashSet to store the characters in current window *[i, j)* (*j = i* initially). Then we slide the index *j* to the right. If it is not in the HashSet, we slide *j* further. Doing so until *s[j]* is already in the HashSet. At this point, we found the maximum size of substrings without duplicate characters start with index *i*. If we do this for all *i*, we get our answer.

    ```
    // C++
    class Solution {
    public:
        int lengthOfLongestSubstring(string s) {
            vector<int> chars(128);

            int left = 0;
            int right = 0;

            int res = 0;
            while (right < s.length()) {
                char r = s[right];
                chars[r]++;

                while (chars[r] > 1) {
                    char l = s[left];
                    chars[l]--;
                    left++;
                }

                res = max(res, right - left + 1);

                right++;
            }

            return res;
        }
    };
    ```

    ```
    // Java
    public class Solution {
        public int lengthOfLongestSubstring(String s) {
            int[] chars = new int[128];

            int left = 0;
            int right = 0;

            int res = 0;
            while (right < s.length()) {
                char r = s.charAt(right);
                chars[r]++;

                while (chars[r] > 1) {
                    char l = s.charAt(left);
                    chars[l]--;
                    left++;
                }

                res = Math.max(res, right - left + 1);

                right++;
            }
            return res;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def lengthOfLongestSubstring(self, s: str) -> int:
            chars = [0] * 128

            left = right = 0

            res = 0
            while right < len(s):
                r = s[right]
                chars[ord(r)] += 1

                while chars[ord(r)] > 1:
                    l = s[left]
                    chars[ord(l)] -= 1
                    left += 1

                res = max(res, right - left + 1)

                right += 1
            return res
    ```

    **Complexity Analysis**

    - Time complexity: *O(2n) = O(n)*. In the worst case each character will be visited twice by *i* and *j*.

    - Space complexity: *O(min(m, n))*. Same as the previous approach. We need *O(k)* space for the sliding window, where *k* is the size of the `Set`. The size of the `Set` is upper bounded by the size of the string *n* and the size of the charset/alphabet *m*.

    ---

    ### Approach 3: Sliding Window Optimized

    The above solution requires at most 2n steps. In fact, it could be optimized to require only n steps. Instead of using a set to tell if a character exists or not, we could define a mapping of the characters to its index. Then we can skip the characters immediately when we found a repeated character.

    The reason is that if *s[j]* have a duplicate in the range *[i, j)* with index *j'*, we don't need to increase *i* little by little. We can skip all the elements in the range *[i, j']* and let *i* to be *j' + 1* directly.

    **Java (Using HashMap)**

    ```
    // Java
    public class Solution {
        public int lengthOfLongestSubstring(String s) {
            int n = s.length(), ans = 0;
            Map<Character, Integer> map = new HashMap<>(); // current index of character
            // try to extend the range [i, j]
            for (int j = 0, i = 0; j < n; j++) {
                if (map.containsKey(s.charAt(j))) {
                    i = Math.max(map.get(s.charAt(j)), i);
                }
                ans = Math.max(ans, j - i + 1);
                map.put(s.charAt(j), j + 1);
            }
            return ans;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def lengthOfLongestSubstring(self, s: str) -> int:
            n = len(s)
            ans = 0
            # mp stores the current index of a character
            mp = {}

            i = 0
            # try to extend the range [i, j]
            for j in range(n):
                if s[j] in mp:
                    i = max(mp[s[j]], i)

                ans = max(ans, j - i + 1)
                mp[s[j]] = j + 1

            return ans
    ```

    Here is a visualization of the above code.
    [![](https://i.vimeocdn.com/video/1003183768-5623f4c10b87ca17146f1c40708c3e251a31a8550642d549f9e316d1b7bebe5e-d?mw=1000&mh=562)](https://player.vimeo.com/video/484238122)

    **Java (Assuming ASCII 128)**

    The previous implements all have no assumption on the charset of the string `s`.

    If we know that the charset is rather small, we can replace the `Map` with an integer array as direct access table.

    Commonly used tables are:

    - `int[26]` for Letters 'a' - 'z' or 'A' - 'Z'
    - `int[128]` for ASCII
    - `int[256]` for Extended ASCII

    ```
    // C++
    class Solution {
    public:
        int lengthOfLongestSubstring(string s) {
            // we will store a senitel value of -1 to simulate 'null'/'None' in C++
            vector<int> chars(128, -1);

            int left = 0;
            int right = 0;

            int res = 0;
            while (right < s.length()) {
                char r = s[right];

                int index = chars[r];
                if (index != -1 and index >= left and index < right) {
                    left = index + 1;
                }
                res = max(res, right - left + 1);

                chars[r] = right;
                right++;
            }
            return res;
        }
    };
    ```

    ```
    // Java
    public class Solution {
        public int lengthOfLongestSubstring(String s) {
            Integer[] chars = new Integer[128];

            int left = 0;
            int right = 0;

            int res = 0;
            while (right < s.length()) {
                char r = s.charAt(right);

                Integer index = chars[r];
                if (index != null && index >= left && index < right) {
                    left = index + 1;
                }

                res = Math.max(res, right - left + 1);

                chars[r] = right;
                right++;
            }

            return res;
        }
    }
    ```

    ```
    # Python3
    class Solution:
        def lengthOfLongestSubstring(self, s: str) -> int:
            chars = [None] * 128

            left = right = 0

            res = 0
            while right < len(s):
                r = s[right]

                index = chars[ord(r)]
                if index != None and index >= left and index < right:
                    left = index + 1

                res = max(res, right - left + 1)

                chars[ord(r)] = right
                right += 1
            return res
    ```

    **Complexity Analysis**

    - Time complexity: *O(n)*. Index *j* will iterate *n* times.

    - Space complexity (HashMap): *O(min(m, n))*. Same as the previous approach.

    - Space complexity (Table): *O(m)*. *m* is the size of the charset.
    $$
);

INSERT INTO solutions (slug, solution) VALUES (
    '4-median-of-two-sorted-arrays',
    $$### Approach 1: Merge Sort

    **Intuition**

    Let's start with the straightforward approach. If we put the elements of two arrays in one array `A` and arrange them in order. Assume the merged arrays has a length of `n`, then the median is:

    - `A[n / 2]`, if `n` is odd.
    - The average of `A[n / 2]` and `A[n / 2 + 1]`, if `n` is even.

    However, we do not really need to merge and sort these arrays. Note that both arrays are already sorted, so the smallest element is either the first element of `nums1` or the first element of `nums2`. Therefore, we can set two pointers `p1` and `p2` at the start of each array, then we can get the smallest element from the `nums1` and `nums2` by comparing the values `nums1[p1]` and `nums2[p2]`.

    **Algorithm**
    Get the total size of two arrays `m + n`

    If `m + n` is odd, we are looking for the `(m + n) / 2`-th element.
    If `m + n` is even, we are looking for the average of the `(m + n) / 2`-th and the `(m + n) / 2 + 1`-th elements.
    Set two pointers `p1` and `p2` at the beginning of arrays `nums1` and `nums2`.

    If both `p1` and `p2` are in bounds of the arrays, compare the values at `p1` and `p2`:

    If `nums1[p1]` is smaller than `nums2[p2]`, we move `p1` one place to the right.
    Otherwise, we move `p2` one place to the right.
    If `p1` is outside `nums1`, just move `p2` one place to the right.
    If `p2` is outside `nums2`, just move `p1` one place to the right.

    Get the target elements and calculate the median:

    If `m + n` is odd, repeat step 3 by `(m + n + 1) / 2` times and return the element from the last step.
    If `m + n` is even, repeat step 3 by `(m + n) / 2 + 1` times and return the average of the elements from the last two steps.

    **Implementation**

    ```cpp
    // cpp
    class Solution {
    public:
        int p1 = 0, p2 = 0;

        // Get the smaller value between nums1[p1] and nums2[p2] and move the pointer forward.

        int getMin(vector<int>& nums1, vector<int>& nums2) {
            if (p1 < nums1.size() && p2 < nums2.size()) {
                return nums1[p1] < nums2[p2] ? nums1[p1++] : nums2[p2++];
            } else if (p1 < nums1.size()) {
                return nums1[p1++];
            } else if (p2 < nums2.size()) {
                return nums2[p2++];
            }
            return -1;
        }

        double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
            int m = int(nums1.size()), n = int(nums2.size());

            if ((m + n) % 2 == 0) {
                for (int i = 0; i < (m + n) / 2 - 1; ++i) {
                    int _ = getMin(nums1, nums2);
                }
                return (double)(getMin(nums1, nums2) + getMin(nums1, nums2)) / 2;
            } else {
                for (int i = 0; i < (m + n) / 2; ++i) {
                    int _ = getMin(nums1, nums2);
                }
                return getMin(nums1, nums2);
            }

            return -1;
        }
    };
    ```

    **Complexity Analysis**

    Let _m_ be the size of array `nums1` and _n_ be the size of array `nums2`.

    Time complexity: _O(m + n)_

    We get the smallest element by comparing two values at `p1` and `p2`, it takes _O(1)_ to compare two elements and move the corresponding pointer to the right.
    We need to traverse half of the arrays before reaching the median element(s).
    To sum up, the time complexity is _O(m + n)_.
    Space complexity: _O(1)_

    We only need to maintain two pointers `p1` and `p2`.

    ---

    **Approach 2: Binary Search, Recursive**

    **Intuition**

    Because the inputs are sorted arrays and the problem asks for a logarithmic time limit, we strongly feel that binary search (or a similar approach) is a promising method. While we're not sure how to cast the same pattern as a normal binary search on this problem, let's go over some steps of a regular binary search and see if we can get any inspiration. (If you are not familiar with binary search, you can refer to our [Binary Search Explore Card](https://leetcode.com/explore/learn/card/binary-search/))

    Here we use binary search to find `target` in a sorted array `A`:

    Locate the middle index (element) of `A`.

    Compare the value of the middle element with `target`.

    Reduce the search space by cutting the current array in half and discarding the half which is guaranteed not to contain `target`.

    Repeat the above process until we either empty the array (move to half a the length of 0) or find `target`.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/bs.png)

    At each step, the search space is cut in half, so we can quickly get the result. Now back to this problem where we have two sorted arrays. For the sake of convenience, let's call them `A` and `B`.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2.png)

    Similarly, we can get and compare their middle values `A_mid` and `B_mid`. Without loss of generality in this example we assume `A_mid <= B_mid` initially, as shown in the yellow boxes.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/3.png)

    **What does this comparison imply?**

    It implies that we can compare sections of `A` and `B`.

    > For the rest of this article, we will use ≤ to represent the relative magnitude of values in arrays. For example, A_left ≤ A_right means that every element in A_left is no larger than any element in A_right. We also 'compare' elements in an array with a single element similarly, for example, A_left ≤ A_mid means that every element in A_left is no larger than the element A_mid.
    > This may not be the most standard way of expressing it, but is easy enough to understand.

    Recall that both arrays are sorted, so we know that:

    - _A_left ≤ A_mid_
    - _B_mid ≤ B_right_

    Combine these observations with the comparison we just made:

    _A_mid ≤ B_mid_

    We have the following result:

    _A_left ≤ A_mid ≤ B_mid ≤ B_right_

    Thus,

    _A_left ≤ B_right_

    As shown in the picture below:

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/4.png)

    Since `A` is sorted, we know that _A_left ≤ A_right_.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/5.png)

    Now we know that `A_left` is smaller than two halves: `A_right` and `B_right`. Although we still don't know where exactly these elements are, what we do know is `A_left` doesn't intersect with `A_right + B_right`! There is an invisible boundary between the `A_left` segment and the mixed segment `A_right + B_right`. As shown in the picture below, the dashed line divides all sorted elements into two halves.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/6.png)

    We can apply all the same logic to the mixed segment `A_left + B_left` and `B_right`, which also do not intersect. You can try to prove it yourself as an exercise.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/7.png)

    It looks somewhat clearer, we have clearly separated some subarrays. How do we continue to leverage this knowledge and use the cut-in-half method repeatedly?

    **The following step is the most important one.**

    Remember that we are looking for the median of `sorted A + B` which is one or two target values. We regard the index of the target value in the `sorted(A + B)` as `k`. For example:

    - If the lengths of `A` and `B` are `6` and `5`, the target index is `k = (6 + 5 + 1) / 2 = 6`, we shall look for the 6th smallest element.
    - If the lengths of `A` and `B` are `6` and `6`, the target indexes are `k = (6 + 6) / 2 = 6` and `k + 1 = 7`, we shall look for the 6th and the 7th smallest elements.

    Depending on whether the total number of elements is odd or even, we need the `kth` (and maybe the `(k+1)th`) elements. What matters is that we set an index `k` at the beginning and we want to find the `kth` smallest element using the Binary Search-like algorithm discussed previously (for convenience, we will discuss only the `kth` element for now).

    However, during the Binary Search-like algorithm, we keep removing one half of an array, so the index `k` might not stay unchanged. Suppose we removed 3 elements that are smaller than the original `kth` smallest element, we shall look for the `(k−3)th` smallest element from the **remaining** arrays.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/exp_1.png)

    More specifically:

    If `k` is larger than half the total number of elements in `sorted(A + B)`, it means that the `kth` element is in the second (larger) half of `sorted(A + B)`, thus `A_left` (or `B_left`, the smaller of the two smaller sections according to the comparison) is guaranteed not to contain this element, and we can safely cut this half, and reduce `k` by the length of the removed half.

    If `k` is not larger than half the total number of elements in `sorted(A + B)`, it means that the `kth` element is in the first (smaller) half of `sorted(A + B)`, thus `B_right` (or `A_right`, the larger of the two larger sections according to the comparison) is guaranteed not to contain this element, and we can safely discard it. Note that we don't need to modify `k` this time, since we removed one larger half that doesn't affect the order of the `kth` smallest element.

    We can continue our search like above in the remaining arrays. The long arrow that starts from the bottom and points to the top-left indicates that we are repeating the process. Once we cut off part of either `A` or `B`, we regard the remaining arrays as modified `A` and `B` and restart this algorithm. Note that the following picture represents one case only: we consider the case that `a_value < b_value`, thus we remove either the smaller half of `A` or the larger half of `B`. If the comparison result is `a_value >= b_value`, we shall remove either the smaller half of `B` or the larger half of `A`.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/9.png)

    That's it. We cut one of the two arrays in half at each step, so this approach has a logarithmic time complexity which we will discuss in detail later.

    > One more thing!

    In the previous picture, we repeat all processes using the modified arrays, but this is just for the sake of understanding. We won't create copies of two arrays repeatedly, because that would introduce a linear time complexity at least. Instead, we just treat a part of the original array as the modified array for the next step, so that we can repeat the process on the original array without making any duplication. To do this, we need to maintain four pointers, two pointers for each array, e.g., `a_start` and `a_end` represent an inclusive range `[a_start, a_end]` of `A`.

    **Algorithm**

    Let's define a function that helps us find the `kth` smallest element from two inclusive ranges `[a_start, a_end]` and `[b_start, b_end]` from arrays `A` and `B`.

    If the range (for example, a range of `A`) is empty, in other words `a_start > a_end`, it means all elements in `A` are passed, we just return the `(k - a_start)`-th element from the other array `B`. Vice versa if `b_start > b_end`.

    Otherwise, get the middle indexes of the two ranges: `a_index = (a_start + a_end) / 2`, `b_index = (b_start + b_end) / 2`.

    Get the middle values of the two ranges: `a_value = A[a_index]`, `b_value = B[b_index]`.

    Cut one array in half, according to:

    - If `a_index + b_index < k`, cut one smaller half.
    - If `a_value < b_value`, cut the smaller half of `A`.
    - Otherwise, cut the smaller half of `B`.
    - Otherwise, cut one larger half.
    - If `b_value < a_value`, cut the larger half of `B`.
    - Otherwise, cut the larger half of `A`.
        Repeat step 1 using the new starting and ending indexes of `A` and `B`.

    Then we move on to find the median elements, and get the length of both arrays `na = len(A)` and `nb = len(B)`.

    If the total number of elements in `A` and `B` is odd, we just use the above function to find the middle element, that is `k = (na + nb) / 2`.
    Otherwise, we use the function to find two middle elements: `k = (na + nb) / 2` - 1 and `k = (na + nb) / 2`, and return their average.

    **Implementation**

    ```cpp
    class Solution {
    public:
        double findMedianSortedArrays(vector<int>& A, vector<int>& B) {
            int na = int(A.size()), nb = int(B.size());
            int n = na + nb;
            if (n % 2) {
                return solve(A, B, n / 2, 0, na - 1, 0, nb - 1);
            } else {
                return 1.0 * (solve(A, B, n/2 - 1, 0, na - 1, 0, nb - 1) + solve(A, B, n/2, 0, na - 1, 0, nb - 1)) / 2;
            }
        }
        int solve(vector<int>& A, vector<int>& B, int k, int aStart, int aEnd, int bStart, int bEnd) {
            // If the segment of on array is empty, it means we have passed all
            // its element, just return the corresponding element in the other array.
            if (aEnd < aStart) {
                return B[k - aStart];
            }
            if (bEnd < bStart) {
                return A[k - bStart];
            }

            // Get the middle indexes and middle values of A and B.
            int aIndex = (aStart + aEnd) / 2, bIndex = (bStart + bEnd) / 2;
            int aValue = A[aIndex], bValue = B[bIndex];

            // If k is in the right half of A + B, remove the smaller left half.
            if (aIndex + bIndex < k) {
                if (aValue > bValue) {
                    return solve(A, B, k, aStart, aEnd, bIndex + 1, bEnd);
                } else {
                    return solve(A, B, k, aIndex + 1, aEnd, bStart, bEnd);
                }
            }
            // Otherwise, remove the larger right half.
            else {
                if (aValue > bValue) {
                    return solve(A, B, k, aStart, aIndex - 1, bStart, bEnd);
                } else {
                    return solve(A, B, k, aStart, aEnd, bStart, bIndex - 1);
                }
            }
            return -1;
        }
    };
    ```

    **Complexity Analysis**

    Let _m_ be the size of array `nums1` and _n_ be the size of array `nums2`.

    Time complexity: _O(log⁡(m⋅n))_

    At each step, we cut one half off from either `nums1` or `nums2`. If one of the arrays is emptied, we can directly get the target from the other array in a constant time. Therefore, the total time spent depends on when one of the arrays is cut into an empty array.
    In the worst-case scenario, we may need to cut both arrays before finding the target element.
    One of the two arrays is cut in half at each step, thus it takes logarithmic time to empty an array. The time to empty two arrays are independent of each other.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/tc.png)

    Therefore, the time complexity is `O(log⁡m+log⁡n)`. `O(log⁡m+log⁡n)=O(log⁡(m⋅n))`

    Space complexity: `O(log⁡m+log⁡n)`

    Similar to the analysis on time complexity, the recursion steps depend on the number of iterations before we cut an array into an empty array. In the worst-case scenario, we need `O(log⁡m+log⁡n)` recursion steps.

    However, during the recursive self-call, we only need to maintain 4 pointers: `a_start`, `a_end`, `b_start` and `b_end`. The last step of the function is to call itself, so if tail call optimization is implemented, the call stack always has _O(1)_ records.

    Please refer to Tail Call for more information on tail call optimization.

    **Approach 3: A Better Binary Search**

    **Intuition**

    Recall the previous approach where we perform a binary search over the 'merged' array consisting of nums1 and nums2, resulting in a time complexity of O(log⁡(m⋅n))O(\log(m \cdot n))O(log(m⋅n)). We could further improve the algorithm by performing the binary search only on the smaller array of nums1 and nums2, thus the time complexity is reduced to O(log⁡(min⁡(m,n)))O(\log(\min(m, n)))O(log(min(m,n))).

    The main idea is similar to approach 2, where we need to find a point of partition in both arrays such that the maximum of the smaller half is less than or equal to the minimum of the larger half.

    However, instead of partitioning over the merged arrays, we can only focus on partitioning the smaller array (let's call this array A). Suppose the partition index is partitionA, we specify that the smaller half contains (m + n + 1) / 2 elements, and we can use this feature to our advantage by directly making partitionB equal to (m + n + 1) / 2 - partitionA, thus the smaller halves of both arrays always contain a total of (m + n + 1) / 2 elements, as shown in the picture below.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2_0.png)

    The next step is to compare these edge elements.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2_1.png)

    If both maxLeftA <= minRightB and maxLeftB <= minRightA hold, it means that we have partitioned arrays at the correct place.

    The smaller half consists of two sections A_left and B_left
    THe larger half consists of two sections A_right and B_right
    We just need to find the maximum value from the smaller half as max(A[maxLeftA], B[maxLeftB]) and the minimum value from the larger half as min(A[minRightA], B[minRightB]). The median value depends on these four boundary values and the total length of the input arrays and we can compute it by situation.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2_2.png)

    If maxLeftA > minRightB, it implies that maxLeftA is too large to be in the smaller half and we should look for a smaller partition value of A.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2_3.png)

    Otherwise, it denotes that minRightA is too small to be in the larger half and we should look for a larger partition value of A.

    ![](https://leetcode.com/problems/median-of-two-sorted-arrays/Figures/4/2_4.png)

    Algorithm
    Assuming nums1 to be the smaller array (If nums2 is smaller, we can swap them). Let m, n represent the size of nums1 and nums2, respectively.

    Define the search space for the partitioning index partitionA by setting boundaries as left = 0 and right = m.

    While left <= right holds, do the following.

    Compute the partition index of nums1 as partitionA = (left + right) / 2. Consequently, the partition index of nums2 is (m + n + 1) / 2 - partitionA.

    Obtain the edge elements:

    Determine the maximum value of the section A_left as maxLeftA = nums1[partitionA - 1]. If partitionA - 1 < 0, set it as maxLeftA = float(-inf).
    Determine the minimum value of the section A_right as minRightA = nums1[partitionA]. If partitionA >= m, set it as minRightA = float(inf).
    Determine the maximum value of the section B_left as maxLeftB = nums2[partitionB - 1]. If partitionB - 1 < 0, set it as maxLeftB = float(-inf).
    Determine the maximum value of the section B_right as minRightB = nums2[partitionB]. If partitionB >= n, set it as minRightB = float(inf).
    Compare and recalculate: Compare maxLeftA with minRightB and maxLeftB with minRightA.

    If maxLeftA > minRightB, it means the maxLeftA is too large to be in the smaller half, so we update right = partitionA - 1 to move to the left half of the search space.
    If maxLeftB > minRightA, it means that we are too far on the left side for partitionA and we need to go to the right half of the search space by updating left = partitionA + 1.
    Repeat step 4.

    When both maxLeftA <= minRightB and maxLeftB <= minRightA are true:

    If (m + n) % 2 = 0, the median value is the average of the maximum value of the smaller half and the minimum value of the larger half, given by answer = (max(maxLeftA, maxLeftB) + min(minRightA, minRightB)) / 2.
    Otherwise, the median value is the maximum value of the smaller half, given by answer = max(maxLeftA, maxLeftB).

    **Implementation**

    ```cpp
    class Solution {
    public:
        double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
            if (nums1.size() > nums2.size()) {
                return findMedianSortedArrays(nums2, nums1);
            }

            int m = nums1.size(), n = nums2.size();
            int left = 0, right = m;

            while (left <= right) {
                int partitionA = (left + right) / 2;
                int partitionB = (m + n + 1) / 2 - partitionA;

                int maxLeftA = (partitionA == 0) ? INT_MIN : nums1[partitionA - 1];
                int minRightA = (partitionA == m) ? INT_MAX : nums1[partitionA];
                int maxLeftB = (partitionB == 0) ? INT_MIN : nums2[partitionB - 1];
                int minRightB = (partitionB == n) ? INT_MAX : nums2[partitionB];

                if (maxLeftA <= minRightB && maxLeftB <= minRightA) {
                    if ((m + n) % 2 == 0) {
                        return (max(maxLeftA, maxLeftB) + min(minRightA, minRightB)) / 2.0;
                    } else {
                        return max(maxLeftA, maxLeftB);
                    }
                } else if (maxLeftA > minRightB) {
                    right = partitionA - 1;
                } else {
                    left = partitionA + 1;
                }
            }

            return 0.0;
        }
    };
    ```

    **Complexity Analysis**

    Let mmm be the size of array nums1 and nnn be the size of array nums2.

    Time complexity: O(log⁡(min⁡(m,n)))O(\log(\min(m, n)))O(log(min(m,n)))

    We perform a binary search over the smaller array of size min⁡(m,n)\min(m, n)min(m,n).
    Space complexity: O(1)O(1)O(1)

    The algorithm only requires a constant amount of additional space to store and update a few parameters during the binary search.$$
);

INSERT INTO solutions (slug, solution) VALUES (
    '5-longest-palindromic-substring',
    $$### Approach 1: Longest Common Substring

    **Common mistake**

    Some people will be tempted to come up with a quick solution, which is unfortunately flawed (however can be corrected easily):

    > Reverse *S* and become *S'*. Find the longest common substring between *S* and *S'*, which must also be the longest palindromic substring.

    This seemed to work, let’s see some examples below.

    For example, *S* = "caba", *S'* = "abac".

    The longest common substring between *S* and *S'* is "aba", which is the answer.

    Let’s try another example: *S* = "abacdfgdcaba", *S'* = "abacdgfdcaba".

    The longest common substring between *S* and *S'* is "abacd". Clearly, this is not a valid palindrome.

    **Algorithm**

    We could see that the longest common substring method fails when there exists a reversed copy of a non-palindromic substring in some other part of *S*. To rectify this, each time we find a longest common substring candidate, we check if the substring’s indices are the same as the reversed substring’s original indices. If it is, then we attempt to update the longest palindrome found so far; if not, we skip this and find the next candidate.

    This gives us an *O(n^2)* Dynamic Programming solution which uses *O(n^2)* space (could be improved to use *O(n)* space). Please read more about Longest Common Substring [here](http://en.wikipedia.org/wiki/Longest_common_substring).

    ---

    **Approach 2: Brute Force**

    The obvious brute force solution is to pick all possible starting and ending positions for a substring, and verify if it is a palindrome.

    Complexity Analysis

    - Time complexity: *O(n^3)*. Assume that *n* is the length of the input string, there are a total of *nC2 = n(n−1) / 2​* such substrings (excluding the trivial solution where a character itself is a palindrome). Since verifying each substring takes *O(n)* time, the run time complexity is *O(n^3)*.

    - Space complexity: *O(1)*.

    ---

    ### Approach 3: Dynamic Programming

    To improve over the brute force solution, we first observe how we can avoid unnecessary re-computation while validating palindromes. Consider the case "ababa". If we already knew that "bab" is a palindrome, it is obvious that "ababa" must be a palindrome since the two left and right end letters are the same.

    We define *P(i, j)* as following:

    *P(i, j) = {true, if the substring Si … Sj is a palindrome. false,otherwise.}*

    Therefore,

    *P(i, j) = (P(i+1, j−1) and Si == Sj)*

    The base cases are:

    *P(i, i) = true*

    *P(i, i+1) = (Si ​== Si+1​)*

    This yields a straight forward DP solution, which we first initialize the one and two letters palindromes, and work our way up finding all three letters palindromes, and so on...

    **Complexity Analysis**

    - Time complexity: *O(n^2)*. This gives us a runtime complexity of *O(n^2)*.

    - Space complexity: *O(n^2)*. It uses *O(n^2)* space to store the table.

    **Additional Exercise**

    Could you improve the above space complexity further and how?

    ---

    ### Approach 4: Expand Around Center

    In fact, we could solve it in *O(n^2)* time using only constant space.

    We observe that a palindrome mirrors around its center. Therefore, a palindrome can be expanded from its center, and there are only *2n - 1* such centers.

    You might be asking why there are *2n - 1* but not *n* centers? The reason is the center of a palindrome can be in between two letters. Such palindromes have even number of letters (such as "abba") and its center are between the two 'b's.

    ```
    // Java
    public String longestPalindrome(String s) {
        if (s == null || s.length() < 1) return "";
        int start = 0, end = 0;
        for (int i = 0; i < s.length(); i++) {
            int len1 = expandAroundCenter(s, i, i);
            int len2 = expandAroundCenter(s, i, i + 1);
            int len = Math.max(len1, len2);
            if (len > end - start) {
                start = i - (len - 1) / 2;
                end = i + len / 2;
            }
        }
        return s.substring(start, end + 1);
    }

    private int expandAroundCenter(String s, int left, int right) {
        int L = left, R = right;
        while (L >= 0 && R < s.length() && s.charAt(L) == s.charAt(R)) {
            L--;
            R++;
        }
        return R - L - 1;
    }
    ```

    **Complexity Analysis**

    - Time complexity: *O(n^2)*. Since expanding a palindrome around its center could take *O(n)* time, the overall complexity is *O(n^2)*.

    - Space complexity: O(1).

    ---

    ### Approach 5: Manacher's Algorithm

    There is even an *O(n)* algorithm called Manacher's algorithm, explained [here in detail](https://en.wikipedia.org/wiki/Longest_palindromic_substring#Manacher's_algorithm). However, it is a non-trivial algorithm, and no one expects you to come up with this algorithm in a 45 minutes coding session. But, please go ahead and understand it, I promise it will be a lot of fun.
    $$
);