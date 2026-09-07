func lengthOfLongestSubstring(s string) int {
    longestLen := 0
    curLen := 0
    windowStart := 0
    curChars := map[rune]bool{}
    for _, char := range s {
        if _, seenChar := curChars[char]; seenChar {
            for ; rune(s[windowStart]) != char; windowStart++ {
                curLen--
                delete(curChars, rune(s[windowStart]))
            }
            windowStart++
        } else {
            curChars[char] = true
            curLen++
            if curLen > longestLen {
                longestLen = curLen
            }
        }
    }
    return longestLen
}
