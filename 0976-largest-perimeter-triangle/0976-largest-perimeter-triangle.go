func largestPerimeter(nums []int) int {
    return sortedNums(nums)
}

func sortedNums(nums []int) int {
    slices.Sort(nums)
    for i := len(nums)-3; i >= 0; i-- {
        j, k := i+1, len(nums) - 1
        for j < k {
            if nums[i] + nums[j] > nums[k] {
                return nums[i] + nums[j] + nums[k]
            } else {
                k--
            }
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