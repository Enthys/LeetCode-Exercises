package main

import (
	"log"
	"time"
)

func main() {
	bigArr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		bigArr[i] = i
	}

	searchFor := 500000

	start := time.Now()
	found := binarySearch(bigArr, searchFor)
	elapsed := time.Since(start)
	log.Printf("Result: %d | Time Spent: %s", found, elapsed)

	start = time.Now()
	for i, val := range bigArr {
		if val == searchFor {
			found = i
		}
	}
	elapsed = time.Since(start)
	log.Printf("Result: %d | Time Spent: %s", found, elapsed)
}

func binarySearch(nums []int, target int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
