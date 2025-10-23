func moveZeroes(nums []int)  {
    pos := 0
    for i, num := range nums {
        if num != 0 {
            nums[i], nums[pos] = nums[pos], nums[i]
            pos++
        }
    }
}