package cipher

import (
	"math"
	"regexp"
	"strings"
	"math/rand"
	"time"
)


// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
type shift int
type vigenere []rune
var my_rand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func NewCaesar() Cipher {
	var s shift = 3
	return s
}

func NewShift(distance int) Cipher {
	if math.Abs(float64(distance)) > 25 || math.Abs(float64(distance)) < 1 {
		return nil
	} else {
		var s shift = shift(distance)
		return s
	}
}
func shift_n_wrap(r rune, s shift) rune {
	n := r + rune(s)
	if n > 'z' {
		n = 'a' + n - 'z' - 1
	} else if n < 'a' {
		n = 'z' - ('a' - n) + 1
	}
	return n
}
func (c shift) Encode(input string) string {
	var b strings.Builder
	r := regexp.MustCompile(`[^a-zA-Z]`)
	input = r.ReplaceAllString(input,"")
	input = strings.ToLower(input)
	//fmt.Println("input: ",input)
	for _, v := range input {
		b.WriteRune(shift_n_wrap(v, c))
	}
	return b.String()
}

func (c shift) Decode(input string) string {
	var b strings.Builder
	for _, v := range input {
		b.WriteRune(shift_n_wrap(v, -c))
	}
	return b.String()
}

func NewVigenere(key string) Cipher {
	var vig vigenere
	r := regexp.MustCompile(`[^a-z]`)
	if r.MatchString(key) {
		return nil
	} else if key == strings.Repeat("a", 25) {
		return nil
	} else if len(key) == 0 {
		for i := 0; i < 100; i++ {
			vig[i] = letterRunes[my_rand.Intn(len(letterRunes))]
		}
	} else {
		for _, v := range key {
			vig = append(vig, v)
		}
	}
	return vig
}

func (v vigenere) Encode(input string) string {
	r := regexp.MustCompile(`[^a-zA-Z]`)
	input = r.ReplaceAllString(input,"")
	input = strings.ToLower(input)

	var b strings.Builder
	for i := 0; i < len(input); i++ {
		b.WriteRune(shift_n_wrap(rune(input[i]), shift(v[i])))
	}
	return b.String()
}

func (v vigenere) Decode(input string) string {
	var b strings.Builder
	for i := 0; i < len(input); i++ {
		b.WriteRune(shift_n_wrap(rune(input[i]), shift(-v[i])))
	}
	return b.String()
}