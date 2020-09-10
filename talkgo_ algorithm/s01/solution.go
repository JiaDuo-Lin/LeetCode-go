package s01

import "math"

// 题目1：
// 输入一个递增排序的数组和一个数字S，在数组中查找两个数，是的他们的和正好是S，
// 如果有多对数字的和等于S，输出两个数的乘积最小的。

// 题解1 由于是排好序的，直接使用双指针即可
func twoSum1(nums []int, s int) []int {
	max := math.MaxInt64
	var answer []int
	for head, tail := 0, len(nums)-1; head<tail; {
		if nums[head]+nums[tail] == s {
			if nums[head]*nums[tail] < max {
				max = nums[head]*nums[tail]
				answer = []int{nums[head], nums[tail]}
			}
			tail--
			head++
		} else if nums[head]+nums[tail] > s {
			tail--
		} else {
			head++
		}
	}
	return answer
}


// 题解2 根据测试题解1可以猜想找到的第一个解就是最优，
// 当然也可以去证明，证明过程直接设，然后求导即可
func twoSum2(nums []int, s int) (answer []int) {
	for head, tail := 0, len(nums)-1; head<tail; {
		if nums[head]+nums[tail] == s {
			return []int{nums[head], nums[tail]}
		} else if nums[head]+nums[tail] > s {
			tail--
		} else {
			head++
		}
	}
	return
}


// 题目二
// 一个有序数组，从随即一位截断，把前段放在后边，如 4 5 6 7 1 2 3求中位数

// 思路找到数组中最小数的位置，然后根据最小数的位置去寻找中位数
// 先判断数组是否被旋转，如果没有旋转，那么最小数就在下标为0处
// 被旋转，那么就使用二叉查找的方法去寻找最小数下标
func findMedian(nums []int) float32 {
	n := len(nums)
	var minIndex int
	if nums[0] < nums[n-1] {
		minIndex = 0
	} else {
		minIndex = biSearch(nums)
	}
	if n % 2 == 1 {
		return float32(nums[(minIndex+n/2)%n])
	}
	x1 := float32(nums[(minIndex+n/2-1)%n])
	x2 := float32(nums[(minIndex+n/2)%n])
	return (x1 + x2) / 2
}

func biSearch(nums []int) (mid int) {
	low, mid, high := 0, len(nums)/2, len(nums)-1
	for nums[mid-1] < nums[mid] {
		if nums[mid] > nums[low] {
			low = mid
		} else {
			high = mid
		}
		mid = (high + low) / 2
	}
	return
}