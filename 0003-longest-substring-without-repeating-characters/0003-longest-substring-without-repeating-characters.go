func lengthOfLongestSubstring(s string) int {
    return approachSlidingWindow(s)
}

// Approach 1 - Brute force:
// Calculate all possible substrings
// Find all substrings with unique characters
// Return one with longest length
// Uses DP
// Time Complexity: O(n^2), Space Complexity: O(n)
func approachBruteForce(s string) int {
    maxLen := 0
    for start := 0; start < len(s); start++ {
        seen := make(map[rune]bool)
        for end := start; end < len(s); end++ {
            if seen[rune(s[end])] {
                break
            }
            seen[rune(s[end])] = true
            maxLen = max(maxLen, end-start+1)
        }
    }
    return maxLen
}

// Approach 2 - Sliding Window
// Check a sliding window, and store each element's index
// If duplicate found, check largest length found
// Time Complexity: O(n^2), Space Complexity: O(n)
func approachSlidingWindow(s string) int {
    start, maxLen := 0, 0
    characters := make(map[rune]int)
    for i, character := range s {
        if index, found := characters[character]; found && index >= start {
            start = index + 1
        }
        characters[character] = i
        maxLen = max(maxLen, i - start + 1)
    }
    return maxLen
}