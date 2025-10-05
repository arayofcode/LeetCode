func removeDigit(number string, digit byte) string {
    maxNum := ""
    for i := 0; i < len(number); i++ {
        if byte(number[i]) == digit {
            newNum := number[:i] + number[i+1:]
            if newNum > maxNum {
                maxNum = newNum
            }
        }
    }
    return maxNum
}