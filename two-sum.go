package letcodeexe

func twoSum(nums []int, target int) []int {
	var arr = make(map[int]int)

	for index, data := range nums {
		left := target - data

		if val, ok := arr[left]; ok {
			return []int{val, index}
		}
		arr[data] = index
	}
	return []int{}
}
