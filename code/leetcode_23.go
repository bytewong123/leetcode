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
