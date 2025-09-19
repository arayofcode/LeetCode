func maxFreqSum(s string) int {
    letters := [26]int{}
    vowel := 0
    consonant := 0
    for _, chr := range s {
        i := int(chr - 'a')
        letters[i]++

        if i == 0 || i == 4 || i == 8 || i == 14 || i == 20 {
            if letters[i] > vowel {
                vowel = letters[i]
            }
        } else {
            if letters[i] > consonant {
                consonant = letters[i]
            }
        }
    }
    return vowel + consonant
}