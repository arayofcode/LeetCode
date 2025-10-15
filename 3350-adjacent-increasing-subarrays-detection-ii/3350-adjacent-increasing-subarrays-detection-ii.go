func maxIncreasingSubarrays(nums []int) int {
    n := len(nums)

    maxK := 0
    prev := 0  // Length of previous increasing subarray
    curr := 1  // Length of current increasing subarray
    
    for i := 1; i < n; i++ {
        if nums[i] > nums[i-1] {
            curr++
        } else {
            prev = curr
            curr = 1
        }
        
        // Two options:
        // 1. Use two complete adjacent subarrays: min(prev, curr)
        //      min ensures taking same subarray size
        // 2. Split current subarray in half: curr / 2
        maxK = max(maxK, max(min(prev, curr), curr/2))
    }
    
    return maxK
}