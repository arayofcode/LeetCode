func triangleNumber(nums []int) int {
    return bruteForceApproach(nums)
}

// func sortedApproach(nums []int) int {
//     sort.Ints(nums)
//     for i := 0; i < len(nums) - 2; i++ {
//         if nums[i] + nums[i+1] <= nums[i+2] {

//         }
//     }
// }

// O(n^3)
func bruteForceApproach(nums []int) int {
    sort.Ints(nums)
    count := 0
    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            for k := j+1; k < len(nums); k++ {
                a, b, c := nums[i], nums[j], nums[k]
                if a + b > c && b + c > a && a + c > b {
                    count++
                }
            }
        }
    }
    return count
}