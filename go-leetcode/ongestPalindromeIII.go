func longestPalindrome(s string) string {
    if len(s) == 0 {
        return ""
    }

    start, end := 0, 0

    for i := 0; i < len(s); i++ {
        // Odd-length palindromes (single center)
        l1 := expand(s, i, i)
        // Even-length palindromes (double center)
        l2 := expand(s, i, i+1)

        longest := max(l1, l2)

        if longest > end-start+1 {
            start = i - (longest-1)/2
            end   = i + longest/2
        }
    }

    return s[start : end+1]
}

func expand(s string, left, right int) int {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--
        right++
    }
    return right - left - 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}