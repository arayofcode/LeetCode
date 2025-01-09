import "slices"

func majorityElement(nums []int) int {
    return BoyerMooreMajority(nums)
}

// Approach 1 - Counting frequency:
// Go through all elements and store their count
// Return one with largest count
// Time Complexity: O(n), Space Complexity: O(n)
func frequencyCountingApproach(nums []int) int {
    var frequencyTable = make(map[int]int)
    for _, elem := range nums {
        frequencyTable[elem]++
    }

    maxCount := 1
    maxNum := nums[0]
    for num, count := range frequencyTable {
        if count > maxCount {
            maxCount = count
            maxNum = num
        }
    }
    return maxNum
}

// Approach 2 - Sorting:
// Sort the array, and return median.
// Majority element is guaranteed to be the median.
// Time Complexity: O(n log n), Space Complexity: O(1)
func medianApproach(nums []int) int {
    slices.Sort(nums)
    return nums[len(nums)/2]
}

// Approach 3 - Boyer-Moore Majority Voting Algorithm
// Works on assumption that majority element exists
// Since non-majority are less than |n/2|, we increase
// or decrease an element's count until it's 0
// Once 0, we look for the next majority
// Time: O(n), Space: O(1)
func BoyerMooreMajority(nums []int) int {
    majority := nums[0]
    majorityCount := 1
    for i := 1; i < len(nums); i++ {
        if majorityCount == 0 {
            majority =  nums[i]
        }
        if nums[i] == majority {
            majorityCount++
        } else {
            majorityCount--
        }
    }
    return majority
}