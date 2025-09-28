func largestPerimeter(nums []int) int {
    return bruteForce(nums)
}

func bruteForce(nums []int) int {
    max := 0
    for i := 0; i < len(nums) - 2; i++ {
        for j := i+1; j < len(nums) - 1; j++ {
            for k := j+1; k < len(nums); k++ {
                if (nums[i] + nums[j] > nums[k]) && (nums[i] + nums[k] > nums[j]) && (nums[k] + nums[j] > nums[i]){
                    if nums[i] + nums[j] + nums[k] > max {
                        max = nums[i] + nums[j] + nums[k]
                    }
                }
            }
        }
    }
    return max
}