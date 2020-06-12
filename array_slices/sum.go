package main

func Sum(numbers []int) (sum int) {
	sum = 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(slices [][]int) []int {
	sum := make([]int, len(slices))
	for i, arr := range slices {
		sum[i] = Sum(arr)
	}
	return sum
}

func SumAllTails(ints [][]int) []int {
	sum := make([]int, len(ints))
	for i, arr := range ints {
		if len(arr) > 1 {
			sum[i] = Sum(arr[1:])
		} else {
			sum[i] = 0
		}
	}
	return sum
}

func SumTails(ints []int) (sum int) {
	for i := 1; i < len(ints); i++ {
		sum += ints[i]
	}
	return sum
}
