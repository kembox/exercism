package main

import (
	"fmt"
)

func main() {
	c := fmt.Sprintf("%c", 65)
	fmt.Printf("%T\n", c)
	fmt.Println(string('A' + 2))
	fmt.Printf("%T\n", string('A'+1))
}
