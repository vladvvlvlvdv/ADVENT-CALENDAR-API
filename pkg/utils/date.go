package utils

import "time"

func DaysInMonth(year int, month time.Month) int {
	return time.Date(year, month+1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, -1).Day()
}
