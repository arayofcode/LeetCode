
type cell struct {
    row, col, height int
}

type minHeap []cell

func (h minHeap) Len() int {
    return len(h)
}

func (h minHeap) Less(i, j int) bool {
    return h[i].height < h[j].height
}

func (h minHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
    *h = append(*h, x.(cell))
}

func (h *minHeap) Pop() interface{} {
    old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
    if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
    m, n := len(heightMap), len(heightMap[0])
    visited := make([][]bool, m)
    for i := range m {
        visited[i] = make([]bool, n)
    }

    h := &minHeap{}
    heap.Init(h)

    // Add boundaries to heap
    // We check minimum of this for finding height of water
    for c := 0; c < n; c++ {
        heap.Push(h, cell{0, c, heightMap[0][c]})
        heap.Push(h, cell{m-1, c, heightMap[m-1][c]})
        visited[0][c] = true
        visited[m-1][c] = true
    }
    for r := 1; r < m-1; r++ {
        heap.Push(h, cell{r, 0, heightMap[r][0]})
        heap.Push(h, cell{r, n-1, heightMap[r][n-1]})
        visited[r][0] = true
        visited[r][n-1] = true
    }
    result := 0
    water := 0
    directions := [][]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    for h.Len() > 0 {
        curCell := heap.Pop(h).(cell)

        if curCell.height > water {
            water = curCell.height
        }

        for _, dir := range directions {
            nRow, nCol := curCell.row + dir[0], curCell.col + dir[1]

            if nRow < 0 || nRow >= m || nCol < 0 || nCol >= n {
                continue
            }

            if visited[nRow][nCol] {
                continue
            }

            visited[nRow][nCol] = true

            if water > heightMap[nRow][nCol] {
                result += water - heightMap[nRow][nCol]
            }

            heap.Push(h, cell{nRow, nCol, heightMap[nRow][nCol]})
        }
    }
    return result
}
