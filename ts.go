package ts

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
	buckets    map[int64]*Bucket
}

// NewTS creates a new TS with the given duration and resolution
func NewTS(duration time.Duration, resolution time.Duration) *TS {
	return &TS{
		Duration:   duration,
		Resolution: resolution,
		buckets:    make(map[int64]*Bucket, 0),
	}
}

func (ts *TS) floor(t time.Time) time.Time {
	return t.Truncate(ts.Resolution)
}

func (ts *TS) index(t time.Time) int64 {
	return int64(math.Mod(float64(ts.floor(t).Unix()), float64(ts.Duration.Seconds())))
}

// Insert takes a given value at a given time and inserts a
// new bucket into the TS given the spec
func (ts *TS) Insert(t time.Time, value int64) {
	b := &Bucket{ts.floor(t), &value}
	idx := ts.index(b.T)
	ts.buckets[idx] = b
}

func (ts *TS) get(t time.Time) *Bucket {
	floor := ts.floor(t)
	idx := ts.index(t)

	bucket := ts.buckets[idx]
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

	now := time.Now()
	firstPossibleFloor := ts.floor(now.Add(-1 * ts.Duration))

	for x := startFloor; x.Before(endFloor) || x.Equal(endFloor); x = x.Add(ts.Resolution) {
		if x.Before(firstPossibleFloor) || x.After(now) {
			continue
		}

		bucket := ts.get(x)
		if bucket == nil {
			continue
		}
		buckets = append(buckets, bucket)
	}

	return buckets
}
