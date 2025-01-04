import ( 
    "sort"
    "strings"
)

// Approach 3 - Improvise Hash Table by using arrays:
// Instead of creating a frequency map, store array of size 26
// for chr a to z. Each index stores frequency. Use this in map
func groupAnagrams(strs []string) (results [][]string) {
    frequencyMap := make(map[string][]string)
    for _, word := range strs {
        frequency := [26]int{}
        for _, chr := range word {
            frequency[chr - 97]++
        }
        frequencyStr := fmt.Sprintf("%v", frequency)
		frequencyMap[frequencyStr] = append(frequencyMap[frequencyStr], word)
    }
    for _, arr := range frequencyMap {
        results = append(results, arr)
    }
    return
}

// Approach 2 - Hash Table: for each word, store letter frequency
// Create another hash table where key is letter frequency hash table
// and value is a slice containing words
// append values of second hash table to another slice
func groupAnagramsHashTableApproach(strs []string) (results [][]string) {
    // Can't use map as key, so converting map to string
    frequencyMap := make(map[string][]string)
    for _, word := range strs {
        chrMap := make(map[rune]int)
        for _, chr := range word {
            chrMap[chr]++
        }
        chrMapString := mapToString(chrMap)
        frequencyMap[chrMapString] = append(frequencyMap[chrMapString], word)
    }
    for _, strings := range frequencyMap {
        results = append(results, strings)
    }
    return results
}

func mapToString(m map[rune]int) string {
    var keys []rune
    for k := range m {
        keys = append(keys, k)
    }
    sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
    
    var sb strings.Builder
    for _, k := range keys {
        sb.WriteString(fmt.Sprintf("%c:%d,", k, m[k]))
    }
    return sb.String()
}

// Approach 1 - Sorting: For all strings, sort characters
// Bucket same ones together
// Time: O(m log m) for sorting each string + O(n)
// for creating final results = O(m log m + n)
func groupAnagramsSortingApproach(strs []string) (results [][]string) {
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