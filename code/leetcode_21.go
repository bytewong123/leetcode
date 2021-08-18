/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    n := new(ListNode)
    head := n
    for l1 != nil && l2 != nil {
        if l1.Val < l2.Val {
            n.Next = l1
            l1 = l1.Next
        } else {
            n.Next = l2
            l2 = l2.Next
        }
        n = n.Next
    }
    for l1 != nil {
        n.Next = l1
        l1 = l1.Next
        n = n.Next
    }
    for l2 != nil {
        n.Next = l2
        l2 = l2.Next
        n = n.Next
    }
    return head.Next
}
