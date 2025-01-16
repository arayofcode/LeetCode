func xorAllNums(nums1 []int, nums2 []int) int {
    m, n := len(nums1), len(nums2)
    evenM, evenN := m % 2 == 0, n % 2 == 0
    if evenM && evenN {
        return 0
    } else if evenM {
        result := 0
        for _, num1 := range nums1 {
            result ^= num1
        }
        return result
    } else if evenN {
        result := 0
        for _, num2 := range nums2 {
            result ^= num2
        }
        return result
    }
    i, j, result := 0, 0, 0
    for i < m || j < n {
        if i < m {
            result ^= nums1[i]
            i++
        }
        if j < n {
            result ^= nums2[j]
            j++
        }
    }
    return result
}