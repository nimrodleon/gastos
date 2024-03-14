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

// AdjustDateRange adjusts the date range by returning the start and end of the day for the given dates.
// It takes two date strings (in the format "2006-01-02") as input: fromDateStr and toDateStr.
// It returns the start of the day for fromDate and the end of the day for toDate.
// If there is an error parsing the date strings, it returns zero time values and the error.
func AdjustDateRange(fromDateStr, toDateStr string) (time.Time, time.Time, error) {
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
