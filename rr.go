package rr

import "math"

type Bucket struct {
	Epoch int64
	Value *int64
}

// TS represents a single time-series.
type TS struct {
	Resolution int64
	Size       int64
	Buckets    map[int64]*Bucket
}

func (t *TS) floor(epoch int64) int64 {
	return epoch / t.Resolution * t.Resolution
}

func (t *TS) index(epoch int64) int64 {
	return int64(math.Mod(float64(t.floor(epoch)), float64(t.Size)))
}

func (t *TS) Insert(epoch int64, value int64) {
	p := &Bucket{t.floor(epoch), &value}
	idx := t.index(p.Epoch)
	t.Buckets[idx] = p
}

// TODO - test for buckets that are out of range
// TODO - test for buckets with old epoch values
func (t *TS) Get(epoch int64) *Bucket {
	floor := t.floor(epoch)
	idx := t.index(epoch)

	bucket := t.Buckets[idx]
	if bucket == nil {
		return &Bucket{floor, nil}
	}

	return bucket
}

// TODO - test for start being greater than finish
func (t *TS) GetRange(start int64, end int64) ([]*Bucket, error) {
	buckets := make([]*Bucket, 0)
	start_floor := t.floor(start)
	end_floor := t.floor(end)

	for x := start_floor; x <= end_floor; x = x + t.Resolution {
		bucket := t.Get(x)
		if bucket == nil {
			continue
		}
		buckets = append(buckets, bucket)
	}

	return buckets, nil
}
