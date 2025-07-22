/* 
    XOR all numbers from 0 to n
    XOR all values in nums
    Only the missing value will remain
*/
func missingNumber(nums []int) int {
    // 0 xor 0 = 0 so this works as base value
    xor := 0
    for i, num := range nums {
        xor ^= i ^ num
    }
    xor ^= len(nums)
    return xor
}