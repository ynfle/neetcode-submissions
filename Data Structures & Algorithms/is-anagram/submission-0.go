func isAnagram(s string, t string) bool {
    sChars := [26]int{}
    tChars := [26]int{}
    for _, char := range s {
        sChars[char-'a'] += 1
    }
    fmt.Println(sChars)
    for _, char := range t {
        tChars[char-'a'] += 1
    }
    fmt.Println(tChars)
    for idx, tChar := range tChars {
        if sChars[idx] != tChar {
            return false
        }
    }
    return true
}
