func xorAllNums(nums1 []int, nums2 []int) int {
    result := 0
    if len(nums2) % 2 != 0 {
        for _, num1 := range nums1 {
            result ^= num1
        }
    }
    if len(nums1) % 2 != 0 {
        for _, num2 := range nums2 {
            result ^= num2
        }
    }
    return result
}