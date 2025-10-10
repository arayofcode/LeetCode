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
    var head *ListNode = nil
    var tail *ListNode = nil
    results := make(minHeap, 0, len(lists))
    heap.Init(&results)
    for _, node := range lists {
        if node != nil {
            heap.Push(&results, node)
        }
    }

    for results.Len() > 0 {
        node := heap.Pop(&results).(*ListNode)
        if node.Next != nil {
            heap.Push(&results, node.Next)
        }

        if head == nil {
            head = node
            tail = head
        } else {
            tail.Next = node
            tail = tail.Next
        }
    }
    return head
}