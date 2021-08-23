# 链接
https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/
# 题目
```
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
```
# 思路
- 使用双指针法，左指针指向当前位置，右指针去探索不重复的元素
- 若左指针与右指针的值相同，右指针开始右移，直到找到一个和左指针不同的值
- 然后左指针的左边一个位置替换为右指针找到的第一个与左指针不同的值即可
# 时空复杂度
- 时间复杂度
O(n)
- 空间复杂度
O(1)

```go
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    left, right := 0, 0
    for right < len(nums) {
        if nums[left] == nums[right] {
            right++
            continue
        } 
        left++
        nums[left] = nums[right]
        right++
    }
    return left + 1
}
```
