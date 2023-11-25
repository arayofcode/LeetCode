func reverseWords(s string) (result string) {
    word := ""
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if len(word) > 0 {
                result = word + " " + result
            }
            word = ""
        } else {
            word += string(s[i])
        }
    }
    if len(word) > 0 {
        result = word + " " + result
    }
    result = result[:len(result) - 1]
    return
}

func reverseWordsArray(s string) (result string) {
    // Split string into words array, join elements with space
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