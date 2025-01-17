func maxProduct(nums []int) int {
    results, currMax, currMin := nums[0], nums[0], nums[0]
    for _, val := range nums[1:] {
        currMin, currMax = min(currMin * val, val, currMax * val), max(currMin * val, val, currMax * val)
        results = max(results, currMax)
    }
    return results
}