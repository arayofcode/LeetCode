func maxFrequencyElements(nums []int) int {
    frequencies := make(map[int]int)
    maxFreq := 0
    for _, num := range nums {
        frequencies[num]++
        if frequencies[num] > maxFreq {
            maxFreq = frequencies[num]
        }
    }
    
    result := 0
    for _, val := range frequencies {
        if val == maxFreq {
            result += val
        }
    }
    return result
}