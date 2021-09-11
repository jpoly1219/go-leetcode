const fs = require("fs")
const tc = require("../testcase.json")

// insert Solution class here
var twoSum = function(nums, target) {
    const output = []
    nums.forEach(num => {
        output.push(num + target)
    });

    return output
};

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
        fs.writeFile("result.json", data, (err) => {
            if (err) {
                throw err
            }
        })
        isOk = false
    }
}

if (isOk) {
    const output = {
        "result": "OK"
    }
    const data = JSON.stringify(output)
    fs.writeFile("result.json", data, (err) => {
        if (err) {
            throw err
        }
    })
}

console.log("done")