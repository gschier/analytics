package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestRollupEvents(t *testing.T) {
	t.Run("Counts and uniques work", func(t *testing.T) {
		events := []AnalyticsEvent{
			{CreatedAt: time.Now().Add(-time.Hour), Name: "MyEvents", SessionKey: "user-1"},

			{CreatedAt: time.Now().Add(-time.Minute * 1), Name: "MyEvent", SessionKey: "user-1"},
			{CreatedAt: time.Now().Add(-time.Minute * 1), Name: "MyEvent", SessionKey: "user-2"},

			{CreatedAt: time.Now().Add(-time.Minute * 2), Name: "MyEvent", SessionKey: "user-1"},
			{CreatedAt: time.Now().Add(-time.Minute * 2), Name: "MyEvent", SessionKey: "user-1"},
			{CreatedAt: time.Now().Add(-time.Minute * 2), Name: "MyEvent", SessionKey: "user-1"},

			{CreatedAt: time.Now().Add(time.Hour), Name: "MyEvents", SessionKey: "user-1"},
		}

		buckets := Rollup(time.Now().Add(-time.Minute*5), time.Now(), time.Minute, events)

		require.Equal(t, 5, len(buckets), "Should have a bucket per minute")

		require.Equal(t, 0, int(buckets[0].Count))
		require.Equal(t, 0, int(buckets[0].Unique))

		require.Equal(t, 0, int(buckets[1].Count))
		require.Equal(t, 0, int(buckets[1].Unique))

		require.Equal(t, 3, int(buckets[2].Count))
		require.Equal(t, 1, int(buckets[2].Unique))

		require.Equal(t, 2, int(buckets[3].Count))
		require.Equal(t, 2, int(buckets[3].Unique))

		require.Equal(t, 0, int(buckets[4].Count))
		require.Equal(t, 0, int(buckets[4].Unique))
	})

	t.Run("Buckets at start of period", func(t *testing.T) {
		events := []AnalyticsEvent{
			{CreatedAt: time.Now().Add(-time.Second * 32), Name: "MyEvents", SessionKey: "user-1"},
		}

		buckets := Rollup(time.Now().Add(-time.Minute*5), time.Now(), time.Minute, events)

		require.Equal(t, 5, len(buckets), "Should have a bucket per minute")

		require.Equal(t, 0, int(buckets[4].Count))
		require.Equal(t, 0, int(buckets[4].Unique))
	})
}
