type seats []int

func (h seats) Len() int           { return len(h) }
func (h seats) Less(i, j int) bool { return h[i] < h[j] }
func (h seats) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *seats) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *seats) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type SeatManager struct {
    freeSeats *seats
}

func Constructor(n int) SeatManager {
    freeSeats := make(seats, 0, n)
    heap.Init(&freeSeats)

    for i := range n {
        heap.Push(&freeSeats, i+1)
    }

    return SeatManager{
        freeSeats: &freeSeats,
    }
}

func (this *SeatManager) Reserve() int {
    return heap.Pop(this.freeSeats).(int)
}

func (this *SeatManager) Unreserve(seatNumber int)  {
    heap.Push(this.freeSeats, seatNumber)
}
