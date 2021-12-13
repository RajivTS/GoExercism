package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:4:5"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime
}

// HasPassed returns whether a date has passed
func HasPassed(date string) bool {
	// "July 25, 2019 13:45:00"
	layout := "January 1, 2006 15:4:5"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon
func IsAfternoonAppointment(date string) bool {
	// "Thursday, July 25, 2019 13:45:00"
	layout := "Monday, January 1, 2006 15:4:5"
	parsedTime, _ := time.Parse(layout, date)
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time
func Description(date string) string {
	layout := "1/2/2006 15:4:5"
	t, _ := time.Parse(layout, date)
	formatStr := "You have an appointment on %s, %s %d, %d, at %d:%d."
	return fmt.Sprintf(formatStr, t.Weekday().String(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	year := time.Now().Year()
	return time.Date(year, time.September, 15, 0, 0, 0, 0, time.UTC)
}
