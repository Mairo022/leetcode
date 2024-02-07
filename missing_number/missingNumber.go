package missing_number

func MissingNumber(nums []int) int {
	temp := make(map[int]bool)
	min, max := nums[0], nums[0]

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		temp[num] = true
	}

	for num := min + 1; num < max; num++ {
		if !temp[num] {
			return num
		}
	}

	if len(nums) == max && min == 1 {
		return 0
	}

	return max + 1
}

var Data = [...][]int{
	{3, 0, 1},
	{0, 1},
	{9, 6, 4, 2, 3, 5, 7, 0, 1},
	{1, 2},
}
