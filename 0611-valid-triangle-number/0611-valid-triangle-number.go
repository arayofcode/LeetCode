func triangleNumber(nums []int) int {
    return twoPointersApproach(nums)
}

// O(n log n + n^2) = O(n^2)
func twoPointersApproach(nums []int) int {
    sort.Ints(nums)
    count := 0

    // when k is the largest, find count of all possible triplets
    for k := 2; k < len(nums); k++ {
        i, j := 0, k-1
        for i < j {
            if nums[i] + nums[j] > nums[k] {
                count += j - i
                j--
            } else {
                i++
            }
        }
    }
    
    return count
}

// O(n^3)
func bruteForceApproach(nums []int) int {
    sort.Ints(nums)
    count := 0
    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            for k := j+1; k < len(nums); k++ {
                a, b, c := nums[i], nums[j], nums[k]
                if a + b > c && b + c > a && a + c > b {
                    count++
                }
            }
        }
    }
    return count
}