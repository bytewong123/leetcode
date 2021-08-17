/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dumb := new(ListNode)
    dumb.Next = head
    node := dumb
    slow := node
    fast := node
    for n > 0 {
        fast = fast.Next
        n--
    }
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return node.Next
}
