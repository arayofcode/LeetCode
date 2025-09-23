func compareVersion(version1 string, version2 string) int {
    v1 := split(version1)
    v2 := split(version2)

    if len(v1) < len(v2) {
        v1 = paddedSlice(v1, len(v2))
    } else if len(v2) < len(v1) {
        v2 = paddedSlice(v2, len(v1))
    }

    for i := 0; i < len(v1); i++ {
        if v1[i] < v2[i] {
            return -1
        }
        if v1[i] > v2[i] {
            return 1
        }
    }
    return 0
}

func split(s string) []int {
    nums := []int{}
    start := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '.' {
            nums = append(nums, toInt(s[start:i]))
            start = i+1
        }
    }
    if start < len(s) {
        nums = append(nums, toInt(s[start:]))
    }
    return nums
}

func paddedSlice(s []int, length int) []int {
    for i := len(s); i < length; i++ {
        s = append(s, 0)
    }
    return s
}

func toInt(str string) (result int) {
    for _, chr := range str {
        result = result * 10 + int(chr - '0')
    }
    return result
}