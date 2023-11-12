func productExceptSelf(nums []int) []int {
    n := len(nums)
    var result = make([]int, n)
    left_product, right_product := 1, 1
    for i := 0; i < n; i++ {
        result[i] = left_product
        left_product *= nums[i]
    }
    for i := n - 1; i >= 0; i-- {
        result[i] *= right_product
        right_product *= nums[i]
    }
    return result
}

func productExceptSelfArray(nums []int) []int {
    // Implementation using slices storing prefix and suffix
    n := len(nums)
    var prefix = make([]int, n)
    var suffix = make([]int, n)
    var result = make([]int, n)
    
    for i := 0; i < n; i++ {
        if i == 0 {
            prefix[i] = 1
            suffix[n - 1] = 1
        } else {
            prefix[i] = nums[i - 1] * prefix[i - 1]
            suffix[n - 1 - i] = nums[n - i] * suffix[n - i]
        }
    }
    for i := 0; i < n; i++ {
        result[i] = prefix[i] * suffix[i]
    }
    return result
}
