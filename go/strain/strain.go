package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](col []T, predicate func(T) bool) []T {
	var my_col []T
	for _, v := range col {
		if predicate(v) {
			my_col = append(my_col, v)
		}
	}
	return my_col
}

func Discard[T any](col []T, predicate func(T) bool) []T {
	var my_col []T
	for _, v := range col {
		if !predicate(v) {
			my_col = append(my_col, v)
		}
	}
	return my_col
}
