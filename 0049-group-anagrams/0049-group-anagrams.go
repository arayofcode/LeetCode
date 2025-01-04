// Approach 1 - Sorting: For all strings, sort characters
// Bucket same ones together
// Time: O(m log m) for sorting each string + O(n)
// for creating final results = O(m log m + n)
import (
    "sort"
    "strings"
)

func groupAnagrams(strs []string) (results [][]string) {
	anagrams := make(map[string][]string)
	for _, str := range strs {
		sorted := sortString(str)
		anagrams[sorted] = append(anagrams[sorted], str)
	}

	for _, group := range anagrams {
		results = append(results, group)
	}
	return results
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}