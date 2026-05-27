package clock

import "fmt"

type Clock struct {
    hours int
    minutes int
}

func New(h, m int) Clock {
    h += m/60
    m %= 60
    h %= 24
    if m < 0 {
        m += 60
        h--
    }
    if h < 0 {
        h += 24
    }
    return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	c.minutes += m
    c.hours += c.minutes/60
    c.minutes %= 60
    c.hours %= 24
    if c.minutes < 0 {
        c.minutes += 60
        c.hours--
    }
    if c.hours < 0 {
        c.hours += 24
    }
    return c
}

func (c Clock) Subtract(m int) Clock {
	c.minutes -= m
    c.hours += c.minutes/60
    c.minutes %= 60
    c.hours %= 24
    if c.minutes < 0 {
        c.minutes += 60
        c.hours--
    }
    if c.hours < 0 {
        c.hours += 24
    }
    return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
