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