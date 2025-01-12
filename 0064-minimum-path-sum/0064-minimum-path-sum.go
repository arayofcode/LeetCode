import "slices"

func minPathSum(grid [][]int) int {
    row := slices.Repeat([]int{-1}, len(grid[0]))
    memo := make([][]int, len(grid))
    for i := range memo {
        memo[i] = append([]int{}, row...)
    }
    return findPathSum(grid, 0, 0, memo)
}

func findPathSum(grid [][]int, i int, j int, memo [][]int) int {
    if memo[i][j] > 0 {
        return memo[i][j]
    }

    if i >= len(grid) - 1 && j >= len(grid[0]) - 1 {
        memo[i][j] = grid[i][j]
        return memo[i][j]
    }

    if i >= len(grid) - 1 {
        memo[i][j] = grid[i][j] + findPathSum(grid, i, j+1, memo)
        return memo[i][j]
    }
    
    if j >= len(grid[0]) - 1 {
        memo[i][j] = grid[i][j] + findPathSum(grid, i+1, j, memo)
        return memo[i][j]
    }

    memo[i][j] = grid[i][j] + min(findPathSum(grid, i+1, j, memo), findPathSum(grid, i, j+1, memo))
    return memo[i][j]
}