package main

import (
	"fmt"
	"time"
)

type Bucket struct {
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Total  int64     `json:"total"`
	Unique int64     `json:"unique"`
}

const (
	PeriodDay    = time.Hour * 24
	PeriodHour   = time.Hour
	PeriodMinute = time.Minute
)

func CeilToPeriod(t time.Time, period time.Duration) time.Time {
	if period == PeriodDay {
		return time.Date(t.Year(), t.Month(), t.Day()+1, 0, 0, 0, 0, time.UTC)
	} else if period == PeriodHour {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, 0, 0, 0, time.UTC)
	} else if period == PeriodMinute {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute()+1, 0, 0, time.UTC)
	}

	panic(fmt.Sprintf("Invalid rollup period %s", period))
}
