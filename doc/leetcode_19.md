# 链接
https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
# 题目
```
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
示例 2：

输入：head = [1], n = 1
输出：[]
示例 3：

输入：head = [1,2], n = 1
输出：[1]
```
# 思路
- 使用快慢指针，快指针比慢指针开始时多移动n步
- 然后当快指针移到尾部，慢指针所指的位置即为删除位置
- 需要处理好边界情况，使用一个虚拟头节点可以方便移动
```go
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
```
