func minimumTotal(triangle [][]int) int {
    for i := len(triangle) - 2; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            if triangle[i+1][j] < triangle[i+1][j+1] {
                triangle[i][j] += triangle[i+1][j]
            } else {
                triangle[i][j] += triangle[i+1][j+1]
            }
        }
    }
    return triangle[0][0]
}

type mapKey struct {
    a int
    b int
}

func sumTillNApproach(triangle [][]int, i int, j int, memo map[mapKey]int) int {
    if i >= len(triangle) || j >= len(triangle[i]) {
        return 0
    }
    key := mapKey{i, j}
    if currSum, found := memo[key]; found {
        return currSum
    }

    result := triangle[i][j] + min(
        sumTillNApproach(triangle, i+1, j, memo),
        sumTillNApproach(triangle, i+1, j+1, memo),
    )
    memo[key] = result
    return result
}

func findSumApproach(triangle [][]int, currSum int, i int, j int) int {
    if i >= len(triangle) || j >= len(triangle[i]) {
        return currSum
    }
    currSum += triangle[i][j]
    return min(findSumApproach(triangle, currSum, i + 1, j), findSumApproach(triangle, currSum, i + 1, j + 1))
}
