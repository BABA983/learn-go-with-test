package arrayandslice

func Sum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func SumAll(nums1 []int, nums2 []int) []int {
// 	return []int{Sum(nums1), Sum(nums2)}
// }

// better one
// func SumAll(numbersToSum ...[]int) []int {
// 	lengthOfNumbers := len(numbersToSum)
// 	sums := make([]int, lengthOfNumbers)
//
// 	for i, numbers := range numbersToSum {
// 		sums[i] = Sum(numbers)
// 	}
//
// 	return sums
// }

// refactor
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// func SumAllTails(numbersToSum ...[]int) []int {
// 	var sums []int
// 	for _, numbers := range numbersToSum {
// 		tail := numbers[1:]
// 		sums = append(sums, Sum(tail))
// 	}
//
// 	return sums
// }

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}
