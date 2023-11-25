func gcdOfStrings(str1 string, str2 string) string {
    // Two strings with same factor essentially sum up to same merged string
    if str1 + str2 != str2 + str1 {
        return ""
    }
    // GCD of their length is the solution
    return str1[:gcd(len(str1), len(str2))]
}

func gcd(a int, b int) int {
    for b != 0 {
        a, b = b, a % b
    }
    return a
}

func gcdRecursive(a int, b int) int {
    if b == 0 {
        return a
    } else {
        return gcd(b, a % b)
    }
}

func gcdOfStrings_factorsApproach(str1 string, str2 string) string {
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