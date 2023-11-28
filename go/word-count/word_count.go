package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	r := regexp.MustCompile(`\w+('\w+)?`)
	p := r.FindAllString(strings.ToLower(phrase), -1)
	var f Frequency = make(Frequency)
	for _, v := range p {
		f[strings.Trim(v, "'")] += 1
	}
	return f
}
