func twoSum(numbers []int, target int) []int {
    for i, j := 0, len(numbers)-1; ; {
        if numbers[i] + numbers[j] == target {
            return []int{i+1, j+1}
        } else if numbers[i] + numbers[j] > target {
            j--
        } else {
            i++
        }
    }
    return nil
}
