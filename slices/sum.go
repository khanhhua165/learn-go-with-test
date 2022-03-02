package slices

func Sum(number []int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, len(numbers))
	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}
	return sums
}

func sumTails(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	return Sum(numbers[1:])
}

func SumAllTails(numbers ...[]int) []int {
	numLength := len(numbers)
	sums := make([]int, numLength)
	for i, numbers := range numbers {
		sums[i] = sumTails(numbers)
	}
	return sums
}
