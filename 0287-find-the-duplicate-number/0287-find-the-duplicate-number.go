// Floyd's Hare and Tortoise algorithm
func findDuplicate(nums []int) int {
    fast := nums[0]
    slow := nums[0]
    // Finding if there's a cycle
    for {
        fast = nums[nums[fast]]
        slow = nums[slow]
        if fast == slow {
            break
        }
    }
    // slow points to the element where 
    // Finding the element where cycle starts
    fast = nums[0]
    for fast != slow {
        fast = nums[fast]
        slow = nums[slow]
    }
    return fast
}