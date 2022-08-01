/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {    
    let f = listToArray(l1)
    let sec = listToArray(l2)
    
    f = BigInt(f.reverse().join(''))
    sec = BigInt(sec.reverse().join(''))   
    
    const sum = (f + sec).toString().split('').reverse()
    const result = arrayToList(sum)
    
    return result
};

const listToArray = (list) => {
    let helperList = list
    let arr = []
    
    while (helperList) {
        arr.push(helperList.val)
        helperList = helperList.next
    }
    
    return arr
}

const arrayToList = (arr) => {
    let list = new ListNode(arr[0]);
    let head = list
    
    for (let i = 1; i < arr.length; i++) {
        head.next = new ListNode(arr[i])
        head = head.next
    }
    
    return list
}
