type cell struct {
    elevation int
    x int
    y int
}

// min heap containing all neighbouring elements found in path
type minHeap []*cell

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool {
	return h[i].elevation < h[j].elevation
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*cell))
}

func (h *minHeap) Pop() interface{} {
	old := *h
    n := len(old)
    item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func swimInWater(grid [][]int) int {
    result := -1
    visited := make([][]bool, len(grid))
    for i := range len(grid) {
        visited[i] = make([]bool, len(grid[0]))
    }

    directions := [][]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    h := &minHeap{}
    heap.Init(h)
    heap.Push(h, &cell{
        elevation: grid[0][0],
        x: 0,
        y: 0,
    })
    visited[0][0] = true

    for h.Len() > 0 {
        curCell := heap.Pop(h).(*cell)
        result = max(result, curCell.elevation)
        
        if curCell.x == len(grid) - 1 && curCell.y == len(grid[0]) - 1 {
            return result
        }

        for _, dir := range directions {
            newX, newY := curCell.x + dir[0], curCell.y + dir[1]
            if newX >= 0 && newX < len(grid) && newY >= 0 && newY < len(grid[0]) && !visited[newX][newY] {
                visited[newX][newY] = true
                heap.Push(h, &cell{
                    elevation: grid[newX][newY],
                    x: newX,
                    y: newY,
                })
            }
        }
    }
    return result
}