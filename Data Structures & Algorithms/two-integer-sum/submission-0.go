func twoSum(nums []int, target int) []int {
    set := map[int]int{}
    for idx, num := range nums {
        if i, hasVal := set[target-num]; hasVal {
                return []int{i, idx}
        }
        set[num] = idx
    }
    return nil
}
