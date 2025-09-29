func minScoreTriangulation(values []int) int {
    memo := make([][]int, len(values))
    for i := range memo {
        memo[i] = make([]int, len(values))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return dp(values, 0, len(values) - 1, memo)
}

func dp(values []int, i int, j int, memo [][]int) int {
    if j < i+2 {
        return 0
    }

    if memo[i][j] != -1 {
        return memo[i][j]
    }
    result := int(1e9)
    for k := i+1; k < j; k++ {
        cost := values[i] * values[j] * values[k] + dp(values, i, k, memo) + dp(values, k, j, memo)
        result = min(cost, result)
    }
    memo[i][j] = result
    return result
}