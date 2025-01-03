func longestCommonPrefix(strs []string) string {
	prefix := ""
	for i := 0; i < smallestStringSize(strs); i++ {
		curr := strs[0][i]
		for _, currStr := range strs {
			if curr != currStr[i] {
				return prefix
			}
		}
		prefix += string(curr)
	}
	return prefix
}

func smallestStringSize(strs []string) int {
	minimum := len(strs[0])
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < minimum {
			minimum = len(strs[i])
		}
	}
	return minimum
}