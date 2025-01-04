// Binary Search Solution: O(n log k), k = len(shortest string)
// Find smallest word size
// From size / 2, check if all substrings of this size have common prefix
// If yes, check between size / 2 + 1, high
// Else check low to size / 2 - 1
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
        return strs[0]
    }

    minSize := smallestStringSize(strs)
	low := 0
	high := minSize
	for low <= high {
		mid := (low + high) / 2
        if hasCommonPrefix(strs, mid) {
            low = mid + 1
        } else {
            high = mid - 1
        }
	}
    return strs[0][:(low + high)/2]
}

func hasCommonPrefix(strs []string, mid int) bool {
    prefix := strs[0][:mid]
    for _, currStr := range strs {
        if prefix != currStr[:mid] {
            return false
        }
    }
    return true
}

// Brute Force: O(m*n)
// For each index in a string i, check if string[i]
// is same for all strings in array
func longestCommonPrefixBruteForce(strs []string) string {
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