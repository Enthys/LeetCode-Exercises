package main

import "log"

func main() {
	log.Printf("%d", searchInsert([]int{1, 3, 5, 6}, 5))
}

func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)

	for low < high {
		mid := low + (high-low)/2

		if nums[mid] >= target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
