# 链接
https://leetcode-cn.com/problems/reverse-integer/
# 题目
```
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。

假设环境不允许存储 64 位整数（有符号或无符号）。
 

示例 1：

输入：x = 123
输出：321
示例 2：

输入：x = -123
输出：-321
示例 3：

输入：x = 120
输出：21
示例 4：

输入：x = 0
输出：0
```
# 思路
- 不断对x以10取模，得到最后一位
- 最后一位不断*10 + 下一次的模
- x每次除以10，以便向左移动，直到<=0
```go
func reverse(x int) int {
    if x == 0 {
        return 0
    }
    negative := false
    if x < 0 {
        negative = true
        x = -x
    }
    num := 0
    sum := 0
    for x > 0 {
        num = x % 10
        sum = sum * 10 + num
        x /= 10
    }
    if negative {
        if sum > 1 << 31 {
            return 0
        }
        return -sum
    }
    if sum > 1 << 31 - 1 {
        return 0
    }
    return sum
}
```
