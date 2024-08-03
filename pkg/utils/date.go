package utils

import (
	"time"
)

func GetDaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Day()
}

func GetDaysCount(month int, ignore bool) int {
	now := time.Now()

	day := now.Day()
	daysInMonth := GetDaysInMonth(now.Year(), time.Month(month))

	if day >= daysInMonth || ignore {
		return daysInMonth
	}

	return day
}
