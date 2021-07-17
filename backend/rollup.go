package main

import (
	"github.com/axiomhq/hyperloglog"
	"log"
	"sort"
	"time"
)

type Bucket struct {
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
	Count  uint64    `json:"count"`
	Unique uint64    `json:"unique"`
	HLL    []byte    `json:"-"`
}

type RollupPeriod int

const (
	PeriodDay    = RollupPeriod(time.Hour * 24)
	PeriodHour   = RollupPeriod(time.Hour)
	PeriodMinute = RollupPeriod(time.Minute)
)

type RollupEvent struct {
	Time      time.Time
	UniqueKey string
}

func RollupPageviews(start time.Time, window int, period RollupPeriod, pageviews []AnalyticsPageview) []Bucket {
	events := make([]RollupEvent, len(pageviews))
	for i := range pageviews {
		events[i] = RollupEvent{
			Time:      pageviews[i].CreatedAt,
			UniqueKey: pageviews[i].SID,
		}
	}
	return Rollup(start, window, period, events)
}

// func RollupEvents(start time.Time, window int, period RollupPeriod, pageviews []AnalyticsEvent) []Bucket {
// 	events := make([]RollupEvent, len(pageviews))
// 	for i := range pageviews {
// 		events[i] = RollupEvent{
// 			Time:      pageviews[i].CreatedAt,
// 			UniqueKey: pageviews[i].Name,
// 		}
// 	}
// 	return Rollup(start, window, period, events)
// }

func Rollup(start time.Time, window int, period RollupPeriod, events []RollupEvent) []Bucket {
	end := start.Add(time.Duration(window) * time.Duration(period))

	if period == PeriodDay {
		start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
		end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, end.Location())
	} else if period == PeriodHour {
		start = time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), 0, 0, 0, start.Location())
		end = time.Date(end.Year(), end.Month(), end.Day(), end.Hour(), 0, 0, 0, end.Location())
	} else if period == PeriodMinute {
		start = time.Date(start.Year(), start.Month(), start.Day(), start.Hour(), start.Minute(), 0, 0, start.Location())
		end = time.Date(end.Year(), end.Month(), end.Day(), end.Hour(), end.Minute(), 0, 0, end.Location())
	} else {
		log.Fatalf("Invalid rollup period %s", time.Duration(period))
	}

	// Ensure sorted by date oldest -> newest
	sort.Slice(events, func(i, j int) bool {
		return events[j].Time.Before(events[i].Time)
	})

	// Put events into bucket
	buckets := make([]Bucket, end.Sub(start)/time.Duration(period))
	for n := 0; n < len(buckets); n++ {
		b := &buckets[n]
		b.Start = start.Add(time.Duration(period) * time.Duration(n))
		b.End = start.Add(time.Duration(period) * time.Duration(n+1))

		hll := hyperloglog.New()
		for _, e := range events {
			if e.Time.Before(b.Start) || e.Time.After(b.End) || e.Time.Equal(b.End) {
				continue
			}

			b.Count += 1
			hll.Insert([]byte(e.UniqueKey))
		}

		b.Unique = hll.Estimate()
		data, err := hll.MarshalBinary()
		if err != nil {
			panic("Failed to marshall HLL: " + err.Error())
		}

		b.HLL = data
	}

	return buckets
}
