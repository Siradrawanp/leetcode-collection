package letcodeexe

import "strings"

func romanToInt(s string) int {
	ret := 0
	arr := strings.Split(s, "")
	roman := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50,
		"C": 100, "D": 500, "M": 1000,
	}

	last := 0
	for i := len(arr); i > 0; i-- {
		idx := i - 1
		rom := arr[idx]
		value := roman[rom]

		if value < last {
			ret = ret - value
		} else {
			ret = ret + value
		}
		last = value
	}

	return ret
}
