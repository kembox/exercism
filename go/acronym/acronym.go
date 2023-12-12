package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	r := regexp.MustCompile(`[^a-zA-Z -]`)
	s = r.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "-", " ")

	if len(s) == 0 {
		return ""
	}
	var sl []string = strings.Fields(s)
	var b strings.Builder
	for _, v := range sl {
		b.WriteByte(v[0])
	}
	return strings.ToUpper(b.String())
}
