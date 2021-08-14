# 链接
https://leetcode-cn.com/problems/palindrome-number/
# 题目
```
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。

```
# 思路
- 只需要将一个整数从右往左算一个新的整数，然后比较前后两个整数即可
- 从右往左求整数的方式，每次获取到右侧的一位，然后不断新数*10 + 右侧的第一位作为新数，不断循环直到老数变为0

```go
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    origin := x
    newX := 0
    base := 10
    for x > 0 {
        extra := x % 10
        newX = extra + base * newX
        x /= 10
    }
    return newX == origin
}
```
