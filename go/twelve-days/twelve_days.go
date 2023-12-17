package twelve

import (
	"fmt"
	"strings"
)

// {12, "On the twelfth day of Christmas my true love gave to me:
var gifts = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"second":   "two Turtle Doves, and a Partridge in a Pear Tree.",
	"third":    "three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"fourth":   "four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"fifth":    "five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"sixth":    "six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"seventh":  "seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"eighth":   "eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"ninth":    "nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"tenth":    "ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"eleventh": "eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming, eleven Pipers Piping, ten Lords-a-Leaping, nine Ladies Dancing, eight Maids-a-Milking, seven Swans-a-Swimming, six Geese-a-Laying, five Gold Rings, four Calling Birds, three French Hens, two Turtle Doves, and a Partridge in a Pear Tree.",
}

var numbers = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", numbers[i-1], gifts[numbers[i-1]])
}

func Song() string {
	var b strings.Builder
	for i := 1; i < 13; i++ {
		b.WriteString(Verse(i) + "\n")
	}
	return strings.TrimRight(b.String(), "\n")
}
