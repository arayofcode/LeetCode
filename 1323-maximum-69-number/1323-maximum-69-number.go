func maximum69Number (num int) (result int) {
    numStr := fmt.Sprintf("%d", num)
    flag := 1
    for i := range numStr {
        chr := numStr[i]
        if chr == '6' && flag != 0 {
            chr = '9'
            flag = 0
        }
        result = result * 10 + int(chr - '0')
    }
    return result
}