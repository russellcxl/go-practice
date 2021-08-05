package main

import "fmt"

/*
A peak element is an element that is strictly greater than its neighbors.

Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

You may imagine that nums[-1] = nums[n] = -∞.

You must write an algorithm that runs in O(log n) time.
*/

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	for i := 1; i < len(nums); i++ {
		if i < len(nums)-1 {
			if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
				return i
			}
			if nums[i] > nums[i+1] {
				i++
			}
		} else {
			if nums[i] > nums[i - 1] {
				return i
			}
		}
	}

	return 0
}
