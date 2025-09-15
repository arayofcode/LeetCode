func canBeTypedWords(text string, brokenLetters string) int {
    brokenLettersMap := make(map[rune]struct{})
    for _, chr := range brokenLetters {
        brokenLettersMap[chr] = struct{}{}
    }
    words := strings.Fields(text)
    count := 0
    for _, word := range words {
        brokenLetterFound := false
        for _, chr := range word {
            if _, found := brokenLettersMap[chr]; found {
                brokenLetterFound = true
                break
            }
        }
        if !brokenLetterFound {
            count += 1
        }
    }
    return count
}