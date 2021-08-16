# 链接
https://leetcode-cn.com/problems/3sum/submissions/
# 题目
```
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
示例 2：

输入：nums = []
输出：[]
示例 3：

输入：nums = [0]
输出：[]

```
# 思路
- 本题求三数之和，故使用双指针的方法，固定一个基准指针为数组的开头，然后另外两位开始不断靠拢移动，直到找到答案，然后基准指针开始右移
- 但是需要将数组排序，这样才能满足左小右大，方便通过判断当前和的大小来移动相应的指针
- 需要注意的是存在重复解的情况
    - 当基准指针与上一次基准指针相同，那么双指针的左指针若上一次是当前的基准指针的位置，且满足解，那么这个解会被保留；但是若左指针不是当前基准指针的位置，而是当前基准指针的右侧，那么重复解就出现了，若不止一个，那么当前基准指针的右侧会有多个重复解，因此需要排除这种情况，这种情况不会漏掉解
    - 若左指针与上一次左指针的位置相同，那么也需要过滤掉，不用担心丢解的情况，因为上一轮已经将解都保留下来了
```go
func threeSum(nums []int) [][]int {
    quickSort(nums, 0, len(nums) - 1)
    result := [][]int{}
    for i := 0; i <= len(nums) - 2; i++ {
        if i > 0 && nums[i] == nums[i - 1] {
            continue
        }
        left, right := i + 1, len(nums) - 1
        other := 0 - nums[i]
        for left < right {
            if nums[left] == nums[left - 1] && left - 1 != i {
                left++
                continue
            }
            if nums[left] + nums[right] == other {
                result = append(result, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            } else if nums[left] + nums[right] > other {
                right--
            } else if nums[left] + nums[right] < other {
                left++
            }
        }
    } 
    return result
}

func quickSort(nums []int, left, right int) {
    if left < right {
        mid := partition(nums, left, right)
        quickSort(nums, left, mid - 1)
        quickSort(nums, mid + 1, right)
    }
}

func partition(nums []int, left, right int) int {
    pivot := nums[left]
    for right > left {
        for right > left && nums[right] >= pivot {
            right--
        }
        nums[left] = nums[right]
        for right > left && nums[left] <= pivot {
            left++
        }
        nums[right] = nums[left]
    }
    nums[left] = pivot
    return left
}
```
