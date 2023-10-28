package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, 世界"
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, string(v))
	}
}
