func largestPerimeter(nums []int) int {
    return sortedNums(nums)
}

func sortedNums(nums []int) int {
    slices.Sort(nums)
    for i := len(nums)-1; i >= 2; i-- {
        if nums[i -1] + nums[i - 2] > nums[i] {
            return nums[i] + nums[i-1] + nums[i-2]
        }
    }
    return 0
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