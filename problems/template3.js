// template file
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

console.log("test completed")