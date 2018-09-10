package clock

import "fmt"

// Clock is define the clock type
type Clock int //

// New it is constructor
func New(h, m int) Clock {
	h = (h + (m / 60)) % 24
	if h < 0 {
		h = 24 + h
	}
	return Clock((h * 60) + (m % 60))
}

// String it return string of Clock
func (c Clock) String() string {
	h := (c / 60) % 24
	m := c % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Add it return added Clock
func (c Clock) Add(minutes int) Clock {
	return Clock(int(c) + minutes)
}

// Subtract it return substract Clock
func (c Clock) Subtract(minutes int) Clock {
	t := int(c) - minutes
	if t < 0 {
		t = 1440 + (t % 1440)
	}
	return Clock(t)
}
