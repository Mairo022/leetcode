package all_disappeared_numbers

func FindDisappearedNumbers(nums []int) []int {
	temp := make(map[int]bool)
	max := len(nums)

	disappearedNumbers := []int{}

	for _, num := range nums {
		if num > max {
			max = num
		}
		temp[num] = true
	}

	for num := 1; num <= max; num++ {
		if !temp[num] {
			disappearedNumbers = append(disappearedNumbers, num)
		}
	}

	return disappearedNumbers
}

var Data = [...][]int{
	{1},
	{5, 4, 6, 7, 9, 3, 22, 9, 5, 6},
	{1, 1},
	{2, 2},
	{1, 1, 2, 2},
}
