func mergeAlternately(word1 string, word2 string) (merged string) {
    i := 0
    for i < len(word1) && i < len(word2) {
        if i < len(word1) {
            merged += string(word1[i])
        }
        if i < len(word2) {
            merged += string(word2[i])
        }
        i++
    }
    if i < len(word1) {
        merged += word1[i:]
    }
    if i < len(word2) {
        merged += word2[i:]
    }
    return merged
}