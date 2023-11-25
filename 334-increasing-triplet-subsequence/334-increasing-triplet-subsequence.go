func increasingTriplet(nums []int) bool {
    minimum := math.MaxInt
    secondMinimum := math.MaxInt
    for i := 0; i < len(nums); i++ {
        if nums[i] <= minimum {
            minimum = nums[i]
        } else if nums[i] <= secondMinimum {
            secondMinimum = nums[i]
        } else {
            return true
        }
    }
    return false
}