package internal

import "strconv"

func Color(hex string) int {
	color, _ := strconv.ParseInt(hex, 16, 0)

	return int(color)
}
