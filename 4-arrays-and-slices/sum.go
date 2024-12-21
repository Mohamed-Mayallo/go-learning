package main

func Sum(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	return sum
}

func SumAll(slices ...[]int) []int {
	lenOfSlices := len(slices)
	sumAll := make([]int, lenOfSlices)

	for i, s := range slices {
		sum := Sum(s)

		sumAll[i] = sum
	}

	return sumAll
}

func SumAllTails(slices ...[]int) []int {
	sumAll := []int{}

	for _, s := range slices {
		tail := s[1:]
		sumAll = append(sumAll, Sum(tail))
	}

	return sumAll
}
