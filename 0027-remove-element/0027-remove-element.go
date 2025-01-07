// Approach 1 - Two Pointers
// One pointer to iterate, the other to swap
// Start from 0. If non-val, swap, and increase count
// Time: O(n) Space: O(1)
func removeElement(nums []int, val int) int {
    k := 0
    for _, num := range nums {
        if num != val {
            nums[k] = num
            k++
        }
    }
    return k
}