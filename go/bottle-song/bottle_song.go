package bottlesong

import (
	"fmt"
	"strings"
)

var n_to_t = map[int]string{
	0:  "no",
	1:  "One",
	2:  "Two",
	3:  "Three",
	4:  "Four",
	5:  "Five",
	6:  "Six",
	7:  "Seven",
	8:  "Eight",
	9:  "Nine",
	10: "Ten",
}

var str_template = `%s %s hanging on the wall,|%s %s hanging on the wall,|And if one green bottle should accidentally fall,|There'll be %s %s hanging on the wall.||`

func Recite(startBottles, takeDown int) []string {
	var b strings.Builder
	for i := 0; i < takeDown; i++ {
		var last_phrase = "green bottles"
		var first_phrase = "green bottles"
		if (startBottles - i - 1) == 1 {
			last_phrase = "green bottle"
		}
		if (startBottles - i) == 1 {
			first_phrase = "green bottle"
		}
		s := fmt.Sprintf(str_template,
			n_to_t[startBottles-i],
			first_phrase,
			n_to_t[startBottles-i],
			first_phrase,
			strings.ToLower(n_to_t[startBottles-i-1]),
			last_phrase)

		if i == takeDown-1 {
			b.WriteString(strings.TrimRight(s, "|"))
		} else {
			b.WriteString(s)
		}
	}

	return strings.Split(b.String(), "|")
}
