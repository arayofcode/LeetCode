func gcdOfStrings(str1 string, str2 string) string {
    // Find factors of the smaller string
    // Check if they are factors of larger string
    // Return the largest such factor
    smallerString := ""
    largerString := ""
    if len(str1) > len(str2){
        smallerString = str2
        largerString = str1
    } else {
        smallerString = str1
        largerString = str2
    }
    factors := factorsOfString(smallerString)
    largestFactor := ""
    for _, factor := range factors {
        if isFactor(largerString, factor) {
            largestFactor = factor
        }
    }
    return largestFactor
}

func factorsOfString(str string) (factors []string) {
    factor := ""
    for _, character := range str {
        factor += string(character)
        if isFactor(str, factor) {
            factors = append(factors, factor)
        }
    }
    return
}

func isFactor(str string, factor string) bool {
    i := len(str) / len(factor)
    if len(factor) * i != len(str) {
        return false
    }
    merged := ""
    for i != 0 {
        merged += factor
        i--
    }
    if merged == str {
        return true
    }
    return false
}