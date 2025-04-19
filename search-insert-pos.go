package letcodeexe

func searchInsert(nums []int, target int) int {
	last := 0
	if target < nums[0] {
		return 0
	}

	for idx, val := range nums {
		if val == target {
			return idx
		}
		if val >= target && target > last {
			return idx
		}
		last = val
	}

	return len(nums)
}
