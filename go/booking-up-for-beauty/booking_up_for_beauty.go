package booking

import (
	"fmt"
	"time"
)

const LAYOUT = "1/2/2006 15:04:05"

func Schedule(date string) time.Time {
	t, _ := time.Parse(LAYOUT, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	h, _, _ := t.Clock()
	return 12 <= h && h < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	f := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", f)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	t := time.Now()
	t = time.Date(
		t.Year(), 9, 15,
		0, 0, 0, 0, time.UTC)
	return t
}
