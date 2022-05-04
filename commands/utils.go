package commands

import "strings"

func loadingBar(min, max, current int) string {
	str := ""

	if current > max {
		current = max
	}
	if current < min {
		current = min
	}

	space := max - min
	if space > 40 {

	} else {
		load := current - min
		str = strings.Repeat("=", load)
		if load > 1 {
			str = str[:len(str)-1] + ">"
		}
		str += strings.Repeat(".", max-current)
	}

	return str
}
