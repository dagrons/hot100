package i23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	mid := (len(lists) - 1) / 2
	leftList := mergeKLists(lists[:mid+1])
	rightList := mergeKLists(lists[mid+1:])
	i, j := leftList, rightList
	var newHead, newTail *ListNode
	for i != nil || j != nil {
		if i == nil || j != nil && i.Val > j.Val {
			if newHead == nil {
				newHead, newTail = j, j
			} else {
				newTail.Next = j
				newTail = newTail.Next
			}
			j = j.Next
		} else {
			if newHead == nil {
				newHead, newTail = i, i
			} else {
				newTail.Next = i
				newTail = newTail.Next
			}
			i = i.Next
		}
	}
	return newHead
}
