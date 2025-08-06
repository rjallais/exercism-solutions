package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var (
		layout = "1/2/2006 15:04:05"
		t, err = time.Parse(layout, date)
	)

	if err != nil {
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var (
		layout    = "January 2, 2006 15:04:05"
		t, err    = time.Parse(layout, date)
		hasPassed = t.Compare(time.Now())
	)
	if err != nil {
		panic(err)
	}

	if hasPassed < 0 {
		return true
	} else {
		return false
	}
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var (
		layout = "Monday, January 2, 2006 15:04:05"
		t, err = time.Parse(layout, date)
	)
	if err != nil {
		panic(err)
	}

	if t.Hour() >= 12 && t.Hour() < 18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var t = Schedule(date)

	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.",
		t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
