package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	// make可以在创建切片的时候指定我们需要的长度和容量。
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
