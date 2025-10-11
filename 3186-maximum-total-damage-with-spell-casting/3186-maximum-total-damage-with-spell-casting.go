func maximumTotalDamage(power []int) int64 {
    freq := make(map[int]int64)
    for _, p := range power {
        freq[p] += int64(p)
    }

    // get sorted keys
    keys := make([]int, 0, len(freq))
    for key := range freq {
        keys = append(keys, key)
    }
    
    slices.Sort(keys)
    dp := make([]int64, len(keys))
    dp[0] = freq[keys[0]]
    for i := 1; i < len(keys); i++ {
        j := i-1
        
        take := freq[keys[i]]
        for j >= 0 && keys[i] - keys[j] <= 2 {
            j--
        }

        if j >= 0 {
            take += dp[j]
        }

        dp[i] = max(dp[i-1], take)
    }
    return dp[len(keys) - 1]
}