package util

import (
	"log"
	"time"
)

// DateIntervals receives a start and end strings representing dates
// and returns a map of strings with all 30 days intervals between both dates
func DateIntervals(start, end string) map[string]string {
	layout := "02 Jan 06"

	startDate, err := time.Parse(layout, start)
	if err != nil {
		log.Fatal(err)
	}

	endDate, err := time.Parse(layout, end)
	if err != nil {
		log.Fatal(err)
	}

	date := startDate
	intervals := make(map[string]string)

	for date.Before(endDate.AddDate(0, -1, 0)) {
		intervals[date.String()[:10]] = date.AddDate(0, 1, 0).String()[:10]
		date = date.AddDate(0, 1, 0)
	}

	diff := endDate.Sub(date).Hours() / 24

	intervals[date.String()[:10]] = date.AddDate(0, 0, int(diff)).String()[:10]
	return intervals
}
