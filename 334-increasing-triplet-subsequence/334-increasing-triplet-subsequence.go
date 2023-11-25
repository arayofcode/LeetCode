func increasingTriplet(nums []int) bool {
    minimum := math.MaxInt
    secondMinimum := minimum
    n := len(nums)
    for i := 0; i < n; i++ {
        num := nums[i]
        if num <= minimum {
            minimum = num
        } else if num <= secondMinimum {
            secondMinimum = num
        } else {
            return true
        }
    }
    return false
}