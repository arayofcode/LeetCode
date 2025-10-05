func pacificAtlantic(heights [][]int) [][]int {
    rows := len(heights)
    columns := len(heights[0])

    pacific  := make([][]bool, rows)
    atlantic := make([][]bool, rows)
    for i := range rows {
        pacific[i] = make([]bool, columns)
        atlantic[i] = make([]bool, columns)
    }

    // Finding all cells that reach pacific and atlantic
    for i := 0; i < rows; i++ {
        dfs(heights, pacific, i, 0, rows, columns)
        dfs(heights, atlantic, i, columns-1, rows, columns)
    }
    for j := 0; j < columns; j++ {
        dfs(heights, pacific, 0, j, rows, columns)
        dfs(heights, atlantic, rows-1, j, rows, columns)
    }

    // Finding intersection
    result := [][]int{}
    for i := 0; i < rows; i++ {
        for j := 0; j < columns; j++ {
            if pacific[i][j] && atlantic[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func dfs(heights [][]int, visited [][]bool, i, j, rows, columns int) {
    visited[i][j] = true
    directions := [][]int{
        {-1, 0},
        {1, 0},
        {0, -1},
        {0, 1},
    }

    for _, dir := range directions {
        newR, newC := i+dir[0], j+dir[1]
        if newR >= 0 && newR < rows && newC >= 0 && newC < columns && !visited[newR][newC] && heights[newR][newC] >= heights[i][j] {
            dfs(heights, visited, newR, newC, rows, columns)
        }
    }
}