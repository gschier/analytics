package main

import (
	"github.com/axiomhq/hyperloglog"
	"sort"
	"time"
)

type Bucket struct {
	Start  time.Time
	End    time.Time
	Count  uint64
	Unique uint64
	HLL    []byte
}

func Rollup(start, end time.Time, period time.Duration, events []AnalyticsEvent) []Bucket {
	// Ensure sorted by date oldest -> newest
	sort.Slice(events, func(i, j int) bool {
		return events[j].CreatedAt.Before(events[i].CreatedAt)
	})

	// Put events into bucket
	buckets := make([]Bucket, end.Sub(start)/period)
	for n := 0; n < len(buckets); n++ {
		b := &buckets[n]
		b.Count = 0
		b.Start = start.Add(period * time.Duration(n))
		b.End = start.Add(period * time.Duration(n+1))

		hll := hyperloglog.New()
		for _, e := range events {
			if e.CreatedAt.Before(b.Start) || e.CreatedAt.After(b.End) {
				continue
			}

			b.Count += 1
			hll.Insert([]byte(e.SessionKey))
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

// OptimalPeriod calculates the bucket size to use for each rollup between two dates
func OptimalPeriod(start, end time.Time) time.Duration {
	delta := end.Sub(start)
	if delta/time.Minute < 30 {
		return time.Minute
	}
	if delta/time.Hour > 30 {
		return time.Hour
	}
	return time.Hour * 24
}
