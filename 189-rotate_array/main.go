package main

import "log"

func main() {
	nums := []int{-1, -100, 3, 99}
	rotate(nums, 2)

	log.Printf("%+v", nums)

	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)

	log.Printf("%+v", nums)
}

func highMemoryRotate(nums []int, k int) {
	k = k % len(nums)
	final := append(nums[len(nums)-k:], nums[:len(nums)-k]...)

	for i := range nums {
		nums[i] = final[i]
	}
}

func rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := nums[len(nums)-1]
		for j := len(nums) - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}

		nums[0] = tmp
	}
}
