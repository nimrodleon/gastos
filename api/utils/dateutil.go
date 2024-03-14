package utils

import "time"

// CreateStartAndEndOfDay returns the start or end of the day for a given date.
// If isEndDate is true, it returns the end of the day (23:59:59.999999999).
// If isEndDate is false, it returns the start of the day (00:00:00).
func CreateStartAndEndOfDay(date time.Time, isEndDate bool) time.Time {
	if isEndDate {
		return date.Truncate(time.Hour * 24).Add(time.Hour*24 - time.Nanosecond)
	}
	return date.Truncate(time.Hour * 24)
}

// ParseAndAdjustDate parses the fromDateStr and toDateStr strings into time.Time values
// using the "2006-01-02" format. It also adjusts the parsed dates to the start and end of the day.
// It returns the adjusted start and end of the day time.Time values, along with any error encountered.
func ParseAndAdjustDate(fromDateStr, toDateStr string) (time.Time, time.Time, error) {
	fromDate, err := time.Parse("2006-01-02", fromDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	toDate, err := time.Parse("2006-01-02", toDateStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	startOfDay := CreateStartAndEndOfDay(fromDate, false)
	endOfDay := CreateStartAndEndOfDay(toDate, true)
	return startOfDay, endOfDay, nil
}
