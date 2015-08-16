package ts

import (
	"math"
	"time"
)

// Bucket represents a value at a given point in time
type Bucket struct {
	T time.Time `json:"t"`
	V *float64  `json:"v"`
}

// Series represents a single time-series.
type Series struct {
	Duration   time.Duration
	Resolution time.Duration
	buckets    map[int64]*Bucket
}

// NewSeries creates a new Series with the given duration and resolution
func NewSeries(duration time.Duration, resolution time.Duration) *Series {
	return &Series{
		Duration:   duration,
		Resolution: resolution,
		buckets:    make(map[int64]*Bucket, 0),
	}
}

// calculate the bucket time for a given Time
func (s *Series) floor(t time.Time) time.Time {
	return t.Truncate(s.Resolution)
}

// calculate the index for our bucket
func (s *Series) index(t time.Time) int64 {
	return int64(math.Mod(float64(s.floor(t).Unix()), float64(s.Duration.Seconds())))
}

// get a bucket for a given timestamp
// returns a fresh one if the bucket is non-existent
// or holding an old timestamp
func (s *Series) get(t time.Time) *Bucket {
	floor := s.floor(t)
	idx := s.index(t)

	bucket := s.buckets[idx]
	if bucket == nil || bucket.T != floor {
		b := &Bucket{floor, nil}
		s.buckets[idx] = b
		return b
	}

	return bucket
}

// Insert takes a given value at a given time and inserts a
// new bucket into the Series given the spec
func (s *Series) Insert(t time.Time, value float64) {
	b := s.get(t)
	b.V = &value
}

// Range takes a start and end time and returns a list of buckets that match
func (s *Series) Range(start time.Time, end time.Time) []*Bucket {
	var buckets []*Bucket
	startFloor := s.floor(start)
	endFloor := s.floor(end)

	now := time.Now()
	firstPossibleFloor := s.floor(now.Add(-1 * s.Duration))

	// sweep through our range of buckets
	for x := startFloor; x.Before(endFloor) || x.Equal(endFloor); x = x.Add(s.Resolution) {
		// don't return values beyound our Series boundaries or from the future
		if x.Before(firstPossibleFloor) || x.After(now) {
			continue
		}

		bucket := s.get(x)
		buckets = append(buckets, bucket)
	}

	return buckets
}

// FromDuration returns a range of Buckets, between Now and Now - d
func (s *Series) FromDuration(d time.Duration) []*Bucket {
	now := time.Now()
	start := now.Add(-1 * d).Add(s.Resolution)
	return s.Range(start, now)
}
