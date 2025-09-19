func maxFreqSum(s string) int {
    maxVowelFreq := 0
    vowelFreq := [5]int{}
    maxConsonantFreq := 0
    consonantFreq := make(map[byte]int, 21)
    for i := 0; i < len(s); i++ {
        chr := s[i]
        switch chr {
        case 'a':
            vowelFreq[0]++
            if vowelFreq[0] > maxVowelFreq {
                maxVowelFreq = vowelFreq[0]
            }
        case 'e':
            vowelFreq[1]++
            if vowelFreq[1] > maxVowelFreq {
                maxVowelFreq = vowelFreq[1]
            }
        case 'i':
            vowelFreq[2]++
            if vowelFreq[2] > maxVowelFreq {
                maxVowelFreq = vowelFreq[2]
            }
        case 'o':
            vowelFreq[3]++
            if vowelFreq[3] > maxVowelFreq {
                maxVowelFreq = vowelFreq[3]
            }
        case 'u':
            vowelFreq[4]++
            if vowelFreq[4] > maxVowelFreq {
                maxVowelFreq = vowelFreq[4]
            }
        default:
            consonantFreq[chr]++
            if consonantFreq[chr] > maxConsonantFreq {
                maxConsonantFreq = consonantFreq[chr]
            }
        }
    }
    return maxVowelFreq + maxConsonantFreq
}