func longestCommonSubsequence(text1 string, text2 string) int {
    memo := make(map[[2]int]string)
    longestSubsequence := LCS(text1, text2, len(text1) - 1, len(text2) - 1, memo)
    return len(longestSubsequence)
}

func LCS(text1 string, text2 string, i int, j int, memo map[[2]int]string) string {
    key := [2]int{i, j}
    if i < 0 || j < 0 {
        return ""
    }

    if val, found := memo[key]; found {
        return val
    }

    if text1[i] == text2[j] {
        memo[key] = LCS(text1, text2, i - 1, j - 1, memo) + string(text1[i])
        return memo[key]
    }

    memo[key] = maxLengthString(LCS(text1, text2, i, j-1, memo), LCS(text1, text2, i-1, j, memo))
    return memo[key]
}

func lastRemoved(text string) string {
    if len(text) < 1 {
        return ""
    }
    return text[ : len(text) - 1]
}

func last(text string) string {
    if len(text) < 1 {
        return ""
    }
    return string(text[len(text) - 1])
}

func maxLengthString(text1, text2 string) string {
    if len(text1) > len(text2) {
        return text1
    }
    return text2
}