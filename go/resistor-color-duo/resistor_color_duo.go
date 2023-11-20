package resistorcolorduo

import "strconv"

// Value should return the resistance value of a resistor with a given colors.
var colors_code map[string]int = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func Value(colors []string) int {
	s := ""
	for _, v := range colors[:2] {
		s += strconv.Itoa(colors_code[v])
	}
	ret, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	} else {
		return ret
	}
}
