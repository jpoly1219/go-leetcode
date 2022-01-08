// template file
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

console.log("test completed")