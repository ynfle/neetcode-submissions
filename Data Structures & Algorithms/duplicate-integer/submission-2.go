func hasDuplicate(nums []int) bool {
    set := map[int]struct{}{}
    for _, num := range nums {
        if _, hasVal := set[num]; hasVal {
            return true
        }
        set[num] = struct{}{}
    }
    return false
}
