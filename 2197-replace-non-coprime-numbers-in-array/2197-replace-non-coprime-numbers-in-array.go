func replaceNonCoprimes(nums []int) []int {
    stackSize := 1 // assuming nums[0] is in position

    for i := 1; i < len(nums); i++ {
        current := nums[i]

        for stackSize > 0 {
            gcd := GCD(nums[stackSize - 1], current)
            if gcd > 1 {
                current = LCM(nums[stackSize - 1], current, gcd)
                stackSize--
            } else {
                break
            }
        }

        nums[stackSize] = current
        stackSize++
    }
    return nums[:stackSize]
}

func LCM(a int, b int, gcd int) int {
    return a * b / gcd
}

func GCD(a int, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}