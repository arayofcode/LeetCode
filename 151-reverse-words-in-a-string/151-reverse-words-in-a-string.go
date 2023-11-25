func reverseWords(s string) (result string) {
    var words []string
    word := ""
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if len(word) > 0 {
                words = append(words, strings.TrimSpace(word))
            }
            word = ""
        } else {
            word += string(s[i])
        }
    }
    if len(word) > 0 {
        words = append(words, word)
    }
    result = words[len(words) - 1]
    for i := len(words) - 2; i >= 0; i-- {
        result += " " + words[i]
    }
    return
}