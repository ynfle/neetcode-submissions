import "slices"

func groupAnagrams(strs []string) [][]string {
    set := map[string][]string{}
    for _, i := range strs {
        chrs := []byte(i)
        slices.Sort(chrs)
        if slice, hasSlice := set[string(chrs)]; hasSlice {
            set[string(chrs)] = append(slice, i)
        } else {
            set[string(chrs)] = []string{i}
        }
    }
    
    result := [][]string{}
    for _, val := range set {
        result = append(result, val)
    }
    return result
}
