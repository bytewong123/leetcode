# 链接
https://leetcode-cn.com/problems/merge-two-sorted-lists/
# 题目
```
将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
```
# 思路
- 遍历两个链表，每次将值小的节点加入到当前指针所指位置即可
```go
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
```
