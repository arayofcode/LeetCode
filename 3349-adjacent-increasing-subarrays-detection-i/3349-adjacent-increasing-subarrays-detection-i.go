func hasIncreasingSubarrays(nums []int, k int) bool {
    for i := 0; i + 2*k <= len(nums); i++ {
        arr1 := nums[i:i+k]
        arr2 := nums[i+k:i+k+k]
        if increasingSubarray(arr1) && increasingSubarray(arr2) {
            return true
        }
    }
    return false
}

func increasingSubarray(nums []int) bool {
    last := nums[0]
    flag := true
    for i := 1; i < len(nums); i++ {
        if nums[i] <= last {
            flag = false
            break
        }
        last = nums[i]
    }
    return flag
}