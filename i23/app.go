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
	head := &ListNode{0, nil}
	tail := head
	for i != nil && j != nil {
		if i.Val > j.Val {
			tail.Next = j
			j = j.Next
		} else {
			tail.Next = i
			i = i.Next
		}
		tail = tail.Next
	}
	if i != nil {
		tail.Next = i
	} else if j != nil {
		tail.Next = j
	}
	return head.Next
}
