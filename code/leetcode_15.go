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
