func maxFreqSum(s string) int {
    vowels := map[rune]int{
        'a': 0,
        'e': 0,
        'i': 0,
        'o': 0,
        'u': 0,
    }
    maxVowelFrequency := 0
    consonants := make(map[rune]int)
    maxConsonantFrequency := 0
    for _, chr := range s {
        if frequency, found := vowels[chr]; found {
            frequency++
            vowels[chr] = frequency
            if frequency > maxVowelFrequency {
                maxVowelFrequency = frequency
            }
        } else {
            consonants[chr]++
            if consonants[chr] > maxConsonantFrequency {
                maxConsonantFrequency = consonants[chr]
            }
        }
    }
    return maxVowelFrequency + maxConsonantFrequency
}