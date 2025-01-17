import "slices"

func productExceptSelf(nums []int) []int {
    return approachConstantSpace(nums)
}


// Approach 1 - Prefix and Suffix Arrays
func approachPrefixSuffix(nums []int) []int {
    n := len(nums)
    leftProduct := slices.Repeat([]int{1}, n)
    rightProduct := slices.Repeat([]int{1}, n)
    for i := 1; i < n; i++ {
        leftProduct[i] = leftProduct[i-1] * nums[i-1]
        rightProduct[n-1-i] = rightProduct[n-i] * nums[n-i]
    }
    for i := range leftProduct {
        leftProduct[i] *= rightProduct[i]
    }
    return leftProduct
}

// Optimization on 1: store prefix in result array
// And suffix in current array
func approachConstantSpace(nums []int) []int {
    n := len(nums)
    result := slices.Repeat([]int{1}, n)
    prefix := 1
    for i := 0; i < n; i++ {
        result[i] *= prefix
        prefix *= nums[i]
    }

    suffix := 1
    for i := n - 1; i >= 0; i-- {
        result[i] *= suffix
        suffix *= nums[i]
    }
    return result
}