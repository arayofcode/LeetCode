func removeElement(nums []int, val int) int {
    return twoPointers(nums, val)
}

// Approach 1 - Two Pointers
// One pointer to iterate, the other to swap
// Start from 0. If non-val, swap, and increase count
// Time: O(n) Space: O(1)
func twoPointers(nums []int, val int) int {
    k := 0
    for _, num := range nums {
        if num != val {
            nums[k] = num
            k++
        }
    }
    return k
}

// Approach 2 - Hash Table:
// Create a frequency table for all elements in nums
// From table, if key is not val, insert it table[key]
// times in nums. Return count of such elements
// Time: O(n), Space: O(n)
func hashTable(nums []int, val int) (count int) {
    table := make(map[int]int)
    for _, num := range nums {
        table[num]++
    }

    for element, frequency := range table {
        if element != val {
            for i := 0; i < frequency; i++ {
                nums[count] = element
                count++
            }
        }
    }
    return count
}

// Approach 3 - Slice off any val found
// Create new slice, append non-val elements
// Copy this slice back to nums
// Time: O(n) Space: O(n)
func slicing(nums []int, val int) int {
    var result []int
    for _, num := range nums {
        if num != val {
            result = append(result, num)
        }
    }
    copy(nums, result)
    return len(result)
}