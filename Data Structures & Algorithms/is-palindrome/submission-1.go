func isPalindrome(s string) bool {
    return do(s, 0, len(s)-1)
}

func do(s string, i int, j int) bool {
    if j - i <= 0 {
        return true
    }

    if !isAlphaNumeric(s[i]) {
        return do(s, i+1, j)
    } 

    if !isAlphaNumeric(s[j]) {
        return do(s, i, j-1)
    }

    if toLower(s[i]) == toLower(s[j]) {
        return do(s, i+1, j-1)
    }

    return false
}

func isAlphaNumeric(b byte) bool {
    return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9'
}

func toLower(b byte) byte {
    if b >= 'A' && b <= 'Z' {
        return b - 'A' + 'a'
    }
    return b
}