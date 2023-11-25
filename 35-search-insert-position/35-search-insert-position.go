func searchInsert(nums []int, target int) int {
    low, high := 0, len(nums) - 1
    mid := (low + high) / 2
    for low <= high {
        if nums[mid] == target {
            return mid
        } else if nums[mid] > target {
            high = mid - 1
        } else {
            low = mid + 1
        }
        mid = (low + high) / 2
    }
    return low
}
