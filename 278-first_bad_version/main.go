package main

import (
	"log"
)

func main() {

	log.Printf("%d", firstBadVersion(1000000))

}

func firstBadVersion(n int) int {
	low, high := 1, n

	for low < high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func isBadVersion(version int) bool {
	return version >= 5000
}
