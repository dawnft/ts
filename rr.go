package rr

import (
	"math"
	"time"
)

// Bucket represents a value at a given point in time
type Bucket struct {
	T time.Time
	V *int64
}

// TS represents a single time-series.
type TS struct {
	Duration   time.Duration
	Resolution time.Duration
	Buckets    map[int64]*Bucket
}

func (ts *TS) floor(t time.Time) time.Time {
	return t.Round(ts.Resolution)
}

func (ts *TS) index(t time.Time) int64 {
	return int64(math.Mod(float64(ts.floor(t).Unix()), float64(ts.Duration.Seconds())))
}

// Insert takes a given value at a given time and inserts a
// new bucket into the TS given the spec
func (ts *TS) Insert(t time.Time, value int64) {
	b := &Bucket{ts.floor(t), &value}
	idx := ts.index(b.T)
	ts.Buckets[idx] = b
}

// TODO - test for buckets that are out of range of the TS
func (ts *TS) get(t time.Time) *Bucket {
	floor := ts.floor(t)
	idx := ts.index(t)

	bucket := ts.Buckets[idx]
	if bucket == nil || bucket.T != floor {
		return &Bucket{floor, nil}
	}

	return bucket
}

// Range takes a start and end time and returns a list of buckets that match
func (ts *TS) Range(start time.Time, end time.Time) []*Bucket {
	var buckets []*Bucket
	startFloor := ts.floor(start)
	endFloor := ts.floor(end)

	for x := startFloor; x.Before(endFloor) || x.Equal(endFloor); x = x.Add(ts.Resolution) {
		bucket := ts.get(x)
		if bucket == nil {
			continue
		}
		buckets = append(buckets, bucket)
	}

	return buckets
}
