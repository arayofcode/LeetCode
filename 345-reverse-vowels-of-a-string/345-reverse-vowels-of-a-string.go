func reverseVowels(s string) string {
    vowels := "aeiouAEIOU"
    reversed := ""
    for _, c := range s {
        character := string(c)
        if strings.Contains(vowels, character) {
            reversed = character + reversed
        }
    }

    for i, j := 0, 0; i < len(s); i++ {
        if strings.Contains(vowels, string(s[i])) {
            s = s[:i] + string(reversed[j]) + s[i+1:]
            j++
        }
    }
    return s
}