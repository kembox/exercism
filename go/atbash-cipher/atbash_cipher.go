package atbash

import (
	"regexp"
	"strings"
)

func Atbash(s string) string {
	s = strings.ToLower(s)
	r := regexp.MustCompile(`[^a-z0-9]`)
	rn := regexp.MustCompile(`[0-9]`)

	s = r.ReplaceAllString(s, "")

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var b strings.Builder
	cnt := 0
	for _, v := range s {
		i := strings.Index(alphabet, string(v))
		if i != -1 {
			b.WriteString(string(alphabet[len(alphabet)-i-1]))
		} else if rn.MatchString(string(v)) {
			b.WriteString(string(v))
		}
		cnt++
		if cnt == 5 {
			b.WriteString(" ")
			cnt = 0
		}
	}
	return strings.TrimSpace(b.String())
}
