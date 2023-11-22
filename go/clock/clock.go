package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour   int
	minute int
}

func New(h, m int) Clock {
	total_mins := (h%24)*60 + m
	for total_mins < 0 {
		total_mins = 24*60 + total_mins
	}
	h = (total_mins / 60) % 24
	m = total_mins % 60
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	total_mins := c.hour*60 + c.minute + m
	c.hour = (total_mins / 60) % 24
	c.minute = total_mins % 60
	return c

}

func (c Clock) Subtract(m int) Clock {
	total_mins := c.hour*60 + c.minute - m
	for total_mins < 0 {
		total_mins = 24*60 + total_mins
	}
	c.hour = total_mins / 60
	c.minute = total_mins % 60
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
