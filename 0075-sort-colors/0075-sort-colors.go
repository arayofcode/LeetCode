func sortColors(nums []int)  {
    threePointers(nums)
}

// This is two-pass approach, but completes in O(n) time and O(1) space
func countingSort(nums []int) {
    count := []int{0, 0, 0}
    for _, num := range nums {
        count[num]++
    }
    k := 0
    for i, value := range count {
        for _ = range value {
            nums[k] = i
            k++
        }
    }
}

// Three pointers:
// Before left -> all zeroes
// Beyond right -> all twos
// between left and current -> all ones
// between current and right -> unprocessed ones
// O(n) time, O(1) space, and one pass
func threePointers(nums []int) {
    left, right := 0, len(nums) - 1
    for curr := left; curr <= right; {
        if nums[curr] == 0 {
            nums[curr], nums[left] = nums[left], nums[curr]
            left++
            curr++
        } else if nums[curr] == 2 {
            nums[curr], nums[right] = nums[right], nums[curr]
            right--
        } else {
            curr++
        }
    }
}