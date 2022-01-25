package math

func Add(nums ...int) int {
	var sum = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
