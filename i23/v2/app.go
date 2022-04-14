package v2

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeHeap []*ListNode

func (I NodeHeap) Less(i, j int) bool  { return I[i].Val < I[j].Val }
func (I NodeHeap) Swap(i, j int)       { I[i], I[j] = I[j], I[i] }
func (I NodeHeap) Len() int            { return len(I) }
func (I *NodeHeap) Push(x interface{}) { *I = append(*I, x.(*ListNode)) }
func (I *NodeHeap) Pop() interface{}   { x := (*I)[len(*I)-1]; (*I) = (*I)[:len(*I)-1]; return x }

func mergeKLists(lists []*ListNode) *ListNode {
	q := &NodeHeap{}
	for i := range lists {
		if lists[i] == nil {
			continue
		}
		heap.Push(q, lists[i])
	}
	head := &ListNode{0, nil}
	tail := head
	for q.Len() > 0 {
		t := heap.Pop(q).(*ListNode)
		tail.Next = t
		tail = tail.Next
		if t.Next != nil {
			heap.Push(q, t.Next)
		}
	}
	return head.Next
}
