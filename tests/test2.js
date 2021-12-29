// template file
const tc = require("./tc2.json")

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
var addTwoNumbers = function(l1, l2) {
    let values = [];
    let overflow = 0;
    while (l1.next || l2.next) {
        const value = (l1.val + l2.val + overflow) % 10;
        overflow = l1.val + l2.val + overflow >= 10 ? 1 : 0;
        values.push(value);

        l1 = l1.next || new ListNode(0);
        l2 = l2.next || new ListNode(0);
    }
    const value = (l1.val + l2.val + overflow) % 10;
    overflow = l1.val + l2.val + overflow >= 10 ? 1 : 0;
    values.push(value);
    if (overflow === 1) {
        values.push(1);
    }
    let node = undefined;
    values.reverse().forEach(item => {
        if (node === undefined) {
            node = new ListNode(item);
        } else {
            const newNode = new ListNode(item);
            newNode.next = node;
            node = newNode;
        }
    });
    return node;
};

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