package i25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	END := false                  // 是否已经结束
	prevTail := &ListNode{0, nil} // 上一段的链表尾部
	headS := prevTail             // dummy head
	nextHead := head              // 下一段链表的头
	currentHead := head           // 当前要被反转的链表部分头

	for {
		// 计算下一段将要被翻转的链表头部
		for i := 0; i < k; i++ {
			if nextHead != nil {
				nextHead = nextHead.Next
			} else {
				END = true
			}
		}
		if END {
			return headS.Next
		}
		// 反转[prevHead, nextHead)部分
		tmpHead, tmpTail := reverseList(currentHead, nextHead)
		prevTail.Next = tmpHead
		tmpTail.Next = nextHead
		currentHead = nextHead
		prevTail = tmpTail
	}
}

/*
	翻转链表
*/
func reverseList(head, nextHead *ListNode) (headS, tailS *ListNode) {
	var dfs func(t *ListNode) (headS, tailS *ListNode)
	dfs = func(t *ListNode) (head, tail *ListNode) {
		if t.Next == nextHead {
			return t, t
		}
		headS, tailS := dfs(t.Next)
		tailS.Next = t
		t.Next = nil
		tailS = tailS.Next
		return headS, tailS
	}
	return dfs(head)
}
