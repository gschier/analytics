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

type RollupPeriod int

const (
	PeriodDay    = RollupPeriod(time.Hour * 24)
	PeriodHour   = RollupPeriod(time.Hour)
	PeriodMinute = RollupPeriod(time.Minute)
)

func GetBucketStart(t time.Time, period RollupPeriod) time.Time {
	if period == PeriodDay {
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	} else if period == PeriodHour {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
	} else if period == PeriodMinute {
		return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, t.Location())
	}

	panic(fmt.Sprintf("Invalid rollup period %s", time.Duration(period)))
}
