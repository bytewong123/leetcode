# 链接
https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
# 题目
```
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

进阶：

你可以设计一个只使用常数额外空间的算法来解决此问题吗？
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
```
# 思路
- 每次找到需要逆置的开头和结尾，逆置这一段，并返回新的开头和结尾
- 维护好逆置后开头和结尾的startPre和endNext的指针关系

# 时空复杂度
- 时间复杂度
O(n)
- 空间复杂度
O(1)
```go
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
```
