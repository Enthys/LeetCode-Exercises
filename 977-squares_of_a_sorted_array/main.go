package main

import (
	"log"
	"sort"
)

func main() {
	log.Printf("%+v", squareAndSort([]int{-4, -1, 0, 3, 10}))
}

func simpleSquareAndSort(nums []int) []int {
	for i, val := range nums {
		nums[i] = val * val
	}

	sort.Ints(nums)

	return nums
}

func squareAndSort(nums []int) []int {
	results := make([]int, len(nums))
	left, right := 0, len(nums)-1

	for i := len(nums) - 1; left <= right && left < len(nums) && right >= 0; i-- {
		sqLeft := nums[left] * nums[left]
		sqRight := nums[right] * nums[right]

		log.Printf("i: %d, left: %d, right: %d", i, left, right)
		if sqLeft > sqRight {
			results[i] = sqLeft
			left++
		} else {
			results[i] = sqRight
			right--
		}
	}

	return results
}
