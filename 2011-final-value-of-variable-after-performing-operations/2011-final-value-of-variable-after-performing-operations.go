func finalValueAfterOperations(operations []string) int {
    res := 0
    for _, val := range operations {
        if val[1] == '-' {
            res--
        } else {
            res++
        }
    }
    return res
}