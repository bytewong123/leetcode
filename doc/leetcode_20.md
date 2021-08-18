# 链接
https://leetcode-cn.com/problems/valid-parentheses/
# 题目
```
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
 

示例 1：

输入：s = "()"
输出：true
示例 2：

输入：s = "()[]{}"
输出：true
示例 3：

输入：s = "(]"
输出：false
示例 4：

输入：s = "([)]"
输出：false
示例 5：

输入：s = "{[]}"
输出：true
```
# 思路
- 栈的常规思路，若新括号匹配到栈顶的括号，则栈顶括号弹出，新括号丢弃；未匹配到则新括号入栈
```go
func isValid(s string) bool {
    stack := make([]byte, 0, len(s))
    for i := 0; i < len(s); i++ {
        if len(stack) == 0 {
            stack = append(stack, s[i])
            continue
        } 
        if (s[i] == ')' && stack[len(stack) - 1] == '(') ||
        (s[i] == ']' && stack[len(stack) - 1] == '[') ||
        (s[i] == '}' && stack[len(stack) - 1] == '{') {
            stack = stack[:len(stack) - 1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}
```
