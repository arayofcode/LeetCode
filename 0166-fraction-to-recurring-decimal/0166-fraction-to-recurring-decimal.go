func fractionToDecimal(numerator int, denominator int) string {
    if numerator == 0 {
        return "0"
    }

    result := ""
    if (numerator < 0) != (denominator < 0) {
        result += "-"
        if numerator < 0 {
            numerator *= -1
        }
        if denominator < 0 {
            denominator *= -1
        }
    }
    result += fmt.Sprint(numerator / denominator)
    remainder := numerator % denominator
    if remainder == 0 {
        return result
    }

    result += "."
    // remainder and its index
    remainderMap := make(map[int]int)
    
    pattern := ""
    for remainder != 0 {
        if index, found := remainderMap[remainder]; found {
            result += pattern[:index]
            result += "(" + pattern[index:] + ")"
            return result
        }

        remainderMap[remainder] = len(pattern)
        pattern += fmt.Sprint((remainder * 10) / denominator)
        remainder = (remainder * 10) % denominator

        if len(pattern) > 10000 {
            break
        }
    }
    result += pattern
    return result
}

