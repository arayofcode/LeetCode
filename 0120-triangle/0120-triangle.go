type mapKey struct {
    a int
    b int
}

func minimumTotal(triangle [][]int) int {
    return getMinSum(triangle, 0, 0, make(map[mapKey]int))
}

func getMinSum(triangle [][]int, i int, j int, memo map[mapKey]int) int {
    if i >= len(triangle) || j >= len(triangle[i]) {
        return 0
    }
    key := mapKey{i, j}
    if currSum, found := memo[key]; found {
        return currSum
    }

    result := triangle[i][j] + min(
        getMinSum(triangle, i+1, j, memo),
        getMinSum(triangle, i+1, j+1, memo),
    )
    memo[key] = result
    return result
}