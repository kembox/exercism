package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	s := strconv.Itoa(n)
	sum := 0
	for i := 0; i < len(s); i++ {
		num, _ := strconv.Atoi(string(s[i]))
		sum += int(math.Pow(float64(num), float64(len(s))))
	}
	return sum == n
}
