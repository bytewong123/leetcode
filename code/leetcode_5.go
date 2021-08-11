func longestPalindrome(s string) string {
    dp := make([][]bool, len(s))
    for i := 0; i < len(s); i++ {
        dp[i] = make([]bool, len(s))
        dp[i][i] = true
    }
    left, right := 0, 0
    maxLen := 1
    for step := 0; step < len(s); step++ {
        for i := 0; i + step < len(s); i++ {
            j := i + step
            if i == j {
                continue
            }
            if step == 1 && s[i] == s[j] {
                dp[i][j] = true
                left, right = i, j
                maxLen = 2
                continue
            }
            if dp[i + 1][j - 1] && s[i] == s[j] {
                dp[i][j] = true
                if j - i + 1 > maxLen {
                    left, right = i, j
                    maxLen = j - i + 1
                }
            }
        }
    }
    return s[left:right+1]
}
