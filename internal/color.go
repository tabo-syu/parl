package internal

import "strconv"

func Color(hex string) int {
	color, _ := strconv.ParseInt(hex, 16, 0)

	return int(color)
}

func ColorRed() int {
	return Color("ff0000")
}

func ColorGreen() int {
	return Color("00ff00")
}

func ColorOrange() int {
	return Color("ffa500")
}
