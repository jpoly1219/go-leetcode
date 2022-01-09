// template file
const tc = require("./tc5.json")

// insert Solution class here
var longestPalindrome = function(s) {

    // Manacher's Algorithm
    
    // Justify if {string} s is totally palindrome string
    var i
    , len = Math.floor(s.length / 2) + 1
    , isTotalPalindrome = true

    for (i = 0; i < len; i++) {
    if (s[i] != s[s.length - i - 1]) {
        isTotalPalindrome = false
        break;
    }
    }
    
    if (isTotalPalindrome) return s;
    
    // preprocess, make {string} s must contain a palindrome of odd length
    s = [].join.call(s, '#')
    s = '$#' + s + '#$'

    var p = []
    , C = 1
    , R = 1
    , iMirror
    , max = 0
    , maxIndex
    
    for (i = 1; i < s.length - 1; i++) {
    iMirror = 2 * C - i
    p[i] = (R > i) ? Math.min(R - i, p[iMirror]) : 1

    while (s[i - p[i]] == s[i + p[i]]) p[i]++

    if (i + p[i] > R) {
        R = i + p[i]
        C = i
    }

    if (p[i] > max) {
        max = p[i]
        maxIndex = i
    }
    }

    return s.substr(maxIndex - max + 1, 2 * max - 1).replace(/[$#]/g, '')
}

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

console.log("test completed")