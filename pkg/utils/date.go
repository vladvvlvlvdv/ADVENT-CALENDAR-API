package utils

import (
	"time"
)

func GetDaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Day()
}

func GetDayByTimeZone(timeZone string) (int, error) {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		return 0, err
	}

	return time.Now().In(loc).Day(), nil
}

func GetDaysCount(month int, day int, ignore bool) int {
	now := time.Now()

	daysInMonth := GetDaysInMonth(now.Year(), time.Month(month))

	if day >= daysInMonth || ignore {
		return daysInMonth
	}

	return day
}
