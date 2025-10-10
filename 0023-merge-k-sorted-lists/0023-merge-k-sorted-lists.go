type minHeap []*ListNode

func (h minHeap) Len() int                 { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}

func (h *minHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
    results := make(minHeap, 0, len(lists))

    for _, node := range lists {
        if node != nil {
            heap.Push(&results, node)
        }
    }

    if results.Len() < 1 {
        return nil
    }

    var head *ListNode = heap.Pop(&results).(*ListNode)
    if head.Next != nil {
        heap.Push(&results, head.Next)
    }
    curr := head
    for results.Len() > 0 {
        node := heap.Pop(&results).(*ListNode)
        if node.Next != nil {
            heap.Push(&results, node.Next)
        }
        curr.Next = node
        curr = curr.Next
    }
    return head
}