package times

import "time"

func DaysInMonth(inputString string) (int, error) {
	// Parse the input string to get the year and month
	parsedTime, err := time.Parse("2006-01", inputString)
	if err != nil {
		return 0, err
	}

	// Get the year and month from the parsed time
	year, month, _ := parsedTime.Date()

	// Get the first day of the next month
	firstDayOfNextMonth := time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC)

	// Subtract one day to get the last day of the current month
	lastDayOfCurrentMonth := firstDayOfNextMonth.Add(-time.Second)

	// Calculate the number of days in the month
	daysInMonth := lastDayOfCurrentMonth.Day()

	return daysInMonth, nil
}
