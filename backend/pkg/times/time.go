package times

import (
	"backend/pkg/logger"
	"time"
)

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

func ParseTime(dateTimeString string) (*time.Time, error) {
	// Parse the string into a time.Time value
	parsedTime, err := time.Parse("2006-01-02 15:04:05", dateTimeString)
	if err != nil {
		return nil, err
	}
	return &parsedTime, err
}

func IsTimeBefore(srcTime time.Time, ref time.Time) (bool, error) {
	curTime := srcTime.UTC().Format("150405")
	logger.Console(srcTime)
	logger.Console(curTime)
	curTimeObj, err := time.Parse("150405", curTime)
	if err != nil {
		return false, err
	}
	refTime := ref.UTC().Format("150405")

	logger.Console(ref)
	logger.Console(refTime)
	refTimeObj, err := time.Parse("150405", refTime)
	if err != nil {
		return false, err
	}
	timeDifference := curTimeObj.Sub(refTimeObj)	
	timeOnlyComparison := int(timeDifference.Minutes())
	return timeOnlyComparison < 0, nil
}

func IsTimeAfter(srcTime time.Time, ref time.Time) (bool, error) {
	curTime := srcTime.UTC().Format("150405")
	curTimeObj, err := time.Parse("150405", curTime)
	if err != nil {
		return false, err
	}
	refTime := ref.UTC().Format("150405")

	refTimeObj, err := time.Parse("150405", refTime)
	if err != nil {
		return false, err
	}
	timeDifference := curTimeObj.Sub(refTimeObj)	
	timeOnlyComparison := int(timeDifference.Minutes())
	return timeOnlyComparison > 0, nil
}
