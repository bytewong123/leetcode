/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    dummy := new(ListNode)
    dummy.Next = head
    startPrev, endNext := dummy, dummy
    start, end := head, head
    i := 1
    for {
        for i < k {
            if end == nil {
                return dummy.Next
            }
            end = end.Next
            i++
        } 
        if end == nil {
            break
        }
        endNext = end.Next
        newStart, newEnd := reverse(start, end)
        startPrev.Next = newStart
        newEnd.Next = endNext
        startPrev = newEnd
        start, end = endNext, endNext
        i = 1
    }
    return dummy.Next
}

func reverse(start, end *ListNode) (*ListNode, *ListNode) {
    var pre *ListNode
    n := start
    for {
        next := n.Next
        n.Next = pre
        if n == end {
            break
        }
        pre = n
        n = next
    }
    return end, start
}
