# 链接
https://leetcode-cn.com/problems/string-to-integer-atoi/

# 题目
```
请你来实现一个 myAtoi(string s) 函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数 myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231,  231 − 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

```
# 思路
可以把解题拆接为几个步骤，依次攻克
- 将原始字符串进行处理，得到只包含数字的字符串
- 将只包含数字的字符串去掉前面多余的0
- 将去掉多余的0的字符串转换成整数，注意边界条件处理
    - 边界条件有两个，一个是底数不断*10时是否达到边界
    - 一个是每一次累计到得到当前位为止的整数时，是否达到边界

```go
func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }
    negative := false
    nums := []byte{}
    for i := 0; i < len(s); i++ {
        if i == 0 {
            if s[i] == '-' {
                negative = true
                continue
            }
            if s[i] == '+' {
                continue
            }
        }
        if s[i] >= '0' && s[i] <= '9' {
            nums = append(nums, s[i])
        } else {
            break
        }
    }
    return strToInt(negative ,nums)
}

func strToInt(negative bool, nums []byte) int {
    idx := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == '0' {
            idx++
        } else {
            break
        }
    }
    nums = nums[idx:]
    num := 0
    base := 1
    for i := len(nums) - 1; i >= 0; i-- {
        if notInArea(base, negative) {
            return limit(negative)
        } 
        num += base * charToInt(nums[i])
        if notInArea(num, negative) {
            return limit(negative)
        }
        base *= 10
    }
    if negative {
        return -num
    }
    return num
}

func notInArea(num int, negative bool) bool {
    return (negative && num >= 1 << 31) || (!negative && num >= 1 << 31 - 1)
}

func limit(negative bool) int {
    if negative {
        return -(1 << 31)
    }
    return 1 << 31 - 1
}

func charToInt(c byte) int {
    if c == '0' {
        return 0
    }
    if c == '1' {
        return 1
    }
    if c == '2' {
        return 2
    }
    if c == '3' {
        return 3
    }
    if c == '4' {
        return 4
    }
    if c == '5' {
        return 5
    }
    if c == '6' {
        return 6
    }
    if c == '7' {
        return 7
    }
    if c == '8' {
        return 8
    }
    if c == '9' {
        return 9
    }
    return -1
}
```
