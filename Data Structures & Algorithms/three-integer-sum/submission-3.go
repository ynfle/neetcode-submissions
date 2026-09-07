import "slices"

func threeSum(nums []int) [][]int {
    n := nums[:]
    slices.Sort(n)
    results := map[string][]int{}
    for idx, num := range n {
        if idx == len(nums)-2 {
            break
        }
        if twoSumResult := twoSumSorted(n[idx+1:], -num); twoSumResult != nil {
            for _, result := range twoSumResult {
                r := append(result, num)
                slices.Sort(r)
                key := fmt.Sprintf("%d,%d,%d", r[0], r[1], r[2])
                results[key] = r
            }
        }
    }
    result := make([][]int, 0, len(results))
    for _, r := range results {
        result = append(result, r)
    }
    return result
}

func twoSumSorted(numbers []int, target int) [][]int {
    result := [][]int{}
    for i, j := 0, len(numbers)-1; i < j; {
        if numbers[i] + numbers[j] == target {
            result = append(result, []int{numbers[i], numbers[j]})
            i++
        } else if numbers[i] + numbers[j] > target {
            j--
        } else {
            i++
        }
    }
    return result
}
