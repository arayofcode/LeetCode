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