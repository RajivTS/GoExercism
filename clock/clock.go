package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hours   int
	minutes int
}

func FormatHours(h int) int {
	h = h % 24
	if h < 0 {
		h = 24 + h
	}
	return h
}

func FormatMinutes(m int) int {
	m = m % (24 * 60)
	if m < 0 {
		m += 24 * 60
	}
	return m
}

func New(h, m int) Clock {
	h = FormatHours(h)
	m = FormatMinutes(m)
	h = (h + FormatHours(m/60)) % 24
	m = m % 60
	return Clock{hours: h, minutes: m}
}

func (c Clock) Add(m int) Clock {
	m = FormatMinutes(m)
	c.hours += (c.minutes + m) / 60
	c.hours = c.hours % 24
	c.minutes = (c.minutes + m) % 60
	return c
}

func (c Clock) Subtract(m int) Clock {
	m = FormatMinutes(m)
	m = 24*60 - m
	return c.Add(m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
