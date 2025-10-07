func avoidFlood(rains []int) []int {
    results := make([]int, len(rains))
    for i := range results {
        results[i] = 1
    }

    sunny := []int{}
    fullLakes := make(map[int]int)

    for i, rain := range rains {
        if rain == 0 {
            sunny = append(sunny, i)
        } else {
            results[i] = -1
            if day, found := fullLakes[rain]; found {
                j := sort.SearchInts(sunny, day)
                if j == len(sunny) {
                    return []int{}
                }
                results[sunny[j]] = rain
                copy(sunny[j:], sunny[j+1:])
                sunny = sunny[:len(sunny) - 1]
            }
            fullLakes[rain] = i
        }
    }
    return results
}