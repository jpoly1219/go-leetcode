// template file
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

console.log("test completed")