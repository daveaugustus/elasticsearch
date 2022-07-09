package utils

import (
	"time"
)

// GetPossibleDates takes target year, month, day and past as parameters, where if past param
// is true, the the function returns possible dates from past in YYYY-MM-DD format else future
func GetPossibleDates(year, month, day int, past bool) []string {
	dates := []string{}
	startDate := time.Now()

	if past {
		endDate := startDate.AddDate(-1, 0, 0)
		for d := endDate; d.After(startDate) == false; d = d.AddDate(0, 0, 1) {
			dates = append(dates, d.Format("2006-01-02"))
		}
		return dates
	}

	end := startDate.AddDate(0, 1, 0)
	for d := startDate; d.After(end) == false; d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format("2006-01-02"))
	}
	return dates
}
