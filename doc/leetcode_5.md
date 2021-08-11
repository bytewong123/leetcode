链接：https://leetcode-cn.com/problems/longest-palindromic-substring/

给你一个字符串 s，找到 s 中最长的回文子串。

思路：
1. 我们需要用一个二维数组来保存当前回文字符串的状态。dp[i][j]代表以i开头，j结尾的字符串是否为回文字符串。
2. 然后从头开始，每次圈定一个固定长度的字符串，判断是否为回文子串，然后不断向右移，直到走到头，在这个过程中，我们会不断填充这个dp数组
3. 不断扩大圈定范围，判断圈定范围的字符串是否为回文子串，然后不断右移，在这个过程中更新最大长度和左右索引
4. 最后返回左右索引所圈定的字符串即可

```go
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
