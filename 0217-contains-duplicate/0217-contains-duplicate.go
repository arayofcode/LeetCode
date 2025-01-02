func containsDuplicate(nums []int) bool {
    var set = make(map[int]struct{}, len(nums))
    for _, num := range nums {
        if _, found := set[num]; found {
            return true
        }
        set[num] = struct{}{}
    }
    return false
}