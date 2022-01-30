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
INSERT INTO templates (lang, slug, template) VALUES (

);

INSERT INTO templates (lang, slug, template) VALUES (

);

INSERT INTO templates (lang, slug, template) VALUES (

);

INSERT INTO templates (lang, slug, template) VALUES (

);

INSERT INTO templates (lang, slug, template) VALUES (

);
-- create testcases