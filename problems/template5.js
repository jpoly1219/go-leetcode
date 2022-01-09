// template file
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

console.log("test completed")