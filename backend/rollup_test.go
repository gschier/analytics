package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRollupEvents(t *testing.T) {
	t.Run("Counts and uniques work", func(t *testing.T) {
		d := time.Date(2021, 1, 1, 1, 0, 0, 0, time.Now().Location())
		events := []RollupEvent{
			{Time: d.Add(-1 * time.Minute), UniqueKey: "user-1"},

			{Time: d.Add(0 * time.Minute), UniqueKey: "user-1"},
			{Time: d.Add(1 * time.Minute), UniqueKey: "user-1"},
			{Time: d.Add(1 * time.Minute), UniqueKey: "user-2"},

			{Time: d.Add(2 * time.Minute), UniqueKey: "user-1"},
			{Time: d.Add(2 * time.Minute), UniqueKey: "user-1"},
			{Time: d.Add(2 * time.Minute), UniqueKey: "user-1"},

			{Time: d.Add(5 * time.Minute), UniqueKey: "user-1"}, // Will get missed
		}

		buckets := Rollup(d, 5, PeriodMinute, events)

		require.Equal(t, 5, len(buckets), "Should have a bucket per minute")

		require.Equal(t, 1, int(buckets[0].Count))
		require.Equal(t, 1, int(buckets[0].Unique))

		require.Equal(t, 2, int(buckets[1].Count))
		require.Equal(t, 2, int(buckets[1].Unique))

		require.Equal(t, 3, int(buckets[2].Count))
		require.Equal(t, 1, int(buckets[2].Unique))

		require.Equal(t, 0, int(buckets[3].Count))
		require.Equal(t, 0, int(buckets[3].Unique))

		require.Equal(t, 0, int(buckets[4].Count))
		require.Equal(t, 0, int(buckets[4].Unique))
	})

	t.Run("Buckets at the edges", func(t *testing.T) {
		d := time.Date(2021, 1, 1, 1, 0, 0, 0, time.Now().Location())
		events := []RollupEvent{
			{Time: d.Add(time.Minute * 0), UniqueKey: "user-1"},
			{Time: d.Add(time.Minute * 1), UniqueKey: "user-1"},
			{Time: d.Add(time.Minute * 2), UniqueKey: "user-1"},
			{Time: d.Add(time.Minute * 3), UniqueKey: "user-1"}, // Will get missed
		}

		buckets := Rollup(d, 3, PeriodMinute, events)

		require.Equal(t, 3, len(buckets), "Should have a bucket per minute")

		require.Equal(t, 1, int(buckets[0].Count))
		require.Equal(t, 1, int(buckets[0].Unique))

		require.Equal(t, 1, int(buckets[1].Count))
		require.Equal(t, 1, int(buckets[1].Unique))

		require.Equal(t, 1, int(buckets[2].Count))
		require.Equal(t, 1, int(buckets[2].Unique))
	})

	t.Run("Buckets use rounded start and end time", func(t *testing.T) {
		dExact := time.Date(2021, 1, 1, 1, 1, 0, 0, time.Now().Location())
		dFuzzy := time.Date(2021, 1, 1, 1, 1, 59, 999, time.Now().Location())
		buckets := Rollup(dFuzzy, 3, PeriodMinute, []RollupEvent{})

		require.Equal(t, 3, len(buckets), "Should have a bucket per minute")
		require.Equal(t, dExact, buckets[0].Start)
		require.Equal(t, dExact.Add(time.Minute), buckets[0].End)
	})
}
