type Solution struct{}

func (*Solution) Encode(strs []string) string {
    result := ""
    for _, str := range strs {
        result += fmt.Sprintf("%d#%s", len(str), str)
    }
    return result
}

func (*Solution) Decode(s string) []string {
    result := []string{}
    start := 0
    idx := 0
    for idx < len(s) {
        chr := s[idx]
        if chr == '#' {
            strSize, err := strconv.Atoi(s[start:idx])
            if err != nil {
                panic(err)
            }
            start = strSize + idx + 1
            result = append(result, s[idx+1:start])
            idx = idx + strSize
        }
        idx++
    }
    return result
}
