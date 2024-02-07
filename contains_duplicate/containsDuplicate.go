package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	temp := make(map[int]bool)

	for _, num := range nums {
		if temp[num] {
			return true
		}
		temp[num] = true
	}
	return false
}

var Data = [...][]int{
	{1, 2, 3, 1},
	{1, 2, 3, 4},
	{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
}
