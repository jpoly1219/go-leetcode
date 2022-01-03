// template file
const tc = require("./tc3.json")

// insert Solution class here
var lengthOfLongestSubstring = function(s) {
    var tmp = {},
        currentMaxRange = 0,
        lastRepIndex = 0,
        len = s.length,
        i;
    for(i = 0; i < len; i += 1) {
        var currentChar = s[i];
        if(typeof tmp[currentChar] !== 'undefined') {
            currentMaxRange = Math.max(currentMaxRange, i - lastRepIndex);
            lastRepIndex = Math.max(tmp[currentChar], lastRepIndex);
        } 
        tmp[currentChar] = i + 1;
        
    }
    return Math.max(currentMaxRange, i - lastRepIndex);
};


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