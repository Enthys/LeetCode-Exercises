package pivot_index

func PivotIndex(numbers []int) int {
	for index := range numbers {
		left := sumNumbers(numbers[0:index])
		right := sumNumbers(numbers[index+1:])

		if left == right {
			return index
		}
	}

	return -1
}

func sumNumbers(numbers []int) int {
	result := 0
	for _, num := range numbers {
		result += num
	}

	return result
}
