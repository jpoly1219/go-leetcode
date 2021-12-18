### Approach 1: Brute Force

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