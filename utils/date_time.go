package utils

import (
	"elasticsearch/constants"
	"time"
)

// GetPossibleDates takes target year, month, day and past as parameters, where if past param
// is true, the the function returns possible dates from past in YYYY-MM-DD format else future
func GetPossibleDates(year, month, day int, past bool) []string {
	dates := []string{}
	startDate := time.Now()

	if past {
		endDate := startDate.AddDate(-year, -month, -day)
		for d := endDate; d.After(startDate) == false; d = d.AddDate(0, 0, 1) {
			dates = append(dates, d.Format(constants.YYYYMMDD))
		}
		return dates
	}

	endDate := startDate.AddDate(year, month, day)
	for d := startDate; d.After(endDate) == false; d = d.AddDate(0, 0, 1) {
		dates = append(dates, d.Format(constants.YYYYMMDD))
	}
	return dates
}
