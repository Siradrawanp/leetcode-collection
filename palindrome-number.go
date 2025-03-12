package letcodeexe

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	final := 0
	temp := x
	for temp > 0 {
		sisa := temp % 10
		final = (final * 10) + sisa
		temp = temp / 10
	}

	if final == x {
		return true
	}

	return false
}
