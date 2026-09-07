func longestConsecutive(nums []int) int {
    set := map[int]struct{}{}
    for _, num := range nums {
        set[num] = struct{}{}
    }

    var maxLength int
    for _, num := range nums {
        if _, has := set[num-1]; !has {
            var length int
            for length = 1; shrug(set, num, length) ; length++ {
            }
            if length > maxLength {
                maxLength = length
            }
        }
    }

    return maxLength
}

func shrug(set map[int]struct{}, num int, length int) bool {
    _, has := set[num+length]
    return has
}