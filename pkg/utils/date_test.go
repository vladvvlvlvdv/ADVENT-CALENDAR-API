package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestGetDaysInMonth(t *testing.T) {
	year := 2024
	daysInMonths := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	for i, daysInMonth := range daysInMonths {
		month := time.Month(i + 1)
		testDays := GetDaysInMonth(year, month)

		require.Equal(t, daysInMonth, testDays)
	}
}
