package letcodeexe

import "slices"

func longestCommonPrefix(strs []string) string {
	output := ""
	arr_len := len(strs)
	slices.Sort(strs)

	if arr_len > 1 {
		first := strs[0]
		last := strs[arr_len-1]
		for idx, v := range strs[0] {
			if first[idx] == last[idx] {
				output = output + string(v)
			} else {
				break
			}
		}
	}

	if arr_len == 1 {
		return strs[0]
	}

	return output
}
