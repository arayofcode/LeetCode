func findPeakGrid(mat [][]int) []int {
    return borderPeak(mat)
}

func borderPeak(mat [][]int) []int {
    m, n := len(mat), len(mat[0])

    left, right := 0, n-1
    for left <= right {
        mid := (left + right) / 2
        
        // Finding which row contains potential peak
        maxRow := 0
        for i := 0; i < m; i++ {
            if mat[i][mid] > mat[maxRow][mid] {
                maxRow = i
            }
        }

        leftVal := -1
        if mid > 0 {
            leftVal = mat[maxRow][mid-1]
        }

        rightVal := -1
        if mid < n-1 {
            rightVal = mat[maxRow][mid+1]
        }

        if mat[maxRow][mid] > leftVal && mat[maxRow][mid] > rightVal {
            return []int{maxRow, mid}
        }

        if leftVal > mat[maxRow][mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return []int{-1, -1}
}