func groupAnagrams(strs []string) [][]string {
    set := map[string][]string{}
    for _, i := range strs {
        key := strToKey(i)
        if slice, hasSlice := set[key]; hasSlice {
            set[key] = append(slice, i)
        } else {
            set[key] = []string{i}
        }
    }
    
    result := [][]string{}
    for _, val := range set {
        result = append(result, val)
    }
    return result
}

func strToKey(str string) string {
    key := make([]byte, 26)
    for _, chr := range str {
        key[chr-'a'] += 1
    }
    return string(key)
}
