func maxDistinctElements(nums []int, k int) int {
    slices.Sort(nums)
    last := -99999999
    count := 0
    for _, num := range nums {
        p := max(num - k, last + 1)
        if p <= num + k {
            last = p
            count++
        }
    }
    return count
}