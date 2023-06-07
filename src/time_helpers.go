package main

import (
	"time"
)

func createDiffTimeFromTwoTimes(time1 time.Time, time2 time.Time) time.Duration {
	difference := time1.Sub(time2)
	return difference
}

func createTimeFromString(dateString string) (time.Time, error) {
	date, error := time.Parse("2006-01-02T15:14:52Z", dateString)
	if error != nil {
		return time.Time{}, error
	}
	return date, nil
}

func createTimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

func createTimeFromUnixNano(unixNano int64) time.Time {
	return time.Unix(0, unixNano)
}
