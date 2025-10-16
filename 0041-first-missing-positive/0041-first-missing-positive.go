func firstMissingPositive(nums []int) int {
    n := len(nums)
    
    // Swap element to its correct position
    // Continuously keep swapping the number found at its place
    for i := 0; i < n; i++ {
        for nums[i] > 0 && nums[i] <= n && nums[nums[i] - 1] != nums[i] {
            nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
        }
    }

    for i := 0; i < n; i++ {
        if nums[i] != i+1 {
            return i+1
        }
    }

    return n+1
}

func booleanArray(nums []int) int {
    n := len(nums)
    missing := make([]int, n)
    for _, val := range nums {
        if val > 0 && val < n {
            missing[val-1]++
        }
    }
    for i, val := range missing {
        if val == 0 {
            return i+1
        }
    }
    return n
}