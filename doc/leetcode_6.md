# 链接
https://leetcode-cn.com/problems/zigzag-conversion/
# 题目
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
```
P   A   H   N
A P L S I I G
Y   I   R
```
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

请你实现这个将字符串进行指定行数变换的函数：
```
string convert(string s, int numRows);
```
# 思路
- 可以看出，题目要求多少行，最终结果即为每一行的结果拼接在一起即可。因此我们需要做的就是不断填充每一行的字符串
- 如何填充呢？首先根据行数n，创建n个数组，然后不断移动字符串的索引，最开始不断递增填充每个字符串，直到填充到最后一行，然后翻转开始递减每一行，直到填充了第一行，再次翻转。
- 最终，将各行的字符串拼接在一起，即为最终的答案
- 注意当行数为1时，由于它既是第一行也是最后一行，因此需要特殊处理

```go
func convert(s string, numRows int) string {
    if numRows == 1 {
        return s
    }
    matrix := make([][]byte, numRows)
    for i := 0; i < numRows; i++ {
        matrix[i] = []byte{}
    }
    flag := -1
    i := 0
    for j := 0; j < len(s); j++ {
        matrix[i] = append(matrix[i], s[j])
        if i == numRows - 1 || i == 0 {
            flag = -flag
        }
        i += flag
    }
    result := []byte{}
    for i := 0; i < numRows; i++ {
        result = append(result, matrix[i]...)
    }
    return string(result)
}
```
