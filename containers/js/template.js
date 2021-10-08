// template file
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

console.log("test completed")