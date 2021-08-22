# 链接
https://leetcode-cn.com/problems/merge-k-sorted-lists/
# 题目
```
给你一个链表数组，每个链表都已经按升序排列。

请你将所有链表合并到一个升序链表中，返回合并后的链表。

 

示例 1：

输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6
示例 2：

输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]

```
# 思路
- 首先将所有n个链表的头节点都插入到堆中
- 不断从堆中弹出元素，插入到当前节点，若弹出元素的next不为空，则继续插入到堆中
- 当弹出的元素为空，结束

# 时空复杂度
- 时间复杂度
假设链表数为n，总共m个节点
每次插入元素需要O(logn)，弹出元素O(1)，m个节点入堆时间复杂度O(mlogn)
- 空间复杂度
堆需要O(n)
```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    heap := make([]*ListNode, 0, len(lists))
    for i := 0; i < len(lists); i++ {
        if lists[i] == nil {
            continue
        }
        push(&heap, lists[i])
    }
    n := new(ListNode)
    head := n
    for {
        poped := pop(&heap)
        if poped == nil {
            break
        }
        n.Next = poped
        n = n.Next
        if poped.Next != nil {
            push(&heap, poped.Next)
        }
    }
    return head.Next
}

func push(heap *[]*ListNode, n *ListNode) {
    *heap = append(*heap, n)
    up(heap, len(*heap) - 1)
}

func pop(heap *[]*ListNode) *ListNode {
    if len(*heap) == 0 {
        return nil
    }
    poped := (*heap)[0]
    (*heap)[0] = (*heap)[len(*heap) - 1]
    *heap = (*heap)[:len(*heap) - 1]
    down(heap, 0)
    return poped
}

func up(heap *[]*ListNode, cur int) {
    if cur <= 0 {
        return
    }
    parent := (cur - 1) >> 1
    if (*heap)[cur].Val < (*heap)[parent].Val {
        (*heap)[cur], (*heap)[parent] = (*heap)[parent], (*heap)[cur]
        up(heap, parent)
    }
}

func down(heap *[]*ListNode, cur int) {
    left, right := (cur << 1) + 1, (cur << 1) + 2
    min := cur
    if left < len(*heap) && (*heap)[left].Val < (*heap)[min].Val {
        min = left
    }
    if right < len(*heap) && (*heap)[right].Val < (*heap)[min].Val {
        min = right
    }
    if min != cur {
        (*heap)[cur], (*heap)[min] = (*heap)[min], (*heap)[cur]
        down(heap, min)
    }
}
```
