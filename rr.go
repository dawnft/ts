package rr

import (
	"math"
	"time"
)

type Bucket struct {
	T time.Time
	V *int64
}

// TS represents a single time-series.
type TS struct {
	Resolution time.Duration
	Size       int64
	Buckets    map[int64]*Bucket
}

func (ts *TS) floor(t time.Time) time.Time {
	return t.Round(ts.Resolution)
}

func (ts *TS) index(t time.Time) int64 {
	return int64(math.Mod(float64(ts.floor(t).Unix()), float64(ts.Size)))
}

func (ts *TS) Insert(t time.Time, value int64) {
	p := &Bucket{ts.floor(t), &value}
	idx := ts.index(p.T)
	ts.Buckets[idx] = p
}

// TODO - test for buckets that are out of range
// TODO - test for buckets with old epoch values
func (ts *TS) Get(t time.Time) *Bucket {
	floor := ts.floor(t)
	idx := ts.index(t)

	bucket := ts.Buckets[idx]
	if bucket == nil {
		return &Bucket{floor, nil}
	}

	return bucket
}

// TODO - test for start being greater than finish
func (ts *TS) GetRange(start time.Time, end time.Time) ([]*Bucket, error) {
	buckets := make([]*Bucket, 0)
	start_floor := ts.floor(start)
	end_floor := ts.floor(end)

	for x := start_floor; x.Before(end_floor); x = x.Add(ts.Resolution) {
		bucket := ts.Get(x)
		if bucket == nil {
			continue
		}
		buckets = append(buckets, bucket)
	}

	return buckets, nil
}
