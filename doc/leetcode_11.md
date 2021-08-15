# 链接
https://leetcode-cn.com/problems/container-with-most-water/
# 题目
```
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。
```
# 思路
- 通过设置左右两个指针
- 左右指针的举例为矩形的底边长
- 左右指针所指高度的较小值为矩形的高
- 此时需要移动指针，寻找下一个答案。将较小值的高度所代表的指针向内部移动。因为移动两个指针会导致底边长变短，这样才能找到高更大的值
> 这种做法是否会有面积遗漏？
如果选择固定一根柱子，另外一根变化，水的面积会有什么变化吗？稍加思考可得：

- 当前柱子是最两侧的柱子，水的宽度 为最大，其他的组合，水的宽度都比这个小。
- 左边柱子较短，决定了水的高度为 3。如果移动左边的柱子，新的水面高度不确定，一定不会超过右边的柱子高度 7。
- 如果移动右边的柱子，新的水面高度一定不会超过左边的柱子高度 3，也就是不会超过现在的水面高度。
- 由此可见，如果固定左边的柱子，移动右边的柱子，那么水的高度一定不会增加，且宽度一定减少，所以水的面积一定减少。这个时候，左边的柱子和任意一个其他柱子的组合，其实都可以排除了。也就是我们可以排除掉左边的柱子了

```go
func maxArea(height []int) int {
    if len(height) < 2 {
        return 0
    }
    left, right := 0, len(height) - 1
    maxVal := 0
    for left < right {
        area := (right - left) * min(height[left], height[right])
        if area > maxVal {
            maxVal = area
        }
        if height[left] > height[right] {
            right--
        } else {
            left++
        }
    }
    return maxVal
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```
