package letcodeexe

func isValid(s string) bool {
	open := []string{}

	for _, v := range s {
		idx := len(open)

		if string(v) == "{" || string(v) == "(" || string(v) == "[" {
			open = append(open, string(v))
		} else if idx != 0 {
			switch string(v) {
			case "}":
				if open[idx-1] == "{" {
					open = open[:(idx - 1)]
				} else {
					return false
				}
			case ")":
				if open[idx-1] == "(" {
					open = open[:(idx - 1)]
				} else {
					return false
				}
			case "]":
				if open[idx-1] == "[" {
					open = open[:(idx - 1)]
				} else {
					return false
				}
			}
		} else {
			return false
		}
	}

	if len(open) != 0 {
		return false
	}

	return true
}
