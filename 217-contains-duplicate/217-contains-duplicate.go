func containsDuplicate(nums []int) bool {
    set := make(map[int]bool, len(nums))
    for _, num := range nums {
        _, err := set[num]
        if err {
            return true
        }
        set[num] = true
    }
    return false
}