## 二分查找

[TOC]

#### [34. 在排序数组中查找元素的第一个和最后一个位置(Medium)](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/)

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

```
输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
```

示例 2:

```
输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
```

题解：

使用左闭右开的方式

```go
func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    l := lowBound(nums, target)
    if l >= len(nums) || nums[l] != target {
        return []int{-1, -1}
    }
    return []int{l, upperBound(nums, target)-1}
}

func lowBound(arr []int, x int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m] >= x {
			r = m
		} else {
			l = m+1
		}
	}
	return l
}

func upperBound(arr []int, x int) int {
	l, r := 0, len(arr)
	for l < r {
		m := l + (r-l)/2
		if arr[m] > x {
			r = m
		} else {
			l = m+1
		}
	}
	return l
}
```



#### [153. 寻找旋转排序数组中的最小值(Medium)](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/)

假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。

请找出其中最小的元素。

示例 1：

```
输入：nums = [3,4,5,1,2]
输出：1
```

示例 2：

```
输入：nums = [4,5,6,7,0,1,2]
输出：0
```

示例 3：

```
输入：nums = [1]
输出：1
```

题解：

对二分查找的一种变形，判断arr[m]是否小于nums[r]。如果成立那么说明最小值在nums[l-m]这个区间。反之在nums[m+1:r]这个区间



#### [154. 寻找旋转排序数组中的最小值 II(Hard)](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/)


假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 `[0,1,2,4,5,6,7]` 可能变为 `[4,5,6,7,0,1,2]` )。

请找出其中最小的元素。

注意数组中可能存在重复的元素。

**示例 1：**

```
输入: [1,3,5]
输出: 1
```

**示例 2：**

```
输入: [2,2,2,0,1]
输出: 0
```

**题解：**

对于有重复的元素，如出现nums[m]==nums[r]的时候，只需要将r-1即可。

```go
func findMin(nums []int) int {
    l, r := 0, len(nums)-1

    for l < r {
        m := l + (r-l)/2
        if nums[m] < nums[r] {
            r = m
        } else if nums[m] == nums[r] {
            r--
        } else {
            l = m+1
        }
    }
    return nums[l]
}
```

