/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function(head) {
    let dupHead = head

    while (dupHead) {
        while (dupHead.next && dupHead.val == dupHead.next.val) {
            dupHead.next = dupHead.next.next
        }

        dupHead = dupHead.next   
    }

    return head
};
