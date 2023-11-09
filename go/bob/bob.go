// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"regexp"
)
// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	len_s := len(remark)

	r := regexp.MustCompile(`.*\?\s*$`)

	if strings.TrimSpace(remark) == "" || len_s < 1 {
		return "Fine. Be that way!"
	}

	upper_r := regexp.MustCompile(`[A-Z]+`)
	if upper_r.MatchString(remark) && remark == strings.ToUpper(remark) {
		if r.MatchString(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Whoa, chill out!"
		}
	} 
	if r.MatchString(remark) {
		return "Sure."
	}

	return "Whatever."
}
