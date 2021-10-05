package _4arary

func Sum(arrayOfNumbers []int) int {
	var sum int
	for _, number := range arrayOfNumbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumOfTails(arrayOfNumbers ...[]int) []int {
	var sum []int
	for _, number := range arrayOfNumbers {
		if len(number) == 0 {
			sum = append(sum, 0)
		} else {
			tail := number[1:]
			sum = append(sum, Sum(tail))
		}
	}
	return sum
}
