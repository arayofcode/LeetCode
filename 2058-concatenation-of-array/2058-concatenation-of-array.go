func getConcatenation(nums []int) []int {
    n := len(nums)
    results := make([]int, 2 * n)
    for i := range nums {
        results[i] = nums[i]
        results[n + i] = nums[i]
    }
    return results
}