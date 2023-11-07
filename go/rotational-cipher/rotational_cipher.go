package rotationalcipher

import (
	"bytes"
	"strings"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	shiftKey = shiftKey % 26
	var b bytes.Buffer
	//Use bytes.Buffer to avoid allocate and copy while appending/concat strings

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for _, v := range plain {
		var v_index int
		//v_index = strings.IndexRune(alphabet, unicode.ToLower(v))
		v_index = strings.IndexRune(alphabet, unicode.ToLower(v))

		if v_index == -1 {
			//if not alphabet, keep it as is
			b.WriteString(string(v))
		} else {
			shifted_val := (v_index + shiftKey) % 26
			//To handle wrapped around shift

			if unicode.IsUpper(v) {
				b.WriteString(strings.ToUpper(string(alphabet[shifted_val])))
			} else {
				b.WriteString(string(alphabet[shifted_val]))
			}

		}
	}
	return b.String()
}
