func replaceNonCoprimes(nums []int) []int {
    results := make([]int, 0, len(nums))
    results = append(results, nums[0])
    for i := 1; i < len(nums); i++ {
        current := nums[i]
        for len(results) > 0 {
            gcd := GCD(results[len(results) - 1], current)
            if gcd > 1 {
                current = LCM(results[len(results) - 1], current, gcd)
                results = results[:len(results) - 1]
            } else {
                break
            }
        }
        results = append(results, current)
    }
    return results
}

// Gives TLE
// func replaceNonCoprimes(nums []int) []int {
//     for i := 0; i < len(nums) - 1; {
//         gcd := GCD(nums[i], nums[i+1])
//         if gcd > 1 {
//             nums[i] = LCM(nums[i], nums[i+1], gcd)
//             nums = slices.Delete(nums, i+1, i+2)
//             if i > 0 {
//                 i--
//             }
//             continue
//         }
//         i++
//     }
//     return nums
// }

func LCM(a int, b int, gcd int) int {
    return a * b / gcd
}

func GCD(a int, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}