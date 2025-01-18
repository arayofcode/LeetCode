func strStr(haystack string, needle string) int {
    needleLen := len(needle)
    haystackLen := len(haystack)
    
    for i := 0; i < haystackLen - needleLen + 1; i++ {
        substring := haystack[i: i+needleLen]
        if substring == needle {
            return i
        }
    }
    return -1
}