func findSmallestInteger(nums []int, value int) int {
    // Store frequency of nums[i] mod value
    // This is useful because later on, we can loop 
    // through 1 to n and find the first missing number
    frequency := make([]int, value)
    for _, num := range nums {
        rem := ((num % value) + value) % value
        frequency[rem]++
    }

    val := 0
    for val < len(nums) {
        rem := val % value
        if frequency[rem] == 0 {
            return val
        }
        frequency[rem]--
        val++
    }

    return val
}